package repository

import (
	"go-gin-template/src/config"
	"go-gin-template/src/entity"

	"gorm.io/gorm"
)

type SysAuthTokenRepository interface {
	FindByUserId(userId string) (entity.SysAuthToken, error)
	Save(data entity.SysAuthToken) error
	DeleteByUserId(userId string) error
}

type sysAuthTokenRepository struct {
	db *gorm.DB
}

func NewSysAuthTokenRepository() sysAuthTokenRepository {
	return sysAuthTokenRepository{config.DB}
}


func (r sysAuthTokenRepository) FindByUserId(userId string) (entity.SysAuthToken, error) {
	var sysAuthToken entity.SysAuthToken
	err := r.db.Where("user_id = ? AND validity > NOW() ", userId).Table("sys_auth_token").Find(&sysAuthToken).Error
	if err != nil {
		return sysAuthToken, err
	}

	return sysAuthToken, nil
}

func (r sysAuthTokenRepository) Save(data entity.SysAuthToken) error {
	err := r.db.Table("sys_auth_token").Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r sysAuthTokenRepository) DeleteByUserId(userId string) error {
	var sysAuthToken entity.SysAuthToken
	err := r.db.Table("sys_auth_token").Where("user_id = ?", userId).Delete(&sysAuthToken).Error
	if err != nil {
		return err
	}
	return nil
}
