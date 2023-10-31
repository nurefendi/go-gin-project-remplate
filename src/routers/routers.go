package routers

import (
	"go-gin-template/src/helper"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleRouter() {
	log.Info("Start routing")
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})

	// backend web
	SetWebBackendRouters(router.Group("/backend"))
	// backend api
	SetApiBackendRouters(router.Group("/api/backend"))
	
	router.Use(static.Serve("/components", static.LocalFile("./src/templates/assets/component", true)))
	router.HTMLRender = helper.LoadTemplates("./src/templates")
	router.Static("/css", "./src/templates/assets/css")
	router.Static("/js", "./src/templates/assets/js")
	router.Static("/images", "./src/templates/assets/img")
	router.Static("/font", "./src/templates/assets/font")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Info("Routed to port " + port)
	log.Error(router.Run(":" + port))
}


