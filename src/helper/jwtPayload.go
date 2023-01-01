package helper

import "github.com/dgrijalva/jwt-go"

type JwtPayload struct {
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	GroupId   string `json:"groupId"`
	GroupName string `json:"groupName"`
	jwt.StandardClaims
}