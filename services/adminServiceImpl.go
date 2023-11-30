package services

import (
	"errors"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	service "gobus/services/interfaces"
	"log"
)

type AdminServiceImpl struct {
	repo repository.AdminRepository
	jwt  *middleware.JwtUtil
}

// AddStation implements interfaces.AdminService.
func (as *AdminServiceImpl) AddStation(station *entities.Stations) (*entities.Stations, error) {
	stations, err := as.repo.AddStation(station)
	if err != nil {
		log.Println("Error Creating station, in adminServiceImpl file")
		return stations, err
	}
	return stations, nil
}

// BlockProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) BlockProvider(id int) (*entities.ServiceProvider, error) {
	provider, err := as.repo.BlockProvider(id)
	if err != nil {
		log.Println("Error Blocking provider, in adminServiceImpl file")
		return provider, err
	}
	return provider, nil
}

// BlockUser implements interfaces.AdminService.
func (as *AdminServiceImpl) BlockUser(id int) (*entities.User, error) {
	user, err := as.repo.BlockUser(id)
	if err != nil {
		log.Println("Error Blocking user, in adminServiceImpl file")
		return user, err
	}
	return user, nil
}

// DeleteProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) DeleteProvider(id int) (*entities.ServiceProvider, error) {
	provider, err := as.repo.DeleteProvider(id)
	if err != nil {
		log.Println("Error Deleting provider, in adminServiceImpl file")
		return provider, err
	}
	return provider, nil
}

// DeleteStation implements interfaces.AdminService.
func (as *AdminServiceImpl) DeleteStation(id int) (*entities.Stations, error) {
	station, err := as.repo.DeleteStation(id)
	if err != nil {
		log.Println("Error Deleting station, in adminServiceImpl file")
		return station, err
	}
	return station, nil
}

// DeleteUser implements interfaces.AdminService.
func (as *AdminServiceImpl) DeleteUser(id int) (*entities.User, error) {
	user, err := as.repo.DeleteUser(id)
	if err != nil {
		log.Println("Error Deleting user, in adminServiceImpl file")
		return user, err
	}
	return user, nil
}

// FindAllProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) FindAllProvider() ([]*entities.ServiceProvider, error) {
	providers, err := as.repo.FindAllProviders()
	if err != nil {
		log.Println("Error finding providers, in adminServiceImpl file")
		return nil, err
	}
	return providers, nil
}

// FindAllStations implements interfaces.AdminService.
func (as *AdminServiceImpl) FindAllStations() ([]*entities.Stations, error) {
	stations, err := as.repo.FindAllStations()
	if err != nil {
		log.Println("Error finding stations, in adminServiceImpl file")
		return nil, err
	}
	return stations, nil
}

// FindAllUsers implements interfaces.AdminService.
func (as *AdminServiceImpl) FindAllUsers() ([]*entities.User, error) {
	users, err := as.repo.FindAllUsers()
	if err != nil {
		log.Println("Error finding users, in adminServiceImpl file")
		return nil, err
	}
	return users, nil
}

// FindProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) FindProvider(id int) (*entities.ServiceProvider, error) {
	provider, err := as.repo.FindProviderById(id)
	if err != nil {
		log.Println("Error finding provider, in adminServiceImpl file")
		return nil, err
	}
	return provider, nil
}

// FindStation implements interfaces.AdminService.
func (as *AdminServiceImpl) FindStation(id int) (*entities.Stations, error) {
	station, err := as.repo.FindStationById(id)
	if err != nil {
		log.Println("Error finding station, in adminServiceImpl file")
		return nil, err
	}
	return station, nil
}

// FindStationByName implements interfaces.AdminService.
func (as *AdminServiceImpl) FindStationByName(name string) (*entities.Stations, error) {
	station, err := as.repo.FindStationByName(name)
	if err != nil {
		log.Println("Error finding station, in adminServiceImpl file")
		return nil, err
	}
	return station, nil
}

// FindUser implements interfaces.AdminService.
func (as *AdminServiceImpl) FindUser(id int) (*entities.User, error) {
	user, err := as.repo.FindUserById(id)
	if err != nil {
		log.Println("Error finding user, in adminServiceImpl file")
		return nil, err
	}
	return user, nil
}

// Login implements interfaces.AdminService.
func (as *AdminServiceImpl) Login(loginRequest *dto.LoginRequest) (string, error) {
	user, err := as.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return "", errors.New("no User exists")
	}
	if user.Password != loginRequest.Password {
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}
	if user.Role != "admin" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}
	token, err := as.jwt.CreateToken(loginRequest.Email, "admin")
	if err != nil {
		return "", errors.New("token NOT generated")
	}

	return token, nil
}

// UnBlockProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) UnBlockProvider(id int) (*entities.ServiceProvider, error) {
	provider, err := as.repo.UnBlockProvider(id)
	if err != nil {
		log.Println("Error UnBlocking provider, in adminServiceImpl file")
		return provider, err
	}
	return provider, nil
}

// UnBlockUser implements interfaces.AdminService.
func (as *AdminServiceImpl) UnBlockUser(id int) (*entities.User, error) {
	user, err := as.repo.UnBlockUser(id)
	if err != nil {
		log.Println("Error UnBlocking user, in adminServiceImpl file")
		return user, err
	}
	return user, nil
}

// UpdateProvider implements interfaces.AdminService.
func (as *AdminServiceImpl) UpdateProvider(id int, provider entities.ServiceProvider) (*entities.ServiceProvider, error) {
	updatedProvider, err := as.repo.EditProvider(id, &provider)
	if err != nil {
		log.Println("Error Updating Station, in adminServiceImpl file")
		return updatedProvider, err
	}
	return updatedProvider, nil
}

// UpdateStation implements interfaces.AdminService.
func (as *AdminServiceImpl) UpdateStation(id int, station entities.Stations) (*entities.Stations, error) {
	updatedStation, err := as.repo.EditStation(id, &station)
	if err != nil {
		log.Println("Error Updating Station, in adminServiceImpl file")
		return updatedStation, err
	}
	return updatedStation, nil
}

// UpdateUser implements interfaces.AdminService.
func (as *AdminServiceImpl) UpdateUser(id int, user entities.User) (*entities.User, error) {
	updatedUser, err := as.repo.EditUser(id, &user)
	if err != nil {
		log.Println("Error Updating user, in adminServiceImpl file")
		return updatedUser, err
	}
	return updatedUser, nil
}

func NewAdminService(repository repository.AdminRepository, jwt *middleware.JwtUtil) service.AdminService {
	return &AdminServiceImpl{
		repo: repository,
		jwt:  jwt,
	}
}
