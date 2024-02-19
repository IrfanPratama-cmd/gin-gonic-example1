package main

import (
	"gin-socmed/config"
	"gin-socmed/migration"
	"gin-socmed/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadCOnfig()
	config.LoadDB()
	migration.RunMigration()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.PostRouter(api)

	r.Run() // listen and serve on 0.0.0.0:8080
}
