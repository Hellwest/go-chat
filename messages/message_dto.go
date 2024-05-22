package messages

import "github.com/google/uuid"

type MessageDTO struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

func NewMessageDTO() MessageDTO {
	return MessageDTO{}
}
