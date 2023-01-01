package entity

import "time"

type SysAuthToken struct {
	Id       int
	Token    string
	Validity time.Time
	UserId   string
	GroupId  string
	Ctb      string
	Ctd      time.Time
}
