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

type UserRepositoryImpl struct {
	DB *gorm.DB
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
func (ur *UserRepositoryImpl) GetBaseFare(scheduleId int) (*entities.BaseFare, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	baseFare := &entities.BaseFare{}
	result := ur.DB.Where("schedule_id=?", scheduleId).First(baseFare)
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

// GetBusDetails implements interfaces.UserRepository.
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

// FindCouponById implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindCouponById(id int) (*entities.Coupons, error) {
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
	fmt.Println(retrievedUser.ID)
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

func (ur *UserRepositoryImpl) AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	retrievedUser, _ := ur.FindUserByEmail(email)
	passenger.UserId = retrievedUser.ID
	result := ur.DB.Create(&passenger)
	if result.Error != nil {
		log.Println("Unable to add passenger, UserRepositoryImpl package")
		return nil, errors.New("passenger not added to db")
	}
	return passenger, nil

}

// FindBus implements interfaces.UserRepository.
func (ur *UserRepositoryImpl) FindBus(depart string, arrival string) ([]*entities.Bus_schedule, error) {
	if ur.DB == nil {
		log.Println("Error connecting DB in FindBus method, AdminRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}
	schedule, err := ur.FindSchedule(depart, arrival)
	fmt.Println(schedule.ScheduleId)
	if err != nil {
		log.Println("Schedule not found in DB")
		return nil, errors.New("schedule not Found in DB")
	}
	// fmt.Print(schedule.ScheduleId)
	query := "SELECT * FROM buses b JOIN schedules s ON b.schedule_id = s.schedule_id WHERE b.schedule_id = ?"
	buses := []*entities.Bus_schedule{}
	ur.DB.Raw(query, schedule.ScheduleId).Scan(&buses)
	// result := ur.DB.Where("schedule_id=?", schedule.ScheduleId).Find(&buses)
	// if result.Error != nil {
	// 	log.Println("Bus not available")
	// 	return nil, errors.New("bus not available")
	// }
	return buses, nil
}
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

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}
