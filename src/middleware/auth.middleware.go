package middleware

import (
	"errors"
	"go-gin-template/src/constant"
	"go-gin-template/src/helper"
	"go-gin-template/src/modules/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var CURRENT_USER helper.JwtPayload

func AuthMidleware(c *gin.Context) {
	result := helper.NewRestResult()
	TOKEN := ""
	tokenCookie, err := c.Cookie("token")
	if err != nil {
		tokenHeader, errr := getTokenFromHeader(c)

		if errr != nil {
			result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.Unautorized)
			result.SetErrors(gin.H{"message": constant.Unautorized})
			c.AbortWithStatusJSON(http.StatusUnauthorized, result)
			return
		}
		TOKEN = tokenHeader
	} else {
		TOKEN = tokenCookie
	}

	jwtService := service.NewJwtService()
	tokenValid, err := jwtService.ValidatedToken(TOKEN)
	if err != nil {
		result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.Unautorized)
		result.SetErrors(gin.H{"message": constant.Unautorized})
		c.AbortWithStatusJSON(http.StatusUnauthorized, result)
		return
	}

	claim, ok := tokenValid.Claims.(jwt.MapClaims)

	if !ok || !tokenValid.Valid {
		result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.Unautorized)
		result.SetErrors(gin.H{"message": constant.Unautorized})
		c.AbortWithStatusJSON(http.StatusUnauthorized, result)
		return
	}

	currentLoggin := helper.JwtPayload{
		UserId: claim["userId"].(string),
		Email:  claim["email"].(string),
		Name: claim["name"].(string),
		GroupId: claim["groupId"].(string),
		GroupName: claim["groupName"].(string),
	}

	userService := service.NewUserService()
	dataUser, err := userService.FindById(currentLoggin.UserId)
	if err != nil {
		result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.Unautorized)
		result.SetErrors(gin.H{"message": constant.Unautorized})
		c.AbortWithStatusJSON(http.StatusUnauthorized, result)
		return
	}

	if dataUser.UserId == "" {
		result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.Unautorized)
		result.SetErrors(gin.H{"message": constant.Unautorized})
		c.AbortWithStatusJSON(http.StatusUnauthorized, result)
		return
	}

	if dataUser.Status != 1 {
		result.SetMeta(http.StatusUnauthorized, constant.Failed, constant.UserNotActive)
		result.SetErrors(gin.H{"message": constant.UserNotActive})
		c.AbortWithStatusJSON(http.StatusUnauthorized, result)
		return
	}

	CURRENT_USER = currentLoggin
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	tokenHeader := c.GetHeader("Authorization")

	if !strings.Contains(tokenHeader, "Bearer") {
		return tokenHeader, errors.New("Unautorized")
	}

	arrayToken := strings.Split(tokenHeader, " ")

	if len(arrayToken) != 2 {
		return tokenHeader, errors.New("token is not valid")
	}

	return arrayToken[1], nil

}