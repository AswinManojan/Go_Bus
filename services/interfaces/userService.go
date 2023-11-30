package interfaces

import (
	"gobus/dto"
	"gobus/entities"
)

type UserService interface {
	Login(loginRequest *dto.LoginRequest) (string, error)
	RegisterUser(user *entities.User) (*entities.User, error)
}
