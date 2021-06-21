package auth

import (
	"context"
	"fmt"
	"go_chat/auth/types"
	db "go_chat/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func FindOne(login string) (UserDTO, error) {
	mongoDocument := db.Client.Database("chat").Collection("users").FindOne(
		context.TODO(),
		bson.D{{Key: "Login", Value: login}},
	)

	var result UserModel
	if err := mongoDocument.Decode(&result); err != nil {
		return result.toUserType(), err
	}

	return result.toUserType(), nil
}

func Register(input types.RegisterInput) (UserDTO, error) {
	uuid := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	if err != nil {
		panic(err)
	}

	fmt.Println("The hash:", string(hashedPassword))

	userModel := UserModel{
		Id:       uuid,
		Login:    input.Login,
		Password: input.Password,
	}

	_, err = db.Client.Database("chat").Collection("users").InsertOne(context.TODO(), &userModel)

	if err != nil {
		return userModel.toUserType(), err
	}

	return userModel.toUserType(), nil
}
