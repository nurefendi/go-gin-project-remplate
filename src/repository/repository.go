package repository

// import (
// 	"go-gin-template/src/config"
// 	"go-gin-template/src/entity"

// 	"gorm.io/gorm"
// )

// type Repository interface {
// 	Save(user entity.User) (entity.User, error)
// 	FindByEmail(email string) (entity.User, error)
// }

// type repository struct {
// 	db *gorm.DB 
// }

// func NewRepository() *repository {
// 	return &repository{config.DB}
// }

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