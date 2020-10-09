package users

import (
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {

	//user routes
	r.GET("/users", getUsers())
	r.POST("/users", postUsers())
	r.PUT("/users", putUsers())
	r.DELETE("/users", deleteUsers())
	// r.GET("/user/:id", ctrl.GetUser)

}
