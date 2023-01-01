package repository

import (
	"errors"
	"go-gin-template/src/config"
	"go-gin-template/src/modules/entity"

	"gorm.io/gorm"
)

type SysUserGroupRepository interface {
	FindByUserId(userId string) (entity.SysUserGroup, error)
}

type sysUserGroupRepository struct {
	db *gorm.DB 
}

func NewSysUserGroupRepository() *sysUserGroupRepository {
	return &sysUserGroupRepository{config.DB}
}

func (r sysUserGroupRepository) FindByUserId(userId string) (entity.SysUserGroup, error){
	var userGroup entity.SysUserGroup
	err := r.db.Where("user_id = ? ", userId).Table("sys_user_group").Find(&userGroup).Error
	if err != nil {
		return userGroup, err
	}
	if userGroup.GroupId == "" {
		return userGroup, errors.New("group not found")
	}
	return userGroup, nil
}
