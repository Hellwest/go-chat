package auth

import (
	"go_chat/auth/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Routes(router *gin.Engine) {
	auth := router.Group("/auth")

	// auth.GET("me", func(c *gin.Context) {

	// })

	auth.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := GetUser(id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	auth.POST("register", func(c *gin.Context) {
		var input types.RegisterInput

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := Register(input)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	auth.POST("login", func(c *gin.Context) {
		var input types.LoginInput

		if err := c.Bind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		jwt, err := Login(input)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"jwt": jwt,
		})
	})
}
