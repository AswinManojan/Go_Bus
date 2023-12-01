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

type ProviderServiceImpl struct {
	repo repository.ProviderRepository
	jwt  *middleware.JwtUtil
}

// AddBus implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) AddBus(bus *entities.Buses, email string) (*entities.Buses, error) {
	buses, err := ps.repo.AddBus(bus, email)
	if err != nil {
		log.Println("Error Creating bus, in providerServiceImpl file")
		return buses, err
	}
	return buses, err
}

// AddCoupon implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) AddCoupon(coupon *entities.Coupons) (*entities.Coupons, error) {
	coupons, err := ps.repo.AddCoupon(coupon)
	if err != nil {
		log.Println("Error Creating coupon, in providerServiceImpl file")
		return coupons, err
	}
	return coupons, err
}

// DeleteBus implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) DeleteBus(id int, email string) (*entities.Buses, error) {
	bus, err := ps.repo.DeleteBus(id, email)
	if err != nil {
		log.Println("Error Deleting bus, in providerServiceImpl file")
		return bus, err
	}
	return bus, err
}

// DeleteCoupon implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) DeactivateCoupon(id int) (*entities.Coupons, error) {
	coupon, err := ps.repo.DeactivateCoupon(id)
	if err != nil {
		log.Println("Error Deactivating coupon, in providerServiceImpl file")
		return coupon, err
	}
	return coupon, err
}
func (ps *ProviderServiceImpl) ActivateCoupon(id int) (*entities.Coupons, error) {
	coupon, err := ps.repo.ActivateCoupon(id)
	if err != nil {
		log.Println("Error Activating coupon, in providerServiceImpl file")
		return coupon, err
	}
	return coupon, err
}

// EditBus implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) EditBus(id int, bus *entities.Buses) (*entities.Buses, error) {
	editedBus, err := ps.repo.EditBus(id, bus)
	if err != nil {
		log.Println("Error edit Bus, in providerServiceImpl file")
		return editedBus, err
	}
	return editedBus, err
}

// EditCoupon implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) EditCoupon(id int, coupon *entities.Coupons) (*entities.Coupons, error) {
	editedCoupon, err := ps.repo.EditCoupon(id, coupon)
	if err != nil {
		log.Println("Error edit coupon, in providerServiceImpl file")
		return editedCoupon, err
	}
	return editedCoupon, err
}

// EditProvider implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) EditProvider(email string, provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	editedProvider, err := ps.repo.EditProvider(email, provider)
	if err != nil {
		log.Println("Error edit provider, in providerServiceImpl file")
		return editedProvider, err
	}
	return editedProvider, err
}

// FindAllStations implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindAllStations() ([]*entities.Stations, error) {
	stations, err := ps.repo.FindAllStations()
	if err != nil {
		log.Println("Error finding stations, in providerServiceImpl file")
		return stations, err
	}
	return stations, err
}

// FindBus implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindBus() ([]*entities.Buses, error) {
	buses, err := ps.repo.FindBus()
	if err != nil {
		log.Println("Error finding buses, in providerServiceImpl file")
		return buses, err
	}
	return buses, err
}

// FindBusById implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindBusById(id int) (*entities.Buses, error) {
	bus, err := ps.repo.FindBusById(id)
	if err != nil {
		log.Println("Error finding buses, in providerServiceImpl file")
		return bus, err
	}
	return bus, err
}

// FindCoupon implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindCoupon() ([]*entities.Coupons, error) {
	coupons, err := ps.repo.FindCoupon()
	if err != nil {
		log.Println("Error finding coupon, in providerServiceImpl file")
		return coupons, err
	}
	return coupons, err
}

// FindCouponByCode implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindCouponByCode(code string) (*entities.Coupons, error) {
	coupon, err := ps.repo.FindCouponByCode(code)
	if err != nil {
		log.Println("Error finding coupon, in providerServiceImpl file")
		return coupon, err
	}
	return coupon, err
}

// FindCouponById implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindCouponById(id int) (*entities.Coupons, error) {
	coupon, err := ps.repo.FindCouponById(id)
	if err != nil {
		log.Println("Error finding coupon, in providerServiceImpl file")
		return coupon, err
	}
	return coupon, err
}

// FindProviderByEmail implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindProviderByEmail(email string) (*entities.ServiceProvider, error) {
	provider, err := ps.repo.FindProviderByEmail(email)
	if err != nil {
		log.Println("Error finding provider, in providerServiceImpl file")
		return provider, err
	}
	return provider, err
}

// FindStationById implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindStationById(id int) (*entities.Stations, error) {
	station, err := ps.repo.FindStationById(id)
	if err != nil {
		log.Println("Error finding station, in providerServiceImpl file")
		return station, err
	}
	return station, err
}

// FindStationByName implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) FindStationByName(name string) (*entities.Stations, error) {
	station, err := ps.repo.FindStationByName(name)
	if err != nil {
		log.Println("Error finding station, in providerServiceImpl file")
		return station, err
	}
	return station, err
}

// Login implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) Login(loginRequest *dto.LoginRequest) (string, error) {

	foundProvider, err := ps.repo.FindProviderByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No Provider EXISTS, in adminService file")
		return "", errors.New("no Provider exists")
	}
	dbHashedPassword := foundProvider.Password // Replace with the actual hashed password.

	enteredPassword := loginRequest.Password // Replace with the password entered by the user during login.

	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(enteredPassword)); err != nil {
		// Passwords match. Allow the user to log in.
		log.Println("Password Mismatch, in adminService file")
		return "", errors.New("password Mismatch")
	}
	if foundProvider.Role != "provider" {
		log.Println("Unauthorized, in adminService file")
		return "", errors.New("unauthorized access")
	}
	if foundProvider.IsLocked {
		log.Println("User locked by Admin,Contact admin to unlock the account--- in adminService file")
		return "", errors.New("locked account")
	}
	token, err := ps.jwt.CreateToken(loginRequest.Email, "provider")
	if err != nil {
		return "", errors.New("token NOT generated")
	}

	return token, nil
}

// RegisterProvider implements interfaces.ProviderService.
func (ps *ProviderServiceImpl) RegisterProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if hashedPassword, err := utils.HashPassword(provider.Password); err != nil {
		log.Println("Unable to hash password")
		return nil, err
	} else {
		provider.Password = hashedPassword
	}
	regProvider, err := ps.repo.RegisterProvider(provider)
	if err != nil {
		log.Println("Provider not added, adminService file")
		return regProvider, err
	}
	return regProvider, err
}

func NewProviderService(repo repository.ProviderRepository, jwt *middleware.JwtUtil) interfaces.ProviderService {
	return &ProviderServiceImpl{
		repo: repo,
		jwt:  jwt,
	}
}
