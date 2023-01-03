package entity

import "time"

type SysAuthLog struct {
	LogId string
	User string
	WaktuLogin time.Time
	Ip string
	UserAgent string
	Keterangan string
	Status string
}

