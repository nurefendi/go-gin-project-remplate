package routers

import (
	"go-gin-template/src/controller"

	"github.com/gin-gonic/gin"
)

func SetApiBackendRouters(router *gin.RouterGroup) {
	router.GET("/portal", controller.LisPortal)
	router.POST("/portal", controller.StorePortal)
	router.DELETE("/portal/:id", controller.DeletePortal)
	router.PUT("/portal/:id", controller.UpdatePortal)
}