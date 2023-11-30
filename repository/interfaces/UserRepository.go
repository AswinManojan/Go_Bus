package interfaces

import "gobus/entities"

type UserRepository interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
}
