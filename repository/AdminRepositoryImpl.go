package repository

import (
	"errors"
	"gobus/entities"
	"gobus/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

// FindUserByEmail implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindUserByEmail(mail string) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user := &entities.User{}
	result := ar.DB.Where("email=?", mail).First(user)
	if result.Error != nil {
		log.Println("User doesn't exist")
		return nil, errors.New("no User found with this name")
	}
	return user, nil
}

// BlockProvider implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) BlockProvider(id int) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider, err := ar.FindProviderById(id)
	if err != nil {
		log.Println("provider not found")
		return nil, errors.New("provider not found")
	}
	provider.IsLocked = true
	result := ar.DB.Save(provider)
	if result.Error != nil {
		log.Println("Unable to block provider")
		return nil, errors.New("unable to block provider")
	}
	return provider, nil
}

// UnBlockProvider implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UnBlockProvider(id int) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider, err := ar.FindProviderById(id)
	if err != nil {
		log.Println("provider not found")
		return nil, errors.New("provider not found")
	}
	provider.IsLocked = false
	result := ar.DB.Save(provider)
	if result.Error != nil {
		log.Println("Unable to unblock provider")
		return nil, errors.New("unable to unblock provider")
	}
	return provider, nil
}

// FindStationByName implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindStationByName(name string) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station := &entities.Stations{}
	result := ar.DB.Where("name=?", name).First(station)
	if result.Error != nil {
		log.Println("Station doesn't exist")
		return nil, errors.New("no station found with this name")
	}
	return station, nil
}

// BlockUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) BlockUser(id int) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ar.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	user.IsLocked = true
	result := ar.DB.Save(user)
	if result.Error != nil {
		log.Println("Unable to block user")
		return nil, errors.New("unable to block user")
	}
	return user, nil
}

// UnBlockUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UnBlockUser(id int) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ar.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	user.IsLocked = false
	result := ar.DB.Save(user)
	if result.Error != nil {
		log.Println("Unable to block user")
		return nil, errors.New("unable to block user")
	}
	return user, nil
}

// AddStation implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) AddStation(station *entities.Stations) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}

	_, err := ar.FindStationByName(station.StationName)
	if err == nil {
		log.Println("STATION ALREADY EXISTS")
		return nil, errors.New("STATION exists in db")
	}

	result := ar.DB.Create(&station)
	if result.Error != nil {
		log.Println("Unable to add user, AdminRepositoryImpl package")
		return nil, errors.New("user not added to db")
	}
	return station, nil
}

// DeleteProvider implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) DeleteProvider(id int) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider, err := ar.FindProviderById(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ar.DB.Delete(provider).Where("id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return provider, nil
}

// DeleteStation implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) DeleteStation(id int) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station, err := ar.FindStationById(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ar.DB.Delete(station).Where("id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return station, nil
}

// DeleteUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) DeleteUser(id int) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ar.FindUserById(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ar.DB.Delete(user).Where("id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return user, nil
}

// EditProvider implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) EditProvider(id int, provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundProvider, err := ar.FindProviderById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if foundProvider.Address != "" {
		foundProvider.Address = provider.Address
	}
	if foundProvider.CompanyName != "" {
		foundProvider.CompanyName = provider.CompanyName
	}
	if foundProvider.Email != "" {
		foundProvider.Email = provider.Email
	}
	if foundProvider.PhoneNumber != "" {
		foundProvider.PhoneNumber = provider.PhoneNumber
	}
	// if foundProvider.BusCount != provider.BusCount {
	// 	foundProvider.BusCount = provider.BusCount
	// }
	if foundProvider.Password != "" {
		foundProvider.Password = provider.Address
	}
	result := ar.DB.Save(&foundProvider)
	if result.Error != nil {
		log.Println("User Not Updated maybe the same email already present, AdminRepositoryImpl package")
		return nil, errors.New("user not updated")
	}
	return foundProvider, nil
}

// EditStation implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) EditStation(id int, station *entities.Stations) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundStation, err := ar.FindStationById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if foundStation.Location != "" {
		foundStation.Location = station.Location
	}
	if foundStation.StationCode != "" {
		foundStation.StationCode = station.StationCode
	}
	if foundStation.StationName != "" {
		foundStation.StationName = station.StationName
	}
	result := ar.DB.Save(&foundStation)
	if result.Error != nil {
		log.Println("User Not Updated maybe the same email already present, AdminRepositoryImpl package")
		return nil, errors.New("user not updated")
	}
	return foundStation, nil
}

// EditUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) EditUser(id int, user *entities.User) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundUser, err := ar.FindUserById(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if foundUser.Email != "" {
		foundUser.Email = user.Email
	}
	if foundUser.UserName != "" {
		foundUser.UserName = user.UserName
	}
	if foundUser.DOB != "" {
		foundUser.DOB = user.DOB
	}
	if foundUser.Gender != "" {
		foundUser.Gender = user.Gender
	}
	if foundUser.Password != "" {
		foundUser.Password = user.Password
	}
	if foundUser.PhoneNumber != "" {
		foundUser.PhoneNumber = user.PhoneNumber
	}
	foundUser.IsLocked = true
	result := ar.DB.Save(&foundUser)
	if result.Error != nil {
		log.Println("User Not Updated maybe the same email already present, AdminRepositoryImpl package")
		return nil, errors.New("user not updated")
	}
	return foundUser, nil
}

// FindAllProviders implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindAllProviders() ([]*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	providers := []*entities.ServiceProvider{}
	result := ar.DB.Find(&providers)
	if result.Error != nil {
		return nil, result.Error
	}
	return providers, nil
}

// FindAllStations implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindAllStations() ([]*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	stations := []*entities.Stations{}
	result := ar.DB.Find(&stations)
	if result.Error != nil {
		return nil, result.Error
	}
	return stations, nil
}

// FindAllUsers implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindAllUsers() ([]*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	users := []*entities.User{}
	result := ar.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindProviderById implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindProviderById(id int) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider := &entities.ServiceProvider{}
	result := ar.DB.First(provider, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}

// FindStationById implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindStationById(id int) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station := &entities.Stations{}
	result := ar.DB.First(station, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return station, nil
}

// FindUserById implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindUserById(id int) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user := &entities.User{}
	result := ar.DB.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{
		DB: db,
	}
}
