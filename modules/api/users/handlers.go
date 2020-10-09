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

func postUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /users",
		})
	}

}

func putUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT /users",
		})
	}

}

func deleteUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE /users",
		})
	}

}
