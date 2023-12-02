package interfaces

import (
	"gobus/dto"
	"gobus/entities"
)

type UserService interface {
	Login(login *dto.LoginRequest) (map[string]string, error)
	RegisterUser(user *entities.User) (*entities.User, error)
	FindBus(request *dto.BusRequest) ([]*entities.Bus_schedule, error)
	AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error)
	ViewAllPassengers(email string) ([]*entities.PassengerInfo, error)
	BookSeat(bookreq *dto.BookingRequest, email string) (*entities.Booking, error)
	FindCoupon() ([]*entities.Coupons, error)
	ViewBookings(email string) ([]*entities.Booking, error)
	CancelBooking(bookId int) (*entities.Booking, error)
}
