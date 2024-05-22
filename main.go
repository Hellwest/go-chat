package main

import (
	"context"
	"fmt"
	"go_chat/auth"
	"go_chat/rooms"
	"os"

	db "go_chat/database"

	"github.com/gin-gonic/gin"
)

func getPortFromEnv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return fmt.Sprintf(":%s", value)
}

func main() {
	port := getPortFromEnv("APP_PORT", ":5000")

	router := gin.Default()
	auth.Routes(router)
	rooms.Routes(router)

	defer db.Client.Disconnect(context.TODO())
	router.Run(port)
}
