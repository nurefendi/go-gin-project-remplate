package service

import (
	"errors"
	"go-gin-template/src/dto/request"
	"go-gin-template/src/entity"
	"go-gin-template/src/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input request.RegisterInput) (error)
	Login(input request.LoginInput) (entity.User, error)
}

type userService struct {
	repository repository.Repository
}

func NewUserService() *userService {
	return &userService{repository.NewRepository()}
}

func (s *userService) RegisterUser(input request.RegisterInput) (error) {
	user := entity.User{}

	user.FullName = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Avatar = "default.png"
	user.CreatedDate = time.Now()
	user.CreatedBy = "aplikasi"
	user.Role = "user"
	user.Token = "token123"
	
	passwordHas, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHas)

	_, errr := s.repository.Save(user)

	if errr != nil {
		return err
	}

	return nil
}

func (s *userService) Login(input request.LoginInput) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.IdUser == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) EmailChecker(input request.EmailCheckerInput) (string, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return err.Error(), err
	}

	if user.IdUser != 0 {
		return user.Email, errors.New("email already register")
	}

	return "", nil

}