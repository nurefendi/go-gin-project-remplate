package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (error) {
	user := User{}

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

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.IdUser == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}