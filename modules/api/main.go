package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mrgigabyte/gin-boilerplate/modules/api/users"
	"gorm.io/gorm"
)

// Init : initializing Api
func Init(db *gorm.DB, r *gin.Engine) {
	users.Init(db, r)
}
