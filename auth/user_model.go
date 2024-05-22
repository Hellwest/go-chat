package auth

import "github.com/google/uuid"

type UserModel struct {
	Id       uuid.UUID `bson:"id" json:"id"`
	Login    string    `bson:"login" json:"login"`
	Password string    `bson:"password" json:"password"`
}

func (m *UserModel) toUserDTO() UserDTO {
	return UserDTO{
		Id:    m.Id,
		Login: m.Login,
	}
}
