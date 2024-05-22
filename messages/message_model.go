package messages

import (
	"go_chat/auth"
	"go_chat/rooms"
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id        uuid.UUID `bson:"id" json:"id"`
	Text      string    `bson:"text" json:"text"`
	Sender    auth.UserModel
	Room      rooms.RoomModel
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

func NewMessage() MessageModel {
	return MessageModel{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
