package main

import (
	"go-gin-template/src/config"
	"go-gin-template/src/routers"
	"os"

	"github.com/gin-gonic/gin"
)

// use godot package to load/read the .env file and
// return the value of the key

func main() {
	config.Env()
	gin.SetMode(os.Getenv("SET_MODE"))
	config.ConnectDatabase()
	// config.InitKafkaClient()

	router := gin.Default()
	api := router.Group(os.Getenv("path"))
	api.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to api v1")
	})
	routers.SetRoutes(api)

	router.Run(os.Getenv("host") + ":" + os.Getenv("port"))

}
