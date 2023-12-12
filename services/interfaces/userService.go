package interfaces

import (
	"gobus/dto"
	"gobus/entities"
)

// UserService inteface is used as an interface for UserServiceImplementation
type UserService interface {
	Login(login *dto.LoginRequest) (map[string]string, error)
	RegisterUser(user *entities.User) (*entities.User, error)
	FindBus(request *dto.BusRequest) ([]*entities.BusScheduleCombo, error)
	AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error)
	ViewAllPassengers(email string) ([]*entities.PassengerInfo, error)
	BookSeat(bookreq *dto.BookingRequest, email string) (*entities.Booking, error)
	FindCoupon() ([]*entities.Coupons, error)
	ViewBookings(email string) ([]*entities.Booking, error)
	CancelBooking(bookID int) (*entities.Booking, error)
	SeatAvailabilityChecker(seatReq *dto.SeatAvailabilityRequest) (*dto.SeatAvailabilityResponse, error)
	MakePayment(bookID int) (*dto.MakePaymentResp, error)
	PaymentSuccess(razor *entities.RazorPay) error
	FindBookingByID(ID int) (*entities.Booking, error)
	SubStationDetails(parent string) ([]string, error)
}
