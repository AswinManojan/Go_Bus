package services

import (
	"encoding/json"
	"errors"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	"gobus/services/interfaces"
	"gobus/utils"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	jwt  *middleware.JwtUtil
}

// FindCoupon implements interfaces.UserService.
func (usi *UserServiceImpl) FindCoupon() ([]*entities.Coupons, error) {
	coupons, err := usi.repo.FindCoupon()
	if err != nil {
		log.Println("Error finding coupon, in providerServiceImpl file")
		return coupons, err
	}
	return coupons, err
}

// BookSeat implements interfaces.UserService.
func (usi *UserServiceImpl) BookSeat(bookreq *dto.BookingRequest, email string) (*entities.Booking, error) {
	if len(bookreq.PassengerId) != len(bookreq.SeatsReserved) {
		log.Println("Error seat-passenger mismatch, in userServiceImpl file")
		return nil, errors.New("seat-passenger count mismatch")
	}
	booking := &entities.Booking{}
	booking.SeatReserved = bookreq.SeatsReserved
	booking.BusId = bookreq.BusId
	user, err := usi.repo.FindUserByEmail(email)
	if err != nil {
		log.Println("Error finding user, in userServiceImpl file")
		return nil, err
	}
	booking.UserId = user.ID
	booking.PassengerId = bookreq.PassengerId
	//Getting bus info
	bus, err := usi.repo.GetBusInfo(int(bookreq.BusId))
	if err != nil {
		log.Println("Error fetching bus details, in userServiceImpl file")
		return nil, err
	}
	scheduleId := int(bus.ScheduleId)
	//Getting bus type
	// busType, err := usi.repo.GetBusTypeDetails(bus.BusTypeCode)
	// if err != nil {
	// 	log.Println("Error fetching bus type details, in userServiceImpl file")
	// 	return nil, err
	// }
	//Getting bus seat layout
	// seatLayout, err := usi.repo.GetSeatLayout(int(busType.SeatLayoutId))
	// if err != nil {
	// 	log.Println("Error fetching seat layout details, in userServiceImpl file")
	// 	return nil, err
	// }
	parsedDate, err := time.Parse("02 01 2006", bookreq.BookingDate)
	if err != nil {
		log.Println("Error parsing the date, in userServiceImpl file")
		return nil, err
	}
	//Getting bus chart
	chart, err := usi.repo.GetChart(int(bookreq.BusId), parsedDate)
	if err != nil {
		log.Println("Error fetching bus schedule, in userServiceImpl file")
		return nil, err
	}
	//Fetching the fare
	bFare, err := usi.repo.GetBaseFare(scheduleId)
	if err != nil {
		log.Println("Error fetching base fare, in userServiceImpl file")
		return nil, err
	}
	booking.ActualFare = float64(bFare.BaseFare)
	ACchecker := bus.BusTypeCode[:2]
	// fmt.Println(ACchecker)
	if ACchecker == "AC" {
		booking.ActualFare = booking.ActualFare * 1.3
		// fmt.Println("Amount updated for AC")
	}
	totalFare := 0.0
	if bus.BusTypeCode == "AC_SL" || bus.BusTypeCode == "SL" {
		totalFare = float64(len(bookreq.PassengerId)) * booking.ActualFare * 1.2
	} else if bus.BusTypeCode == "AC_SL_SE" || bus.BusTypeCode == "SL_SE" {
		c1 := 0
		c2 := 0
		for i := 0; i < len(bookreq.SeatsReserved); i++ {
			k := bookreq.SeatsReserved[i][2]
			// num, _ := strconv.Atoi(bookreq.SeatsReserved[i][:2])
			if k == 'A' || k == 'B' || k == 'C' {
				c1++
			} else if k == 'D' || k == 'E' || k == 'F' {
				c2++
			}
		}
		totalFare = (float64(c1) * booking.ActualFare) + (float64(c2) * booking.ActualFare * 1.2)
	} else {
		totalFare = float64(len(bookreq.PassengerId)) * booking.ActualFare
	}
	booking.ActualFare = totalFare
	discount := 0
	coupon, err := usi.repo.FindCouponById(int(bookreq.UsedCouponId))
	if err != nil {
		log.Println("Error finding coupon, in userServiceImpl file")
		return nil, err
	}
	if coupon.IsActive {
		booking.UsedCouponId = bookreq.UsedCouponId
		discount = int(coupon.Discount)
	} else {
		log.Println("Coupon not active or valid, in userServiceImpl file")
		return nil, errors.New("coupon not active or valid")
	}
	booking.FarePostDiscount = booking.ActualFare * float64((100-float64(discount))/100)
	// fmt.Println(chart)
	// fmt.Print(chart.DeckOneSeatLayout)
	//reserving seat as per seatreserved string
	type DeckOneLayoutstr struct {
		DeckLayout [][]bool `json:"deckOneLayout"`
	}
	type DeckTwoLayoutstr struct {
		DeckLayout [][]bool `json:"deckTwoLayout"`
	}
	var unmarshaledLayoutOne DeckOneLayoutstr
	var unmarshaledLayoutTwo DeckTwoLayoutstr
	var copyLayoutOne DeckOneLayoutstr
	var copyLayoutTwo DeckTwoLayoutstr
	// flag := true
	// flag1 := true
	for i := 0; i < len(bookreq.SeatsReserved); i++ {
		k := bookreq.SeatsReserved[i][2]
		num, _ := strconv.Atoi(bookreq.SeatsReserved[i][:2])
		if k == 'A' || k == 'B' || k == 'C' {
			// if flag {
			// 	flag = false
			// }
			json.Unmarshal(chart.DeckOneSeatLayout, &unmarshaledLayoutOne)
			copyLayoutOne = unmarshaledLayoutOne
			if num > len(unmarshaledLayoutOne.DeckLayout) {
				log.Println("you are trying to book an invalid seat, in userServiceImpl file")
				return nil, errors.New("invalid seat entered")
			}
			// fmt.Print(unmarshaledLayout)
			// if err != nil {
			// 	log.Println("Error unmarshalling seat layout, in userServiceImpl file")
			// 	return nil, err
			// }
			if k == 'A' && !unmarshaledLayoutOne.DeckLayout[num-1][0] {
				unmarshaledLayoutOne.DeckLayout[num-1][0] = true
			} else if k == 'B' && !unmarshaledLayoutOne.DeckLayout[num-1][1] {
				unmarshaledLayoutOne.DeckLayout[num-1][1] = true
			} else if k == 'C' && !unmarshaledLayoutOne.DeckLayout[num-1][2] {
				unmarshaledLayoutOne.DeckLayout[num-1][2] = true
			} else {
				log.Println("Seat you are trying to book is already reserved or invalid seat entered, in userServiceImpl file")
				unmarshaledLayoutOne = copyLayoutOne
				Layout, _ := json.Marshal(&unmarshaledLayoutOne)
				chart.DeckOneSeatLayout = Layout
				_, err := usi.repo.UpdateChart(chart)
				if err != nil {
					log.Println("Could not update the chart, in userService file")
					return nil, err
				}
				return nil, errors.New("seat already reserved")
			}
			// fmt.Println(unmarshaledLayoutOne.DeckLayout)
			Layout, _ := json.Marshal(&unmarshaledLayoutOne)
			chart.DeckOneSeatLayout = Layout
		} else if k == 'D' || k == 'E' || k == 'F' {
			// fmt.Println(chart.DeckTwoSeatLayout)
			// if flag1 {
			// 	flag1 = false
			// 	// fmt.Println("Flag set to false")
			// }
			json.Unmarshal(chart.DeckTwoSeatLayout, &unmarshaledLayoutTwo)
			copyLayoutTwo = unmarshaledLayoutTwo
			// fmt.Println(unmarshaledLayoutTwo)
			if num > len(unmarshaledLayoutTwo.DeckLayout) {
				log.Println("Seat you are trying to book an invalid seat, in userServiceImpl file")
				return nil, errors.New("invalid seat entered")
			}
			// if err != nil {
			// 	log.Println("Error unmarshalling seat layout, in userServiceImpl file")
			// 	return nil, err
			// }
			if k == 'D' && !unmarshaledLayoutTwo.DeckLayout[num-1][0] {
				unmarshaledLayoutTwo.DeckLayout[num-1][0] = true
			} else if k == 'E' && !unmarshaledLayoutTwo.DeckLayout[num-1][1] {
				unmarshaledLayoutTwo.DeckLayout[num-1][1] = true
			} else if k == 'F' && !unmarshaledLayoutTwo.DeckLayout[num-1][2] {
				unmarshaledLayoutTwo.DeckLayout[num-1][2] = true
			} else {
				log.Println("Seat you are trying to book is already reserved, in userServiceImpl file")
				unmarshaledLayoutTwo = copyLayoutTwo
				Layout, _ := json.Marshal(&unmarshaledLayoutTwo)
				chart.DeckTwoSeatLayout = Layout
				_, err := usi.repo.UpdateChart(chart)
				if err != nil {
					log.Println("Could not update the chart, in userService file")
					return nil, err
				}
				return nil, errors.New("seat already reserved")
			}
			// fmt.Println(unmarshaledLayoutTwo.DeckLayout)
			Layout, _ := json.Marshal(&unmarshaledLayoutTwo)
			chart.DeckTwoSeatLayout = Layout
		} else {
			log.Println("Seat you are trying to book an invalid seat, in userServiceImpl file")
			return nil, errors.New("invalid seat entered")
		}
		_, err := usi.repo.UpdateChart(chart)
		if err != nil {
			log.Println("Could not update the chart, in userService file")
			return nil, err
		}
	}
	booked, err := usi.repo.MakeBooking(booking)
	if err != nil {
		log.Println("Unable to make the booking, in userServiceImpl file")
		return nil, err
	}
	return booked, nil
}

// ViewAllPassengers implements interfaces.UserService.
func (usi *UserServiceImpl) ViewAllPassengers(email string) ([]*entities.PassengerInfo, error) {
	passengers, err := usi.repo.ViewAllPassengers(email)
	if err != nil {
		log.Println("Error finding passengers, in userServiceImpl file")
		return nil, err
	}
	return passengers, nil
}

// AddPassenger implements interfaces.UserService.
func (usi *UserServiceImpl) AddPassenger(passenger *entities.PassengerInfo, email string) (*entities.PassengerInfo, error) {
	pass, err := usi.repo.AddPassenger(passenger, email)
	if err != nil {
		log.Println("Passenger not added, in userService file")
		return nil, err
	}
	return pass, nil
}

// FindBus implements interfaces.UserService.
func (usi *UserServiceImpl) FindBus(request *dto.BusRequest) ([]*entities.Bus_schedule, error) {
	depart := request.DepartureStation
	arrival := request.ArrivalStation
	buses, err := usi.repo.FindBus(depart, arrival)
	if err != nil {
		log.Println("No Buses EXISTS for this route, in userService file")
		return nil, errors.New("no Bus exists")
	}
	return buses, nil
}

func (usi *UserServiceImpl) Login(login *dto.LoginRequest) (map[string]string, error) {
	user, err := usi.repo.FindUserByEmail(login.Email)
	if err != nil {
		log.Println("No USER EXISTS, in adminService file")
		return nil, errors.New("no User exists")
	}
	dbHashedPassword := user.Password

	enteredPassword := login.Password

	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(enteredPassword)); err != nil {
		log.Println("Password Mismatch, in adminService file")
		return nil, errors.New("password Mismatch")
	}
	if user.Role != "user" {
		log.Println("Unauthorized, in adminService file")
		return nil, errors.New("unauthorized access")
	}
	if user.IsLocked {
		log.Println("User locked by Admin,Contact admin to unlock the account--- in adminService file")
		return nil, errors.New("locked account")
	}

	// token, err := usi.jwt.CreateToken(login.Email, "user")
	// if err != nil {
	// 	return nil, errors.New("token NOT generated")
	// }
	// return token, nil

	accessToken, refreshToken, err := usi.jwt.CreateToken(login.Email, "user")
	if err != nil {
		return nil, errors.New("token pair NOT generated")
	}

	tokenPair := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return tokenPair, nil

}

func (usi *UserServiceImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		log.Println("Unable to hash password")
		return nil, err
	} else {
		user.Password = hashedPassword
	}
	user, err := usi.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return user, err
	}
	return user, err
}

func NewUserService(repo repository.UserRepository, jwt *middleware.JwtUtil) interfaces.UserService {
	return &UserServiceImpl{
		repo: repo,
		jwt:  jwt,
	}
}
