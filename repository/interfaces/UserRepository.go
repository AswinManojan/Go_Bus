package interfaces

import (
	"gobus/entities"
	"time"
)

// UserRepository interface is the interface for User Repository.
type UserRepository interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	FindBus(depart string, arrival string) ([]*entities.BusScheduleCombo, error)
	FindSchedule(depart string, arrival string) (*entities.Schedule, error)
	AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error)
	MakeBooking(booking *entities.Booking) (*entities.Booking, error)
	ViewAllPassengers(email string) ([]*entities.PassengerInfo, error)
	FindCoupon() ([]*entities.Coupons, error)
	FindCouponByID(id int) (*entities.Coupons, error)
	GetBusTypeDetails(code string) (*entities.BusType, error)
	GetChart(busid int, day time.Time) (*entities.BusSchedule, error)
	GetSeatLayout(id int) (*entities.BusSeatLayout, error)
	GetBusInfo(id int) (*entities.Buses, error)
	GetBaseFare(scheduleID int) (*entities.BaseFare, error)
	UpdateChart(chart *entities.BusSchedule) (*entities.BusSchedule, error)
	ViewBookings(email string) ([]*entities.Booking, error)
	CancelBooking(booking *entities.Booking) (*entities.Booking, error)
	FindBookingByID(bookID int) (*entities.Booking, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	GetProviderInfo(providerID int) (*entities.ServiceProvider, error)
	UpdateProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	GetUserInfo(userID int) (*entities.User, error)
	UpdateBooking(booking *entities.Booking) (*entities.Booking, error)
	PaymentSuccess(razor *entities.RazorPay) error
	GetParentLocation(name string) (*entities.SubStation, error)
	GetSubStationDetails(parent string) ([]*entities.SubStation, error)
}
