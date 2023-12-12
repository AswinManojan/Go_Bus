package repository

import (
	"errors"
	"fmt"
	"gobus/entities"
	"gobus/repository/interfaces"
	"log"
	"time"

	"gorm.io/gorm"
)

// UserRepositoryImpl struct is used to define User Repository Implementation.
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// GetSubStationDetails implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetSubStationDetails(parent string) ([]*entities.SubStation, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	stations := []*entities.SubStation{}
	result := ur.DB.Where("parent_location=?", parent).Find(&stations)
	if result.Error != nil {
		return nil, result.Error
	}
	return stations, nil
}

// PaymentSuccess implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) PaymentSuccess(razor *entities.RazorPay) error {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return errors.New("error connecting database")
	}
	result := ur.DB.Create(&razor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBooking implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) UpdateBooking(booking *entities.Booking) (*entities.Booking, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Save(booking)
	if result.Error != nil {
		return nil, result.Error
	}
	return booking, nil
}

// GetUserInfo implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetUserInfo(userID int) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	user := &entities.User{}
	result := ur.DB.Where("id=?", userID).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateProvider implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) UpdateProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Save(provider)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}

// GetProviderInfo implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetProviderInfo(providerID int) (*entities.ServiceProvider, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider := &entities.ServiceProvider{}
	result := ur.DB.Where("provider_id=?", providerID).First(provider)
	if result.Error != nil {
		return nil, result.Error
	}
	return provider, nil
}

// UpdateUser implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) UpdateUser(user *entities.User) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// FindBookingByID implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindBookingByID(bookID int) (*entities.Booking, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	booking := &entities.Booking{}
	result := ur.DB.Where("booking_id=?", bookID).First(booking)
	if result.Error != nil {
		return nil, result.Error
	}
	return booking, nil
}

// CancelBooking implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) CancelBooking(booking *entities.Booking) (*entities.Booking, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Save(booking)
	if result.Error != nil {
		return nil, result.Error
	}
	return booking, nil
}

// ViewBookings implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) ViewBookings(email string) ([]*entities.Booking, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	retrievedUser, _ := ur.FindUserByEmail(email)
	bookings := []*entities.Booking{}
	result := ur.DB.Where("user_id=?", retrievedUser.ID).Find(&bookings)
	if result.Error != nil {
		return nil, result.Error
	}
	return bookings, nil
}

// UpdateChart implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) UpdateChart(chart *entities.BusSchedule) (*entities.BusSchedule, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Save(chart)
	if result.Error != nil {
		return nil, result.Error
	}
	return chart, nil
}

// GetBaseFare implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetBaseFare(scheduleID int) (*entities.BaseFare, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	baseFare := &entities.BaseFare{}
	result := ur.DB.Where("schedule_id=?", scheduleID).First(baseFare)
	if result.Error != nil {
		return nil, result.Error
	}
	return baseFare, nil
}

// GetBusInfo implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetBusInfo(id int) (*entities.Buses, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bus := &entities.Buses{}
	result := ur.DB.Where("bus_id= ?", id).First(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

// GetSeatLayout implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetSeatLayout(id int) (*entities.BusSeatLayout, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	seatLayout := &entities.BusSeatLayout{}
	result := ur.DB.Where("id= ?", id).First(seatLayout)
	if result.Error != nil {
		return nil, result.Error
	}
	return seatLayout, nil
}

// GetChart implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetChart(busid int, day time.Time) (*entities.BusSchedule, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	buschart := &entities.BusSchedule{}
	result := ur.DB.Where("bus_id= ? AND day=?", busid, day).First(buschart)
	if result.Error != nil {
		return nil, result.Error
	}
	return buschart, nil
}

// GetBusTypeDetails implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) GetBusTypeDetails(code string) (*entities.BusType, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bus := &entities.BusType{}
	result := ur.DB.Where("bus_type_code=?", code).First(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

// FindCoupon implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindCoupon() ([]*entities.Coupons, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	coupons := []*entities.Coupons{}
	result := ur.DB.Find(&coupons)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupons, nil
}

// FindCouponByID implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindCouponByID(id int) (*entities.Coupons, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	coupon := &entities.Coupons{}
	result := ur.DB.Where("coupon_id", id).First(coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupon, nil
}

// ViewAllPassengers implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) ViewAllPassengers(email string) ([]*entities.PassengerInfo, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	retrievedUser, _ := ur.FindUserByEmail(email)
	// fmt.Println(retrievedUser.ID)
	passengers := []*entities.PassengerInfo{}
	result := ur.DB.Where("user_id=?", retrievedUser.ID).Find(&passengers)
	fmt.Println(passengers)
	if result.Error != nil {
		return nil, result.Error
	}
	return passengers, nil
}

// MakeBooking implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) MakeBooking(booking *entities.Booking) (*entities.Booking, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	result := ur.DB.Create(&booking)
	if result.Error != nil {
		log.Println("Unable to make booking, UserRepositoryImpl package")
		return nil, errors.New("booking not added to db")
	}
	return booking, nil
}

// AddPassenger function is used to add the passenger
func (ur *UserRepositoryImpl) AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	retrievedUser, _ := ur.FindUserByEmail(email)
	passenger.UserID = retrievedUser.ID
	result := ur.DB.Create(&passenger)
	if result.Error != nil {
		log.Println("Unable to add passenger, UserRepositoryImpl package")
		return nil, errors.New("passenger not added to db")
	}
	return passenger, nil

}

// FindBus implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindBus(depart string, arrival string) ([]*entities.BusScheduleCombo, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindBus method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}
	departStationInfo, err := ur.GetParentLocation(depart)
	if err == nil {
		depart = departStationInfo.ParentLocation
	}
	arrivalStationInfo, err := ur.GetParentLocation(arrival)
	if err == nil {
		arrival = arrivalStationInfo.ParentLocation
	}
	schedule, err := ur.FindSchedule(depart, arrival)
	// fmt.Println(schedule.ScheduleID)
	if err != nil {
		log.Println("Schedule not found in DB")
		return nil, errors.New("schedule not Found in DB")
	}
	// fmt.Print(schedule.ScheduleId)
	query := "SELECT * FROM buses b JOIN schedules s ON b.schedule_id = s.schedule_id WHERE b.schedule_id = ?"
	buses := []*entities.BusScheduleCombo{}
	ur.DB.Raw(query, schedule.ScheduleID).Scan(&buses)
	// result := ur.DB.Where("schedule_id=?", schedule.ScheduleId).Find(&buses)
	// if result.Error != nil {
	// 	log.Println("Bus not available")
	// 	return nil, errors.New("bus not available")
	// }
	return buses, nil
}

//GetParentLocation function is used to fetch the parent location based on the sub station name shared.
func (ur *UserRepositoryImpl) GetParentLocation(name string) (*entities.SubStation, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindBus method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}
	station := &entities.SubStation{}
	if err := ur.DB.Where("sub_station=?", name).First(&station).Error; err != nil {
		log.Println("Unable to find parent station, UserRepositoryImpl package")
		return nil, errors.New("station not added to db")
	}
	return station, nil
}

// FindSchedule function is used to find the schedule
func (ur *UserRepositoryImpl) FindSchedule(depart string, arrival string) (*entities.Schedule, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindBus method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}
	schedule := &entities.Schedule{}
	result := ur.DB.Where("departure_station=? AND arrival_station=?", depart, arrival).First(schedule)
	if result.Error != nil {
		log.Println("Schedule not found in DB")
		return nil, errors.New("schedule not Found in DB")
	}
	return schedule, nil
}

// FindUserByEmail implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindUserByEmail(email string) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}

	user := &entities.User{}
	result := ur.DB.Where("email = ?", email).First(user)

	if result.Error != nil {
		log.Println("User not found in DB")
		return nil, errors.New("user not Found in DB")
	}

	return user, nil
}

// RegisterUser implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindUserById method, AdminRepositoryImpl package")
		return nil, errors.New("error Coonecting Database")
	}

	result := ur.DB.Create(&user)
	if result.Error != nil {
		log.Println("Unable to add user, AdminRepositoryImpl package")
		return user, errors.New("user already exists")
	}

	return user, nil
}

// NewUserRepository function is used to instatiate User Repository
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
