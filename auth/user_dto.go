package auth

import "github.com/google/uuid"

type UserDTO struct {
	Id    uuid.UUID `json:"Id"`
	Login string    `json:"Login"`
}
