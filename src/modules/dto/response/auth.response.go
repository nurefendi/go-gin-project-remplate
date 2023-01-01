package response

import "time"

type AuthResponse struct {
	UserId    string  `json:"userId"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	Expired time.Time `json:"expired"`
}