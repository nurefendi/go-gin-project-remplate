package entity

import "time"

type SysGroup struct {
	GroupId         string
	GroupName       string
	GroupDesc       string
	GroupPortal     string
	GroupRestricted int
	Mdb             int
	Mdd             time.Time
}
