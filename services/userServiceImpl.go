package services

import (
	"errors"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	"gobus/services/interfaces"
	"gobus/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
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
	dbHashedPassword := user.Password // Replace with the actual hashed password.

	enteredPassword := login.Password // Replace with the password entered by the user during login.

	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(enteredPassword)); err != nil {
		// Passwords match. Allow the user to log in.
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}
	if user.Role != "user" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}
	if user.IsLocked {
		log.Println("User locked by Admin,Contact admin to unlock the account--- in adminService file")
		return "", errors.New("locked account")
	}

	token, err := usi.jwt.CreateToken(login.Email, "user")
	if err != nil {
		return "", errors.New("token NOT generated")
	}
	return token, nil
}

func (usi *UserServiceImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		log.Println("Unable to hash password")
		return nil, err
	} else {
		user.Password = hashedPassword
	}
	user, err := usi.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return user, err
	}
	return user, err
}

func NewUserService(repo repository.UserRepository, jwt *middleware.JwtUtil) interfaces.UserService {
	return &UserServiceImpl{
		repo: repo,
		jwt:  jwt,
	}
}
