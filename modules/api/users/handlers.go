package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrgigabyte/gin-boilerplate/modules/models"
	"gorm.io/gorm"
)

type handler struct{}

var userModel = new(models.User)

func (h handler) getUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := db.Find(&userModel).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, userModel)
		}

		// c.JSON(200, gin.H{
		// 	"message": "GET /users",
		// })
	}

}

func (h handler) postUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST /users",
		})
	}

}

func (h handler) putUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT /users",
		})
	}

}

func (h handler) deleteUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE /users",
		})
	}

}
