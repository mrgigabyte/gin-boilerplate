package api

import (
	"github.com/mrgigabyte/gin-boilerplate/modules/api/users"

	"github.com/gin-gonic/gin"
)

// Init : initializing Api
func Init(r *gin.Engine) {
	users.Init(r)
}
