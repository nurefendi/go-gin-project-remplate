package controller

import (
	"crowdfounding/constant"
	"crowdfounding/dto/request"
	"crowdfounding/dto/response"
	"crowdfounding/helper"
	"crowdfounding/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController() *UserController {
	return &UserController{service.NewUserService()}
}

func (h *UserController) RegisterUser(c *gin.Context) {

	var input request.RegisterInput
	response := helper.NewRestResult()

	err := c.ShouldBind(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		response.SetMeta(http.StatusUnprocessableEntity, constant.Failed, constant.RegisterFailed)
		response.SetErrors(gin.H{"message": errors})

		c.JSON(http.StatusBadRequest, response)
		return
	}

	errr := h.userService.RegisterUser(input)
	if errr != nil {
		response.SetMeta(http.StatusUnprocessableEntity, constant.Failed, constant.RegisterFailed)
		response.SetErrors(errr.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.SetMeta(http.StatusOK, constant.Success, constant.RegisterSuccess)
	c.JSON(http.StatusOK, response)
}

func (h *UserController) Login(c *gin.Context) {
	var input request.LoginInput
	result := helper.NewRestResult()

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		result.SetMeta(http.StatusUnprocessableEntity, constant.Failed, constant.LoginFailed)
		result.SetErrors(gin.H{"message": errors})

		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	loggedUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"message": err.Error()}

		result.SetMeta(http.StatusUnprocessableEntity, constant.Failed, constant.LoginFailed)
		result.SetErrors(errorMessage)

		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	result.SetMeta(http.StatusOK, constant.Success, constant.Success)
	result.SetData(response.FormatResponse(loggedUser, "toket"))
	c.JSON(http.StatusOK, result)
}

func (h *UserController) EmailChecker(c *gin.Context) {

}
