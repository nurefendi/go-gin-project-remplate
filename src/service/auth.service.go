package service

import (
	"errors"
	"go-gin-template/src/helper"
	"go-gin-template/src/dto/request"
	"go-gin-template/src/dto/response"
	"go-gin-template/src/entity"
	"go-gin-template/src/mapper"
	"go-gin-template/src/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	DoLogin(input request.LoginRequest) (response.AuthResponse, error)
	InsertLogginAttempts(data entity.SysLoginAttempts) error
	CheckLoginAttempts(email string, maxTry int16) bool
	CheckLogin(userId string) (response.AuthResponse, error)
	LogLogin(log entity.SysAuthLog)
	LogOut(userId string) error
}

type authService struct {
	sysUserRepository          repository.SysUserRepository
	loginAttemptsRepository repository.SysLoginAttemptsRepository
	sysGroupRepository repository.SysGroupRepository
	sysUserGroupRepository repository.SysUserGroupRepository
	sysAuthTokenRepository repository.SysAuthTokenRepository
	sysAuthLogRepository repository.SysAuthLogReoisitory
	jwtService              jwtService
}

func NewAuthService() authService {
	return authService{
		repository.NewSysUserRepository(),
		repository.NewSysLoginAttemptsRepository(),
		repository.NewSysGroupRepository(),
		repository.NewSysUserGroupRepository(),
		repository.NewSysAuthTokenRepository(),
		repository.NewSysAuthLogRepository(),
		NewJwtService(),
	}
}


func (as authService) DoLogin(request request.LoginRequest) (response.AuthResponse, error) {
	email := request.Email
	password := request.Password

	user, err := as.sysUserRepository.FindByEmail(email)
	if err != nil {
		return response.AuthResponse{}, err
	}

	if user.Email == "" || user.KataSandi == "" {
		return response.AuthResponse{}, errors.New("email is not registed")
	}

	// create validasi passeord
	if !checkPasswordHash(password, user.KataSandi) {
		return response.AuthResponse{}, errors.New("wrong Password")
	}

	userGroup, err := as.sysUserGroupRepository.FindByUserId(user.UserId)
	if err != nil {
		return response.AuthResponse{}, err
	}

	group, err := as.sysGroupRepository.FindById(userGroup.GroupId)
	if err != nil {
		return response.AuthResponse{}, err
	}

	authResponse,_ := as.CheckLogin(user.UserId)
	if authResponse.Token != "" {
		return authResponse, nil
	}

	now := time.Now().Local()
    yyyy, mm, dd := now.Date()
    timeExp := time.Date(yyyy, mm, dd+1, 23, 59, 59, 0, now.Location())
	jwtPayload := helper.JwtPayload{
		UserId:    user.UserId,
		Email:     user.Email,
		Name:      user.NamaLengkap,
		GroupId:   group.GroupId,
		GroupName: group.GroupName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeExp.Unix(),
			Issuer:    "bearer token",
		},
	}

	token, err := as.jwtService.GenerateToken(jwtPayload)
	if err != nil {
		return response.AuthResponse{}, err
	}
	exp := timeExp

	saveToken := entity.SysAuthToken{
		Token: token,
		Validity: exp,
		UserId: user.UserId,
		GroupId: group.GroupId,
		Ctb: "SYSTEM",
		Ctd: time.Now(),
	}
	as.sysAuthTokenRepository.Save(saveToken)
	return mapper.AuthResponseMapper(user, token, exp), nil

}

func (sv authService) CheckLoginAttempts(email string, maxTry int16) bool {
	totalAttempts := sv.loginAttemptsRepository.CountUserAttemptsByEmail(email)
	return maxTry <= totalAttempts
}

func (sv authService) InsertLogginAttempts(input entity.SysLoginAttempts) error {
	err := sv.loginAttemptsRepository.Save(input)
	if err != nil {
		return err
	}

	return nil
}

func (sv authService) CheckLogin(userId string) (response.AuthResponse, error) {
	authToken, err := sv.sysAuthTokenRepository.FindByUserId(userId)
	if err != nil {
		return response.AuthResponse{}, err
	}

	user, err := sv.sysUserRepository.FindById(userId)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return mapper.AuthResponseMapper(user, authToken.Token, authToken.Validity), nil
}

func (sv authService) LogLogin(log entity.SysAuthLog){
	sv.sysAuthLogRepository.Save(log)
}

func (sv authService) LogOut(userId string) error{
	err := sv.sysAuthTokenRepository.DeleteByUserId(userId)
	if err != nil {
		return err
	}
	return nil
}

// func hashPassword(password string) (string, error) {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//     return string(bytes), err
// }

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
