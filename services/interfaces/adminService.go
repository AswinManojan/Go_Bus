package interfaces

import (
	"gobus/dto"
	"gobus/entities"
)

// AdminService inteface is used as an interface for AdminServiceImplementation.
type AdminService interface {
	Login(loginRequest *dto.LoginRequest) (map[string]string, error)
	FindUser(id int) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	UpdateUser(id int, user entities.User) (*entities.User, error)
	DeleteUser(id int) (*entities.User, error)
	BlockUser(id int) (*entities.User, error)
	UnBlockUser(id int) (*entities.User, error)
	FindProvider(id int) (*entities.ServiceProvider, error)
	FindAllProvider() ([]*entities.ServiceProvider, error)
	UpdateProvider(id int, provider entities.ServiceProvider) (*entities.ServiceProvider, error)
	DeleteProvider(id int) (*entities.ServiceProvider, error)
	BlockProvider(id int) (*entities.ServiceProvider, error)
	UnBlockProvider(id int) (*entities.ServiceProvider, error)
	FindStation(id int) (*entities.Stations, error)
	FindStationByName(name string) (*entities.Stations, error)
	FindAllStations() ([]*entities.Stations, error)
	UpdateStation(id int, station entities.Stations) (*entities.Stations, error)
	DeleteStation(id int) (*entities.Stations, error)
	AddStation(station *entities.Stations) (*entities.Stations, error)
	AddBusSchedule(schedule *dto.BusSchedule) (*entities.BusSchedule, error)
	AddFareForRoute(baseFare *entities.BaseFare) (*entities.BaseFare, error)
	ViewAllBookings() ([]*entities.Booking, error)
	ViewBookingsPerBus(busID int, day string) ([]*entities.Booking, error)
	CancelBus(busID int, day string) (string, error)
}
