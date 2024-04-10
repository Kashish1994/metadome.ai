package repositories

import (
	"github.com/eduhub/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

type RoomsRepository struct {
	Db *gorm.DB
}

var roomsRepo *RoomsRepository
var once sync.Once

func GetRoomsRepositoryInstance(Db *gorm.DB) *RoomsRepository {
	once.Do(func() {
		roomsRepo = &RoomsRepository{Db: Db}
	})
	return roomsRepo
}

func (r *RoomsRepository) PersistMessage(roomID string, senderID string, receiverID string, message string) error {
	res := strings.Split(roomID, "_")
	roomID2 := res[1] + "_" + res[0]
	rm := &models.Room{}
	r.Db.First(&models.Room{}).Last(rm, "room_id IN ?", []string{roomID2, roomID})
	r.Db.Create(&models.Message{
		RoomID:     rm.ID,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    message,
		CreatedAt:  time.Time{},
	})
	return nil
}

func (r *RoomsRepository) UpsertRoomAndJoinee(roomID string, joinee int64) error {
	res := strings.Split(roomID, "_")
	roomID2 := res[1] + "_" + res[0]
	rm := &models.Room{}
	r.Db.First(&models.Room{}).Last(rm, "room_id IN ?", []string{roomID2, roomID})
	if rm.ID == 0 {
		r.Db.Create(&models.Room{
			Joinees: pq.Int64Array{joinee},
			RoomID:  roomID,
		})
	} else {
		joinees := rm.Joinees
		for _, j := range joinees {
			if j == joinee {
				return nil
			}
		}
		joinees = append(joinees, joinee)
		rm.Joinees = joinees
		r.Db.Save(rm)
	}
	return nil
}
