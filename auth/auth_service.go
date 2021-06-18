package auth

import (
	"context"
	"go_chat/auth/types"
	db "go_chat/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func FindOne(login string) (UserModel, error) {
	mongoDocument := db.Client.Database("chat").Collection("users").FindOne(
		context.TODO(),
		bson.D{{Key: "Login", Value: login}},
	)

	var result UserModel
	if err := mongoDocument.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func Register(input types.RegisterInput) (UserModel, error) {
	uuid := uuid.New()

	userModel := UserModel{
		Id:       uuid,
		Login:    input.Login,
		Password: input.Password,
	}

	_, err := db.Client.Database("chat").Collection("users").InsertOne(context.TODO(), &userModel)

	if err != nil {
		return userModel, err
	}

	return userModel, nil
}
