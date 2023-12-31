package repository

import (
	"errors"
	"fmt"
	"gobus/dto"
	"gobus/entities"
	"gobus/repository/interfaces"
	"log"
	"time"

	"gorm.io/gorm"
)

// AdminRepositoryImpl struct is used to define Admin Repository implementation.
type AdminRepositoryImpl struct {
	DB *gorm.DB
}

// GetRouteByBus implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) GetRouteByBus(scheduleID int) (*entities.Schedule, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	schedule := &entities.Schedule{}
	result := ar.DB.Where("schedule_id=?", scheduleID).Find(schedule)
	if result.Error != nil {
		log.Println("Unable to fetch the buses")
		return nil, result.Error
	}
	return schedule, nil
}

// ViewBookingsToBeCancelled implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) ViewBookingsToBeCancelled(busID int, day string) ([]*entities.Booking, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	str := "Success"
	bookings := []*entities.Booking{}
	result := ar.DB.Where("bus_id=? AND booking_date=? AND status=?", busID, day, str).Find(&bookings)
	if result.Error != nil {
		log.Println("Unable to fetch the buses")
		return nil, result.Error
	}
	return bookings, nil
}

// UpdateBooking implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UpdateBooking(booking *entities.Booking) (*entities.Booking, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ar.DB.Save(booking)
	if result.Error != nil {
		return nil, result.Error
	}
	return booking, nil
}

// UpdateChart implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UpdateChart(chart *entities.BusSchedule) (*entities.BusSchedule, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ar.DB.Save(chart)
	if result.Error != nil {
		return nil, result.Error
	}
	return chart, nil
}

// UpdateProvider implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UpdateProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ar.DB.Save(provider)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}

// UpdateUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) UpdateUser(user *entities.User) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ar.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetBusInfo implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) GetBusInfo(id int) (*entities.Buses, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bus := &entities.Buses{}
	result := ar.DB.Where("bus_id= ?", id).First(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

// GetChart implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) GetChart(busid int, day time.Time) (*entities.BusSchedule, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	buschart := &entities.BusSchedule{}
	result := ar.DB.Where("bus_id= ? AND day=?", busid, day).First(buschart)
	if result.Error != nil {
		return nil, result.Error
	}
	return buschart, nil
}

// ViewBookingsPerBus implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) ViewBookingsPerBus(busID int, day string) ([]*entities.Booking, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bookings := []*entities.Booking{}
	result := ar.DB.Where("bus_id=? AND booking_date=?", busID, day).Find(&bookings)
	if result.Error != nil {
		log.Println("Unable to fetch the buses")
		return nil, result.Error
	}
	return bookings, nil
}

// ViewAllBookings implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) ViewAllBookings() ([]*entities.Booking, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bookings := []*entities.Booking{}
	result := ar.DB.Find(&bookings)
	if result.Error != nil {
		log.Println("Unable to fetch the buses")
		return nil, result.Error
	}
	return bookings, nil
}

// AddFareForRoute implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) AddFareForRoute(baseFare *entities.BaseFare) (*entities.BaseFare, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ar.DB.Create(baseFare)
	if result.Error != nil {
		log.Println("Unable to add bus fare")
		return nil, result.Error
	}
	return baseFare, nil
}

// AddBusSchedule implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) AddBusSchedule(schedule *dto.BusSchedule) (*entities.BusSchedule, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	chart := &entities.BusSchedule{}
	chart.BusID = schedule.BusID
	parsedDate, err := time.Parse("02 01 2006", schedule.Day)
	fmt.Println("Parsed Date:", parsedDate)
	chart.Day = parsedDate
	if err != nil {
		log.Println(err)
	}
	result := ar.DB.Create(chart)
	if result.Error != nil {
		log.Println("Unable to add bus schedule")
		return nil, result.Error
	}
	return chart, nil
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
	provider, err := ar.FindProviderByID(id)
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
	provider, err := ar.FindProviderByID(id)
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
	result := ar.DB.Where("station_name=?", name).First(station)
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
	user, err := ar.FindUserByID(id)
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
	user, err := ar.FindUserByID(id)
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
	provider, err := ar.FindProviderByID(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ar.DB.Delete(provider).Where("provider_id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return provider, nil
}

// DeleteStation implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) DeleteStation(id int) (*entities.Stations, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station, err := ar.FindStationByID(id)
	if err != nil {
		log.Println("User Not found")
		return nil, errors.New("error deleting the user")
	}
	ar.DB.Delete(station).Where("station_id=?", id)
	// ari.DB.Raw("delete from users where id=1")
	return station, nil
}

// DeleteUser implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) DeleteUser(id int) (*entities.User, error) {
	if ar.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user, err := ar.FindUserByID(id)
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
	foundProvider, err := ar.FindProviderByID(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if provider.Address != "" {
		foundProvider.Address = provider.Address
	}
	if provider.CompanyName != "" {
		foundProvider.CompanyName = provider.CompanyName
	}
	if provider.Email != "" {
		foundProvider.Email = provider.Email
	}
	if provider.PhoneNumber != "" {
		foundProvider.PhoneNumber = provider.PhoneNumber
	}
	// if foundProvider.BusCount != provider.BusCount {
	// 	foundProvider.BusCount = provider.BusCount
	// }
	if provider.Password != "" {
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
	foundStation, err := ar.FindStationByID(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	// if station.Location != "" {
	// 	foundStation.Location = station.Location
	// }
	// if station.StationCode != "" {
	// 	foundStation.StationCode = station.StationCode
	// }
	if station.StationName != "" {
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
	foundUser, err := ar.FindUserByID(id)
	if err != nil {
		log.Println("User not found")
		return nil, errors.New("user not found")
	}
	if user.Email != "" {
		foundUser.Email = user.Email
	}
	if user.UserName != "" {
		foundUser.UserName = user.UserName
	}
	if user.DOB != "" {
		foundUser.DOB = user.DOB
	}
	if user.Gender != "" {
		foundUser.Gender = user.Gender
	}
	if user.Password != "" {
		foundUser.Password = user.Password
	}
	if user.PhoneNumber != "" {
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

// FindProviderByID implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindProviderByID(id int) (*entities.ServiceProvider, error) {
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

// FindStationByID implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindStationByID(id int) (*entities.Stations, error) {
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

// FindUserByID implements interfaces.AdminRepository.
func (ar *AdminRepositoryImpl) FindUserByID(id int) (*entities.User, error) {
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

// NewAdminRepository function is used to initialize/instatiate Admin Repository.
func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{
		DB: db,
	}
}
