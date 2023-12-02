package interfaces

import (
	"gobus/entities"
	"time"
)

type UserRepository interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindBus(depart string, arrival string) ([]*entities.Bus_schedule, error)
	FindSchedule(depart string, arrival string) (*entities.Schedule, error)
	AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error)
	MakeBooking(booking *entities.Booking) (*entities.Booking, error)
	ViewAllPassengers(email string) ([]*entities.PassengerInfo, error)
	FindCoupon() ([]*entities.Coupons, error)
	FindCouponById(id int) (*entities.Coupons, error)
	GetBusTypeDetails(code string) (*entities.BusType, error)
	GetChart(busid int, day time.Time) (*entities.BusSchedule, error)
	GetSeatLayout(id int) (*entities.BusSeatLayout, error)
	GetBusInfo(id int) (*entities.Buses, error)
	GetBaseFare(scheduleId int) (*entities.BaseFare, error)
	UpdateChart(chart *entities.BusSchedule) (*entities.BusSchedule, error)
	ViewBookings(email string) ([]*entities.Booking, error)
	CancelBooking(booking *entities.Booking) (*entities.Booking, error)
	FindBookingById(bookId int) (*entities.Booking, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	GetProviderInfo(providerId int) (*entities.ServiceProvider, error)
	UpdateProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	GetUserInfo(userId int) (*entities.User, error)
}
