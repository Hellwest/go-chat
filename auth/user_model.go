package auth

import "github.com/google/uuid"

type UserModel struct {
	Id       uuid.UUID `bson:"Id" json:"Id"`
	Login    string    `bson:"Login" json:"Login"`
	Password string    `bson:"Password" json:"Password"`
}

func (m *UserModel) toUserDTO() UserDTO {
	return UserDTO{
		Id:    m.Id,
		Login: m.Login,
	}
}
