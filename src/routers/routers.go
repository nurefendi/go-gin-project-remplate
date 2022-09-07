package routers

import (
	controller "go-gin-template/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {

	router.POST("/users", controller.NewUserController().RegisterUser)
	router.POST("/login", controller.NewUserController().Login)
}