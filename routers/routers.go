package routers

import (
	"crowdfounding/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {

	router.POST("/users", controller.NewUserController().RegisterUser)
	router.POST("/login", controller.NewUserController().Login)
}