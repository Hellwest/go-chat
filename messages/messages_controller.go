package messages

import (
	"go_chat/messages/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	messages := router.Group("/messages")

	messages.POST("/", func(c *gin.Context) {
		var input types.SendMessageInput

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		message, err := SendMessage(input)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": message})
	})
}
