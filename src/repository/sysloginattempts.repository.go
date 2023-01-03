package repository

import (
	"go-gin-template/src/config"
	"go-gin-template/src/entity"
	"time"

	"gorm.io/gorm"
)

type SysLoginAttemptsRepository interface {
	Save(req entity.SysLoginAttempts) error
	CountUserAttemptsByEmail(email string) int16
	RemoveLoginAttemptsByEmail(email string) error
	RemoveLoginAttemptsAfterDay(email string, time time.Time) error
}


type sysLoginAttemptsRepository struct {
	db *gorm.DB
}

func NewSysLoginAttemptsRepository() sysLoginAttemptsRepository {
	return sysLoginAttemptsRepository{config.DB}
}


func (r sysLoginAttemptsRepository) Save(req entity.SysLoginAttempts) error {
	err := r.db.Create(&req).Error
	if err != nil {
		return err
	}
	return nil;
}

func (r sysLoginAttemptsRepository) CountUserAttemptsByEmail(email string) int16 {
	var table entity.SysLoginAttempts
	var count int64
	r.db.Where(" login = ? AND time BETWEEN DATE_SUB(NOW(), INTERVAL 3 MINUTE) and NOW()", email).Find(&table).Count(&count)
	return int16(count);
}

func (r sysLoginAttemptsRepository) RemoveLoginAttemptsByEmail(email string) error {
	var table entity.SysLoginAttempts

	err := r.db.Where("login = ? ", email).Delete(&table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r sysLoginAttemptsRepository) RemoveLoginAttemptsAfterDay(email string, time time.Time) error {
	var table entity.SysLoginAttempts
	err := r.db.Where("login = ? AND time < ? ", email, time).Delete(&table).Error
	
	if err != nil {
		return err
	}
	return nil
}
