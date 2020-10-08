package users

import (
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {

	//user routes
	r.GET("/users", getUsers())
	// r.GET("/user/:id", ctrl.GetUser)

}
