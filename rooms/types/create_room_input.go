package types

type CreateRoomInput struct {
	Name string `json:"name" form:"name" binding:"required"`
}
