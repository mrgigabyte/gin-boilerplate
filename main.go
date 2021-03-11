package main

import (
	"github.com/mrgigabyte/gin-boilerplate/db"
	"github.com/mrgigabyte/gin-boilerplate/modules/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.LoadHTMLGlob("templates/**/*")
	db := db.Init()
	api.Init(db, r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
