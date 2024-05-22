package rooms

import (
	"context"
	"errors"
	"go_chat/rooms/types"

	db "go_chat/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRoom(input types.CreateRoomInput) (RoomModel, error) {
	existingRoomDocument := db.GetCollection("rooms").FindOne(context.TODO(), bson.D{{Key: "name", Value: input.Name}})

	var existingRoom RoomModel
	err := existingRoomDocument.Decode(&existingRoom)

	if err == nil {
		return RoomModel{}, errors.New("api.roomExists")
	}

	roomModel := NewRoom()
	roomModel.Id = uuid.New()
	roomModel.Name = input.Name

	_, err = db.GetCollection("rooms").InsertOne(context.TODO(), &roomModel)

	if err != nil {
		return RoomModel{}, err
	}

	return roomModel, nil
}
