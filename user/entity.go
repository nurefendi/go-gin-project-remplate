package user

import "time"

type User struct {
	IdUser       int
	Email        string
	FullName     string
	Occupation   string
	Password     string
	Avatar       string
	Role         string
	Token        string
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedDate time.Time
	ModifiedBy   string
}
