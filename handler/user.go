package handler

import (
	"crowdfounding/helper"
	"crowdfounding/msg"
	"crowdfounding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context){

	var input user.RegisterInput
	var response interface{}

	err := c.ShouldBind(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}

		response = helper.APIResponse(msg.RegisterFailed, http.StatusUnprocessableEntity, msg.Failed, nil, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errr := h.userService.RegisterUser(input)
	if errr != nil {
		response = helper.APIResponse(msg.RegisterFailed, http.StatusBadRequest, msg.Failed, nil, errr.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = helper.APIResponse(msg.RegisterSuccess, http.StatusOK, msg.Success, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Login(c *gin.Context){
	var input user.LoginInput
	var response interface{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"message": errors}

		response = helper.APIResponse(msg.LoginFailed, http.StatusUnprocessableEntity, msg.Failed, nil, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"message": err.Error()}

		response = helper.APIResponse(msg.LoginFailed, http.StatusUnprocessableEntity, msg.Failed, nil, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	forater := user.FormatResponse(loggedUser, "toket")
	response = helper.APIResponse(msg.Success, http.StatusOK, msg.Success, forater,nil)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) EmailChecker(c *gin.Context){
	
}