package rooms

import (
	"time"

	"github.com/google/uuid"
)

type RoomModel struct {
	Id        uuid.UUID `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

func NewRoom() RoomModel {
	return RoomModel{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
