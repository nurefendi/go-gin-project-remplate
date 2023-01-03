package service

import (
	"errors"
	"go-gin-template/src/helper"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(claim helper.JwtPayload) (string, error)
	ValidatedToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewJwtService() jwtService {
	return jwtService{}
}

func (s *jwtService) GenerateToken(claim helper.JwtPayload) (string, error){
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := jwt.SignedString(SECRET_KEY)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s jwtService) ValidatedToken(token string) (*jwt.Token, error){
	jwt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return jwt, err
	}

	return jwt, nil

}