package types

type RegisterInput struct {
	Login    string `json:"login" form:"login" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
