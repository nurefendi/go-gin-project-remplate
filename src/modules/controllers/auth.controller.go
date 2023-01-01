package controller

import (
	"go-gin-template/src/constant"
	"go-gin-template/src/helper"
	"go-gin-template/src/middleware"
	"go-gin-template/src/modules/dto/request"
	"go-gin-template/src/modules/entity"
	"go-gin-template/src/modules/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController() AuthController {
	return AuthController{service.NewAuthService()}
}

func (ctrl AuthController) Login(c *gin.Context){
	var input request.LoginRequest
	result := helper.NewRestResult()
	err := c.ShouldBindJSON(&input)
	loginAttempts := entity.SysLoginAttempts{
		Time: time.Now(),
		IpAddress: c.ClientIP(),
		Login: input.Email,
		LoginId: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	if ctrl.authService.CheckLoginAttempts(input.Email, 3) {
		result.SetMeta(http.StatusUnprocessableEntity, constant.Failed, constant.MaxLogin)
		result.SetErrors(gin.H{"message": constant.MaxLogin})
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}

	if err != nil {
		errors := helper.FormatValidationError(err)
		result.SetMeta(http.StatusBadRequest, constant.Failed, constant.LoginFailed)
		result.SetErrors(gin.H{"message": errors})
		c.JSON(http.StatusBadRequest, result)

		ctrl.authService.InsertLogginAttempts(loginAttempts)
		return
	}

	userdata, err := ctrl.authService.DoLogin(input)

	if err != nil {
		result.SetMeta(http.StatusBadRequest, constant.Failed, constant.LoginFailed)
		result.SetErrors(gin.H{"message": err.Error()})
		c.JSON(http.StatusBadRequest, result)

		ctrl.authService.InsertLogginAttempts(loginAttempts)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    userdata.Token,
		MaxAge:   int(userdata.Expired.Sub(time.Now().Local()).Seconds()),
		Path:     "/",
		Secure: false,
		HttpOnly: true,
	})

	ctrl.authService.LogLogin(entity.SysAuthLog{
		LogId: strconv.FormatInt(time.Now().UnixMilli(), 10),
		User: userdata.UserId,
		WaktuLogin: time.Now(),
		Ip: c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
		Keterangan: "Login Berhasil",
		Status: "success",
	})

	result.SetMeta(http.StatusOK, constant.Success, constant.Success)
	result.SetData(userdata)
	c.JSON(http.StatusOK, result)
}

func (ctrl AuthController) Me(c *gin.Context){
	result := helper.NewRestResult()
	result.SetMeta(http.StatusOK, constant.Success, constant.Success)
	result.SetData(middleware.CURRENT_USER)
	c.JSON(http.StatusOK, result)
}