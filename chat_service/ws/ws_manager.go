package ws

import (
	"fmt"
	"github.com/eduhub/helper"
	"github.com/eduhub/repositories"
	"github.com/eduhub/service"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"sync"
)

type WebSocketManager struct {
	Db *gorm.DB
}

var once sync.Once
var Rooms []*helper.Room
var wbManager *WebSocketManager
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetWebSocketManager(Db *gorm.DB) *WebSocketManager {
	once.Do(func() {
		wbManager = &WebSocketManager{Db: Db}
	})
	return wbManager
}

// InitChatRoom -> Authorises & Creates/Adds users to the chat rooms
func (wsm *WebSocketManager) InitChatRoom(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	user, err := service.GetAuthServiceInstance().AuthoriseAndFetchUser(token)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Fprintf(w, "error authorising user")
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("User (%+v)", user.Username)
	sender := user.Username
	receiver := r.URL.Query().Get("receiver")
	roomID := sender + "_" + receiver // Can encrypt the keys and then use it as RoomID
	wsm.AddToRoom(sender, receiver, roomID, conn, int64(user.ID))
}

// AddToRoom --> Manages/Create existing rooms
// Returns:
// The Room to which user is added
func (wsm *WebSocketManager) AddToRoom(senderID, receiverID, roomID string, conn *websocket.Conn, joinee int64) *helper.Room {
	key1 := senderID + "_" + receiverID
	key2 := receiverID + "_" + senderID

	repo := repositories.GetRoomsRepositoryInstance(wsm.Db)
	for _, room := range Rooms {
		if room.RoomID == key1 || room.RoomID == key2 {
			connections := room.Connections
			connections[senderID] = conn
			room.Connections = connections
			fmt.Printf("Subscribing existing room (%+v) (%+v)", room.RoomID, len(room.Connections))
			repo.UpsertRoomAndJoinee(roomID, joinee)
			wsm.Subscribe(conn, room)
			return room
		}
	}
	connMap := make(map[string]*websocket.Conn)
	connMap[senderID] = conn
	room := &helper.Room{
		SenderUsername:   senderID,
		ReceiverUsername: receiverID,
		RoomID:           roomID,
		Broadcaster:      make(chan *helper.Message),
		Connections:      connMap,
	}
	Rooms = append(Rooms, room)

	repo.UpsertRoomAndJoinee(roomID, joinee)
	fmt.Printf("Created new room (%+v) (%+v) \n", room.RoomID, len(room.Connections))
	wsm.Subscribe(conn, room)
	return room
}

// Subscribe --> Subscribes the user to the incoming Broadcasts
func (wsm *WebSocketManager) Subscribe(conn *websocket.Conn, room *helper.Room) {
	go wsm.ListenBroadCastAndSendToConnectedClients(room)
	fmt.Println("Subscribed", room.ReceiverUsername)
	for {
		var msg helper.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		room.Broadcaster <- &msg
		repo := repositories.GetRoomsRepositoryInstance(wsm.Db)
		repo.PersistMessage(room.RoomID, room.SenderUsername, room.ReceiverUsername, msg.Content)
	}
}

// ListenBroadCastAndSendToConnectedClients --> Listens and Sends Broadcasts to all connected clients
func (wsm *WebSocketManager) ListenBroadCastAndSendToConnectedClients(room *helper.Room) {
	for {
		msg := <-room.Broadcaster
		for sender, conn := range room.Connections {
			if conn != nil && sender != msg.Sender {
				err := conn.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					conn.Close()
				}
			}
		}
	}
}
