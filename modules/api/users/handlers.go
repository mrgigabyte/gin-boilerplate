package users

import (
	"github.com/gin-gonic/gin"
)

func getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET /users",
		})
	}

}
