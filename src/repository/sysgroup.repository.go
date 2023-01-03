package repository

import (
	"errors"
	"go-gin-template/src/config"
	"go-gin-template/src/entity"

	"gorm.io/gorm"
)

type SysGroupRepository interface {
	FindById(id string) (entity.SysGroup, error)
}

type sysGroupRepository struct {
	db *gorm.DB 
}

func NewSysGroupRepository() *sysGroupRepository {
	return &sysGroupRepository{config.DB}
}

func (r sysGroupRepository) FindById(id string) (entity.SysGroup, error){
	var group entity.SysGroup
	err := r.db.Where("group_id = ? ", id).Table("sys_group").Find(&group).Error
	if err != nil {
		return group, err
	}

	if group.GroupId == "" {
		return group, errors.New("group not found")
	}
	
	return group, nil
}

// func (r *repository) Save(user entity.User) (entity.User, error) {
// 	err := r.db.Create(&user).Error

// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (r *repository) FindByEmail(email string) (entity.User, error){
// 	var user entity.User

// 	err := r.db.Where("email = ?", email).Find(&user).Error
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }