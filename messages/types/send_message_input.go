package types

type SendMessageInput struct {
	RoomId string `json:"roomId" form:"roomId" binding:"required"`
	Text   string `json:"text" form:"text" binding:"required"`
}
