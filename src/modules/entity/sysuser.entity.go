package entity

import "time"

type SysUser struct {
	UserId       string
	UserName     string
	KataSandi    string
	NamaLengkap  string
	Telepon      string
	Email        string
	JenisKelamin string
	Foto         string
	Status       int8
	LastLogin    time.Time
	RegistedBy   string
	Mdd          time.Time
}

