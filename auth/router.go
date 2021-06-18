package auth

import (
	"go_chat/auth/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	auth := router.Group("/auth")

	auth.GET("/:login", func(c *gin.Context) {
		login := c.Param("login")

		user, err := FindOne(login)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	auth.POST("register", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")

		user, err := Register(types.RegisterInput{Login: login, Password: password})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})
}
