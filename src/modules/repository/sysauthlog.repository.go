package repository

import (
	"go-gin-template/src/config"
	"go-gin-template/src/modules/entity"

	"gorm.io/gorm"
)

type SysAuthLogReoisitory interface {
	Save(log entity.SysAuthLog)
}

type sysAuthLogReoisitory struct {
	db *gorm.DB
}

func NewSysAuthLogRepository() sysAuthLogReoisitory {
	return sysAuthLogReoisitory{config.DB}
}


func (r sysAuthLogReoisitory) Save(log entity.SysAuthLog) {
	r.db.Table("sys_auth_log").Create(&log)
}
