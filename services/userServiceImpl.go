package services

import (
	"errors"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	"log"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	jwt  *middleware.JwtUtil
}

func (usi *UserServiceImpl) Login(login *dto.LoginRequest) (string, error) {
	user, err := usi.repo.FindUserByEmail(login.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return "", errors.New("no User exists")
	}

	if user.Password != login.Password {
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}

	if user.Role != "user" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}

	token, err := usi.jwt.CreateToken(login.Email, "user")
	if err != nil {
		return "", errors.New("token NOT generated")
	}
	return token, nil
}

func (usi *UserServiceImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	user, err := usi.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return user, err
	}
	return user, err
}

func NewUserService(repo repository.UserRepository, jwt *middleware.JwtUtil) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
		jwt:  jwt,
	}
}
