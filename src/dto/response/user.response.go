package response

import "go-gin-template/src/entity"

type UserResponse struct {
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatResponse(user entity.User, token string) UserResponse {
	formater := UserResponse{
		UserId:     user.IdUser,
		Name:       user.FullName,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formater
}