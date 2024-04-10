package ws

import (
	"fmt"
	"github.com/eduhub/helper"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type WebSocketManager struct{}

var once sync.Once
var Rooms []*helper.Room
var wbManager *WebSocketManager
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetWebSocketManager() *WebSocketManager {
	once.Do(func() {
		wbManager = &WebSocketManager{}
	})
	return wbManager
}

// InitChatRoom -> Creates/Adds users to the chat rooms
func (wsm *WebSocketManager) InitChatRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	sender := r.URL.Query().Get("sender")
	receiver := r.URL.Query().Get("receiver")
	roomID := sender + "_" + receiver // Can encrypt the keys and then use it as RoomID
	wsm.AddToRoom(sender, receiver, roomID, conn)
}

// AddToRoom --> Manages/Create existing rooms
// Returns:
// The Room to which user is added
func (wsm *WebSocketManager) AddToRoom(senderID, receiverID, roomID string, conn *websocket.Conn) *helper.Room {
	key1 := senderID + "_" + receiverID
	key2 := receiverID + "_" + senderID

	for _, room := range Rooms {
		fmt.Printf("Room (%+v) Keys (%+v) (%+v) \n", room.RoomID, key2, key1)
		if room.RoomID == key1 || room.RoomID == key2 {
			connections := room.Connections
			connections[senderID] = conn
			room.Connections = connections
			fmt.Printf("Subscribing existing room (%+v) (%+v)", room.RoomID, len(room.Connections))
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
	fmt.Println(len(room.Connections))
	fmt.Printf("Created new room (%+v) (%+v)", room.RoomID, len(room.Connections))
	wsm.Subscribe(conn, room)
	return room
}

// Subscribe --> Subscribes the user to the incoming Broadcasts
func (wsm *WebSocketManager) Subscribe(conn *websocket.Conn, room *helper.Room) {
	go wsm.ListenBroadCastAndSendToConnectedClients(room)
	for {
		var msg helper.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Subscribed", room.ReceiverUsername)
		room.Broadcaster <- &msg
	}
}

// ListenBroadCastAndSendToConnectedClients --> Listens and Sends Broadcasts to all connected clients
func (wsm *WebSocketManager) ListenBroadCastAndSendToConnectedClients(room *helper.Room) {
	for {
		msg := <-room.Broadcaster
		for sender, conn := range room.Connections {
			if conn != nil && sender != msg.Username {
				err := conn.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					conn.Close()
				}
			}
		}
	}
}
