package mapper

import (
	"go-gin-template/src/dto/response"
	"go-gin-template/src/entity"
	"time"
)


func AuthResponseMapper(e entity.SysUser, token string, exp time.Time) response.AuthResponse{
	return  response.AuthResponse{
		UserId: e.UserId,
		Email: e.Email,
		Token: token,
		Expired: exp,
	}
}