package auth

import (
	"context"
	"go_chat/auth/types"
	db "go_chat/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func FindOneById(id uuid.UUID) (UserModel, error) {
	mongoDocument := db.Client.Database("chat").Collection("users").FindOne(
		context.TODO(),
		bson.D{{Key: "Id", Value: id}},
	)

	var model UserModel
	err := mongoDocument.Decode(&model)

	return model, err
}

func FindOneByLogin(login string) (UserModel, error) {
	mongoDocument := db.Client.Database("chat").Collection("users").FindOne(context.TODO(), bson.D{{Key: "Login", Value: login}})

	var model UserModel
	err := mongoDocument.Decode(&model)

	return model, err
}

// func Me(tokenString string) (bool, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return false, errors.New("Unexpected signing method")
// 		}

// 		return "jwt_secret", nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println(claims["Id"])
// 		return true, nil
// 	} else {
// 		panic(err)
// 	}
// }

func GetUser(id uuid.UUID) (UserDTO, error) {
	entity, err := FindOneById(id)

	if err != nil {
		return NewUserDTO(), err
	}

	return entity.toUserDTO(), nil
}

func Register(input types.RegisterInput) (UserDTO, error) {
	login, password := input.Login, input.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return NewUserDTO(), err
	}

	userModel := UserModel{
		Id:       uuid.New(),
		Login:    login,
		Password: string(hashedPassword),
	}

	_, err = db.Client.Database("chat").Collection("users").InsertOne(context.TODO(), &userModel)

	if err != nil {
		return NewUserDTO(), err
	}

	return userModel.toUserDTO(), nil
}

func Login(input types.LoginInput) (string, error) {
	user, err := FindOneByLogin(input.Login)

	if err != nil {
		return "", err
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return "", err
	}

	// Encode JWT
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		"iat": time.Now().Unix(),
		"sub": user.Id,
	}

	tokenString, err := token.SignedString([]byte("jwt_secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
