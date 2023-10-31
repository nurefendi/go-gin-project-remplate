package routers

import (
	"go-gin-template/src/controller"

	"github.com/gin-gonic/gin"
)

func SetApiBackendRouters(router *gin.RouterGroup) {
	router.GET("/portal", controller.LisPortal)
}