package users

import (
	"github.com/gin-gonic/gin"
)

// Init : initializing Users
func Init(r *gin.Engine) {
	router(r)
}
