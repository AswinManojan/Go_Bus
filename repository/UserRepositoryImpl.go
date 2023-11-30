package repository

import (
	"errors"
	"gobus/entities"
	"gobus/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// FindUserByEmail implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}

	user := &entities.User{}
	result := ur.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		log.Println("User not found in DB")
		return nil, errors.New("user not Found in DB")
	}

	return user, nil
}

// RegisterUser implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Coonecting Database")
	}

	result := ur.DB.Create(&user)
	if result.Error != nil {
		log.Println("Unable to add user, AdminRepositoryImpl package")
		return user, errors.New("user already exists")
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
