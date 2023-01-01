package repository

import (
	"go-gin-template/src/config"
	"go-gin-template/src/modules/entity"

	"gorm.io/gorm"
)
var sys_user = "sys_user"
type SysUserRepository interface {
	FindByEmail(email string) (entity.SysUser, error)
	FindById(id string) (entity.SysUser, error)
}

type sysUserRepository struct {
	db *gorm.DB
}

func NewSysUserRepository() sysUserRepository {
	return sysUserRepository{config.DB}
}

func (r sysUserRepository) FindByEmail(email string) (entity.SysUser, error) {
	var user entity.SysUser
	err := r.db.Where("email = ?", email).Table(sys_user).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r sysUserRepository) FindById(id string) (entity.SysUser, error){
	var user entity.SysUser
	err := r.db.Where("user_id = ?", id).Table(sys_user).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}