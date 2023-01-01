package entity

import "time"

type SysLoginAttempts struct {
	LoginId   string
	IpAddress string
	Login     string
	Time      time.Time
}
