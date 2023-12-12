package services

import (
	"errors"
	"fmt"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	service "gobus/services/interfaces"
	"log"
	"os"
	"time"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func sendCancellationEmail(recipientEmail, cancel string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "gobusaswin@gmail.com")
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "GoBus: Bus Cancelled")

	m.SetBody("text/plain", "Message: "+cancel)

	d := gomail.NewDialer("smtp.gmail.com", 587, "gobusaswin@gmail.com", "zfej mjdj hhzq lxve")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// AdminServiceImpl struct is used to Implement the Admin Service.
type AdminServiceImpl struct {
	repo repository.AdminRepository
	jwt  *middleware.JwtUtil
}

// WhatsappNotifier function was added to notify the customer on bus cancellation, but will not notify via whatsapp only via sms.
func WhatsappNotifier(messageText string, toNumber string) {

	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	params.SetBody(messageText)
	params.SetFrom("+15152001155")
	// params.SetTo(toNumber)
	params.SetTo(os.Getenv("MY_NUMBER"))

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

// CancelBus implements interfaces.AdminService.
func (as *AdminServiceImpl) CancelBus(busID int, day string) (string, error) {
	parsedDate, _ := time.Parse("02 01 2006", day)
	chart, _ := as.repo.GetChart(busID, parsedDate)
	if chart.Status != "Active" {
		return "Bus was already in Inactive or Cancelled state", errors.New("bus already in inactive or cancelled state")
	}
	chart.Status = "Cancelled"
	bookings, _ := as.repo.ViewBookingsToBeCancelled(busID, day)
	result := make(chan error)
	bus, _ := as.repo.GetBusInfo(busID)
	schedule, _ := as.repo.GetRouteByBus(int(bus.ScheduleID))
	for i := 0; i < len(bookings); i++ {
		// if bookings[i].Status != "Success" {
		// 	continue
		// }
		go func(booking *entities.Booking) {
			defer close(result)
			amount := booking.FarePostDiscount
			user, _ := as.repo.FindUserByID(int(booking.UserID))
			provider, _ := as.repo.FindProviderByID(int(bus.ProviderID))
			user.UserWallet += int(amount)
			provider.ProviderWallet -= int(amount)
			booking.Status = "Cancelled by Admin"
			if _, err := as.repo.UpdateProvider(provider); err != nil {
				log.Println("Error updating the provider, in adminServiceImpl file")
				result <- err
			}
			if _, err := as.repo.UpdateUser(user); err != nil {
				log.Println("Error updating the user, in adminServiceImpl file")
				result <- err
			}
			if _, err := as.repo.UpdateBooking(booking); err != nil {
				log.Println("Error updating the booking, in adminServiceImpl file")
				result <- err
			}
			message := fmt.Sprintf("The bus %d has been cancelled for the day %s due to unforeseen circumstances. Sorry for the inconvinience caused. Your amount has been refunded to your wallet. \n Booking Info \n Booking ID: %d \n UserID: %d \n UsedCouponID: %d \n Actual Fare: %d \n FarePostDiscount: %d \n BusID: %d \n Departure Location: %s \n Arrival Location: %s \n BookingDate: %s", busID, day, booking.BookingID, booking.UserID, booking.UsedCouponID, int(booking.ActualFare), int(booking.FarePostDiscount), booking.BusID, schedule.DepartureStation, schedule.ArrivalStation, booking.BookingDate)
			if err := sendCancellationEmail(user.Email, message); err != nil {
				log.Println("Error sending bus cancellation email, in adminServiceImpl file")
				result <- err
			}
			fmt.Println(user.PhoneNumber)
			WhatsappNotifier(message, user.PhoneNumber)
			fmt.Printf("Email sent to %s \n", user.Email)
			result <- nil
		}(bookings[i])
	}
	for i := 0; i < len(bookings); i++ {
		if err := <-result; err != nil {
			log.Println("Error in goroutine:", err)
			return "", err
		}
	}
	if _, err := as.repo.UpdateChart(chart); err != nil {
		log.Println("Error updating the schedule(chart), in adminServiceImpl file")
		return "", err
	}
	response := fmt.Sprintf("Cancelled the bus %d scheduled for date %s", busID, day)
	return response, nil
}

// ViewAllBookings implements interfaces.AdminService.
func (as *AdminServiceImpl) ViewAllBookings() ([]*entities.Booking, error) {
	bookings, err := as.repo.ViewAllBookings()
	if err != nil {
		log.Println("Error fetching the bookings, in adminServiceImpl file")
		return nil, err
	}
	return bookings, nil
}

// ViewBookingsPerBus implements interfaces.AdminService.
func (as *AdminServiceImpl) ViewBookingsPerBus(busID int, day string) ([]*entities.Booking, error) {
	bookings, err := as.repo.ViewBookingsPerBus(busID, day)
	if err != nil {
		log.Println("Error fetching the bookings, in adminServiceImpl file")
		return nil, err
	}
	return bookings, nil
}

// AddFareForRoute implements interfaces.AdminService.
func (as *AdminServiceImpl) AddFareForRoute(baseFare *entities.BaseFare) (*entities.BaseFare, error) {
	baseFares, err := as.repo.AddFareForRoute(baseFare)
	if err != nil {
		log.Println("Error Adding baseFare, in adminServiceImpl file")
		return nil, err
	}
	return baseFares, nil
}

// AddBusSchedule implements interfaces.AdminService.
func (as *AdminServiceImpl) AddBusSchedule(schedule *dto.BusSchedule) (*entities.BusSchedule, error) {
	schedules, err := as.repo.AddBusSchedule(schedule)
	if err != nil {
		log.Println("Error Creating schedule, in adminServiceImpl file")
		return schedules, err
	}
	return schedules, nil
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
	provider, err := as.repo.FindProviderByID(id)
	if err != nil {
		log.Println("Error finding provider, in adminServiceImpl file")
		return nil, err
	}
	return provider, nil
}

// FindStation implements interfaces.AdminService.
func (as *AdminServiceImpl) FindStation(id int) (*entities.Stations, error) {
	station, err := as.repo.FindStationByID(id)
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
	user, err := as.repo.FindUserByID(id)
	if err != nil {
		log.Println("Error finding user, in adminServiceImpl file")
		return nil, err
	}
	return user, nil
}

// // Login implements interfaces.AdminService.
// func (as *AdminServiceImpl) Login(loginRequest *dto.LoginRequest) (string, error) {
// 	user, err := as.repo.FindUserByEmail(loginRequest.Email)
// 	if err != nil {
// 		log.Println("No USER EXISTS, in adminService file")
// 		return "", errors.New("no User exists")
// 	}
// 	dbHashedPassword := user.Password // Replace with the actual hashed password.

// 	enteredPassword := loginRequest.Password // Replace with the password entered by the user during login.

// 	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(enteredPassword)); err != nil {
// 		// Passwords match. Allow the user to log in.
// 		log.Println("Password Mismatch, in adminService file")
// 		return "", errors.New("password Mismatch")
// 	}
// 	if user.Role != "admin" {
// 		log.Println("Unauthorized, in adminService file")
// 		return "", errors.New("unauthorized access")
// 	}
// 	token, err := as.jwt.CreateToken(loginRequest.Email, "admin")
// 	if err != nil {
// 		return "", errors.New("token NOT generated")
// 	}

//		return token, nil
//	}

// Login function is used to log the user in to the application
func (as *AdminServiceImpl) Login(loginRequest *dto.LoginRequest) (map[string]string, error) {
	user, err := as.repo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return nil, errors.New("no User exists")
	}
	dbHashedPassword := user.Password

	enteredPassword := loginRequest.Password

	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(enteredPassword)); err != nil {
		log.Println("Password Mismatch, in adminService file")
		return nil, errors.New("password Mismatch")
	}
	if user.Role != "admin" {
		log.Println("Unauthorized, in adminService file")
		return nil, errors.New("unauthorized access")
	}
	accessToken, refreshToken, err := as.jwt.CreateToken(loginRequest.Email, "admin")
	if err != nil {
		return nil, errors.New("token pair NOT generated")
	}

	tokenPair := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return tokenPair, nil
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

// NewAdminService function return AdminServiceImpl of type AdminService interface
func NewAdminService(repository repository.AdminRepository, jwt *middleware.JwtUtil) service.AdminService {
	return &AdminServiceImpl{
		repo: repository,
		jwt:  jwt,
	}
}
