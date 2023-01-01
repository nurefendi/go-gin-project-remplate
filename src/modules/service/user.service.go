package service

import (
	"errors"
	"go-gin-template/src/modules/dto/request"
	"go-gin-template/src/modules/entity"
	"go-gin-template/src/modules/repository"
)

type UserService interface {
	RegisterUser(input request.RegisterInput) (error)
	FindById(input string) (entity.SysUser, error)
}

type userService struct {
	sysUserRepository repository.SysUserRepository
}

func NewUserService() *userService {
	return &userService{repository.NewSysUserRepository()}
}

// func (s *userService) RegisterUser(input request.RegisterInput) (error) {
// 	user := entity.User{}

// 	user.FullName = input.Name
// 	user.Email = input.Email
// 	user.Occupation = input.Occupation
// 	user.Avatar = "default.png"
// 	user.CreatedDate = time.Now()
// 	user.CreatedBy = "aplikasi"
// 	user.Role = "user"
// 	user.Token = "token123"
	
// 	passwordHas, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(passwordHas)

// 	_, errr := s.repository.Save(user)

// 	if errr != nil {
// 		return err
// 	}

// 	return nil
// }


func (s *userService) EmailChecker(input request.EmailCheckerInput) (string, error){
	email := input.Email

	user, err := s.sysUserRepository.FindByEmail(email)

	if err != nil {
		return err.Error(), err
	}

	if user.UserId != "" {
		return user.Email, errors.New("email already register")
	}

	return email, nil

}

func (s userService) FindById(userId string) (entity.SysUser, error){
	user, err := s.sysUserRepository.FindById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}