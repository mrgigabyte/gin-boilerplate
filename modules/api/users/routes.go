package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func router(db *gorm.DB, r *gin.Engine) {

	handlers := new(handler)

	//user routes
	r.GET("/users", handlers.getUsers(db))
	r.POST("/users", handlers.postUsers(db))
	r.PUT("/users", handlers.putUsers(db))
	r.DELETE("/users", handlers.deleteUsers(db))
	// r.GET("/user/:id", ctrl.GetUser)

}
