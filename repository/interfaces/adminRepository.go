package interfaces

import (
	"gobus/dto"
	"gobus/entities"
	"time"
)

// AdminRepository interface is the interface used for admin repository
type AdminRepository interface {
	FindUserByID(id int) (*entities.User, error)
	FindUserByEmail(mail string) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	EditUser(id int, user *entities.User) (*entities.User, error)
	DeleteUser(id int) (*entities.User, error)
	BlockUser(id int) (*entities.User, error)
	UnBlockUser(id int) (*entities.User, error)
	FindProviderByID(id int) (*entities.ServiceProvider, error)
	FindAllProviders() ([]*entities.ServiceProvider, error)
	EditProvider(id int, provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	DeleteProvider(id int) (*entities.ServiceProvider, error)
	BlockProvider(id int) (*entities.ServiceProvider, error)
	UnBlockProvider(id int) (*entities.ServiceProvider, error)
	FindStationByID(id int) (*entities.Stations, error)
	FindStationByName(name string) (*entities.Stations, error)
	FindAllStations() ([]*entities.Stations, error)
	EditStation(id int, station *entities.Stations) (*entities.Stations, error)
	DeleteStation(id int) (*entities.Stations, error)
	AddStation(station *entities.Stations) (*entities.Stations, error)
	AddBusSchedule(schedule *dto.BusSchedule) (*entities.BusSchedule, error)
	AddFareForRoute(baseFare *entities.BaseFare) (*entities.BaseFare, error)
	ViewAllBookings() ([]*entities.Booking, error)
	ViewBookingsPerBus(busID int, day string) ([]*entities.Booking, error)
	GetChart(busid int, day time.Time) (*entities.BusSchedule, error)
	GetBusInfo(id int) (*entities.Buses, error)
	UpdateProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	UpdateChart(chart *entities.BusSchedule) (*entities.BusSchedule, error)
	UpdateBooking(booking *entities.Booking) (*entities.Booking, error)
	ViewBookingsToBeCancelled(busID int, day string) ([]*entities.Booking, error)
	GetRouteByBus(scheduleID int) (*entities.Schedule, error)
}
