package routers

import (
	"go-gin-template/src/middleware"
	controller "go-gin-template/src/modules/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {

	router.POST("/auth/login", controller.NewAuthController().Login)
	router.GET("/auth/me", middleware.AuthMidleware, controller.NewAuthController().Me)
}