package api

import (
	"gin-boilerplate/modules/api/users"

	"github.com/gin-gonic/gin"
)

// Init : initializing Api
func Init(r *gin.Engine) {
	users.Init(r)
}
