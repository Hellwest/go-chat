package rooms

import (
	"go_chat/rooms/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	rooms := router.Group("/rooms")

	rooms.POST("/", func(c *gin.Context) {
		var input types.CreateRoomInput

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		room, err := CreateRoom(input)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"room": room})
	})
}
