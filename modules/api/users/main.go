package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Init : initializing Users
func Init(db *gorm.DB, r *gin.Engine) {
	router(db, r)
}
