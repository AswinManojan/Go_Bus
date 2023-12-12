package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"gobus/dto"
	"gobus/entities"
	"gobus/middleware"
	repository "gobus/repository/interfaces"
	"gobus/services/interfaces"
	"gobus/utils"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/razorpay/razorpay-go"

	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl struct is used to Implement the UserService.
type UserServiceImpl struct {
	repo repository.UserRepository
	jwt  *middleware.JwtUtil
}

// SubStationDetails implements interfaces.UserService.
func (usi *UserServiceImpl) SubStationDetails(parent string) ([]string, error) {
	stations, err := usi.repo.GetSubStationDetails(parent)
	if err != nil {
		log.Println("Error fetching the stations, in userServiceImpl file")
		return nil, err
	}
	var substations []string
	for _, student := range stations {
		substations = append(substations, student.SubStation)
	}
	return substations, nil
}

// FindBookingByID implements interfaces.UserService.
func (usi *UserServiceImpl) FindBookingByID(ID int) (*entities.Booking, error) {
	book, err := usi.repo.FindBookingByID(int(ID))
	if err != nil {
		log.Println("Error fetching the booking, in userServiceImpl file")
		return nil, err
	}
	return book, nil
}

// PaymentSuccess implements interfaces.UserService.
func (usi *UserServiceImpl) PaymentSuccess(razor *entities.RazorPay) error {
	bookingID := razor.BookID
	book, err := usi.repo.FindBookingByID(int(bookingID))
	if err != nil {
		log.Println("Error fetching the booking, in userServiceImpl file")
		return err
	}
	book.Status = "Success"
	if _, err := usi.repo.UpdateBooking(book); err != nil {
		log.Println("Error fetching the booking, in userServiceImpl file")
		return err
	}
	if err := usi.repo.PaymentSuccess(razor); err != nil {
		log.Println("Error updating Payment Info, in userServiceImpl file")
		return err
	}
	return nil
}

type pageVariables struct {
	OrderID string
}

// MakePayment implements interfaces.UserService.
func (usi *UserServiceImpl) MakePayment(bookID int) (*dto.MakePaymentResp, error) {
	booking, err := usi.repo.FindBookingByID(bookID)
	if err != nil {
		log.Println("Error fetching the booking, in userServiceImpl file")
		return nil, err
	}
	user, err := usi.repo.GetUserInfo(int(booking.UserID))
	if err != nil {
		log.Println("Error fetching the user info, in userServiceImpl file")
		return nil, err
	}
	amount := booking.FarePostDiscount * 100
	razorKey := os.Getenv("RAZOR_KEY_ID")
	razorSecret := os.Getenv("RAZOR_SECRET")
	client := razorpay.NewClient(razorKey, razorSecret)

	data := map[string]interface{}{
		"amount":   amount,
		"currency": "INR",
		"receipt":  strconv.Itoa(int(booking.BookingID)),
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Printf("Problem getting repositorys information: %v\n", err)
		return nil, err
	}

	value := body["id"]
	str := value.(string)

	homepageVariables := pageVariables{
		OrderID: str,
	}
	paymentResp := &dto.MakePaymentResp{}
	paymentResp.AmountInRupees = int(booking.FarePostDiscount)
	paymentResp.BookingID = int(booking.BookingID)
	paymentResp.Email = user.Email
	paymentResp.PhoneNumber = user.PhoneNumber
	paymentResp.OrderID = homepageVariables.OrderID
	return paymentResp, nil
}

// SeatNameDecoder function is used to decode the seat name from rows and columns
func SeatNameDecoder(row int, column int) string {
	seatName := ""
	if row < 9 {
		seatName = fmt.Sprintf("0%d%c", row+1, 'A'+column)
	} else {
		seatName = fmt.Sprintf("%d%c", row+1, 'A'+column)

	}
	return seatName
}

// SeatAvailabilityChecker implements interfaces.UserService.
func (usi *UserServiceImpl) SeatAvailabilityChecker(seatReq *dto.SeatAvailabilityRequest) (*dto.SeatAvailabilityResponse, error) {
	busID := seatReq.BusID
	parsedDate, _ := time.Parse("02 01 2006", seatReq.Date)
	chart, err := usi.repo.GetChart(busID, parsedDate)
	if err != nil {
		log.Println("Error fetching the chart, in userServiceImpl file")
		return nil, err
	}
	bus, err := usi.repo.GetBusInfo(busID)
	if err != nil {
		log.Println("Error fetching the bus Info, in userServiceImpl file")
		return nil, err
	}
	busTypeCode := bus.BusTypeCode
	totalSleeper := bus.TotalSleeperSeats
	totalSeater := bus.TotalPushBackSeats
	seatStatus := &dto.SeatAvailabilityResponse{}
	seatStatus.BusID = seatReq.BusID
	seatStatus.Date = seatReq.Date
	seatStatus.BusStatus = chart.Status
	type DeckOneLayoutstr struct {
		DeckLayout [][]bool `json:"deckOneLayout"`
	}
	type DeckTwoLayoutstr struct {
		DeckLayout [][]bool `json:"deckTwoLayout"`
	}
	deckOneCount := 0
	deckTwoCount := 0
	var deckOneSlots []string
	var deckTwoSlots []string
	var unmarshaledLayoutOne DeckOneLayoutstr
	var unmarshaledLayoutTwo DeckTwoLayoutstr
	json.Unmarshal(chart.DeckOneSeatLayout, &unmarshaledLayoutOne)
	json.Unmarshal(chart.DeckTwoSeatLayout, &unmarshaledLayoutTwo)
	for i := 0; i < len(unmarshaledLayoutOne.DeckLayout); i++ {
		// var seatCode string
		for j := 0; j < len(unmarshaledLayoutOne.DeckLayout[i]); j++ {
			if unmarshaledLayoutOne.DeckLayout[i][j] {
				deckOneCount++
			} else {
				deckOneSlots = append(deckOneSlots, SeatNameDecoder(i, j))
			}
		}
	}
	for i := 0; i < len(unmarshaledLayoutTwo.DeckLayout); i++ {
		for j := 0; j < len(unmarshaledLayoutTwo.DeckLayout[i]); j++ {
			if unmarshaledLayoutTwo.DeckLayout[i][j] {
				deckOneCount++
			} else {
				deckTwoSlots = append(deckTwoSlots, SeatNameDecoder(i, j+3))
			}
		}
	}
	if busTypeCode == "AC_SL" || busTypeCode == "SL" {
		totalSleeper = totalSleeper - uint(deckOneCount+deckTwoCount)
		seatStatus.AvailableSleeperSlots = append(seatStatus.AvailableSleeperSlots, deckOneSlots...)
		seatStatus.AvailableSleeperSlots = append(seatStatus.AvailableSleeperSlots, deckTwoSlots...)
	} else if busTypeCode == "AC_SE" || busTypeCode == "SE" {
		totalSeater = totalSeater - uint(deckOneCount+deckTwoCount)
		seatStatus.AvailableSeaterSlots = append(seatStatus.AvailableSeaterSlots, deckOneSlots...)
		seatStatus.AvailableSeaterSlots = append(seatStatus.AvailableSeaterSlots, deckTwoSlots...)
	} else {
		totalSleeper = totalSleeper - uint(deckTwoCount)
		totalSeater = totalSeater - uint(deckOneCount)
		seatStatus.AvailableSeaterSlots = append(seatStatus.AvailableSeaterSlots, deckOneSlots...)
		seatStatus.AvailableSleeperSlots = append(seatStatus.AvailableSleeperSlots, deckTwoSlots...)
	}
	seatStatus.SleeperSlotsLeft = int(totalSleeper)
	seatStatus.SeaterSlotsLeft = int(totalSeater)
	seatStatus.BusType = busTypeCode
	return seatStatus, nil
}

// CancelBooking implements interfaces.UserService.
func (usi *UserServiceImpl) CancelBooking(bookID int) (*entities.Booking, error) {
	booking, err := usi.repo.FindBookingByID(bookID)
	if err != nil {
		log.Println("Error finding booking that has to be cancelled, in userServiceImpl file")
		return nil, err
	}
	parsedDate, err := time.Parse("02 01 2006", booking.BookingDate)
	if err != nil {
		log.Println("Error parsing the date, in userServiceImpl file")
		return nil, err
	}
	//Getting bus chart
	chart, _ := usi.repo.GetChart(int(booking.BusID), parsedDate)
	type DeckOneLayoutstr struct {
		DeckLayout [][]bool `json:"deckOneLayout"`
	}
	type DeckTwoLayoutstr struct {
		DeckLayout [][]bool `json:"deckTwoLayout"`
	}
	var unmarshaledLayoutOne DeckOneLayoutstr
	var unmarshaledLayoutTwo DeckTwoLayoutstr
	for i := 0; i < len(booking.SeatReserved); i++ {
		k := booking.SeatReserved[i][2]
		num, _ := strconv.Atoi(booking.SeatReserved[i][:2])
		if k == 'A' || k == 'B' || k == 'C' {
			json.Unmarshal(chart.DeckOneSeatLayout, &unmarshaledLayoutOne)
			if k == 'A' && unmarshaledLayoutOne.DeckLayout[num-1][0] {
				unmarshaledLayoutOne.DeckLayout[num-1][0] = false
			} else if k == 'B' && unmarshaledLayoutOne.DeckLayout[num-1][1] {
				unmarshaledLayoutOne.DeckLayout[num-1][1] = false
			} else if k == 'C' && unmarshaledLayoutOne.DeckLayout[num-1][2] {
				unmarshaledLayoutOne.DeckLayout[num-1][2] = false
			}
			Layout, _ := json.Marshal(&unmarshaledLayoutOne)
			chart.DeckOneSeatLayout = Layout
		} else if k == 'D' || k == 'E' || k == 'F' {
			json.Unmarshal(chart.DeckTwoSeatLayout, &unmarshaledLayoutTwo)
			if k == 'D' && unmarshaledLayoutTwo.DeckLayout[num-1][0] {
				unmarshaledLayoutTwo.DeckLayout[num-1][0] = false
			} else if k == 'E' && unmarshaledLayoutTwo.DeckLayout[num-1][1] {
				unmarshaledLayoutTwo.DeckLayout[num-1][1] = false
			} else if k == 'F' && unmarshaledLayoutTwo.DeckLayout[num-1][2] {
				unmarshaledLayoutTwo.DeckLayout[num-1][2] = false
			}
			Layout, _ := json.Marshal(&unmarshaledLayoutTwo)
			chart.DeckTwoSeatLayout = Layout
		}
		_, err := usi.repo.UpdateChart(chart)
		if err != nil {
			log.Println("Could not update the chart, in userService file")
			return nil, err
		}
	}
	refundPostCancellationCharge := booking.FarePostDiscount * 0.9
	user, _ := usi.repo.GetUserInfo(int(booking.UserID))
	user.UserWallet += int(refundPostCancellationCharge)
	busID := booking.BusID
	bus, _ := usi.repo.GetBusInfo(int(busID))
	provider, _ := usi.repo.GetProviderInfo(int(bus.ProviderID))
	provider.ProviderWallet -= int(refundPostCancellationCharge)
	usi.repo.UpdateUser(user)
	usi.repo.UpdateProvider(provider)
	booking.Status = "Cancelled by User"
	cancelledBooking, err := usi.repo.CancelBooking(booking)
	if err != nil {
		log.Println("unable to cancel the booking, in userServiceImpl file")
		return nil, err
	}
	WhatsappNotifier("The booking has been cancelled successfully and the refund has been transferred to your wallet.", user.PhoneNumber)
	return cancelledBooking, nil
}

// ViewBookings implements interfaces.UserService.
func (usi *UserServiceImpl) ViewBookings(email string) ([]*entities.Booking, error) {
	bookings, err := usi.repo.ViewBookings(email)
	if err != nil {
		log.Println("Error finding bookings, in userServiceImpl file")
		return nil, err
	}
	return bookings, err
}

// FindCoupon implements interfaces.UserService.
func (usi *UserServiceImpl) FindCoupon() ([]*entities.Coupons, error) {
	coupons, err := usi.repo.FindCoupon()
	if err != nil {
		log.Println("Error finding coupon, in userServiceImpl file")
		return nil, err
	}
	return coupons, err
}

// BookSeat implements interfaces.UserService.
func (usi *UserServiceImpl) BookSeat(bookreq *dto.BookingRequest, email string) (*entities.Booking, error) {
	if len(bookreq.PassengerID) != len(bookreq.SeatsReserved) {
		log.Println("Error seat-passenger mismatch, in userServiceImpl file")
		return nil, errors.New("seat-passenger count mismatch")
	}
	booking := &entities.Booking{}
	booking.BookingDate = bookreq.BookingDate
	booking.SeatReserved = bookreq.SeatsReserved
	booking.BusID = bookreq.BusID
	user, err := usi.repo.FindUserByEmail(email)
	if err != nil {
		log.Println("Error finding user, in userServiceImpl file")
		return nil, err
	}
	userBalance := uint(user.UserWallet)
	booking.UserID = user.ID
	passengers, _ := usi.repo.ViewAllPassengers(email)
	var passengerIdlist []int
	for i := 0; i < len(passengers); i++ {
		passengerIdlist = append(passengerIdlist, int(passengers[i].PassengerID))
	}
	// fmt.Println(passengerIdlist)
	// fmt.Println(bookreq.PassengerId)
	for i := 0; i < len(bookreq.PassengerID); i++ {
		count := 0
		for j := 0; j < len(passengerIdlist); j++ {
			if bookreq.PassengerID[i] == int64(passengerIdlist[j]) {
				count++
				break
			}
		}
		fmt.Println(count)
		if count != 1 {
			log.Println("Passenger Id not valid, in userServiceImpl file")
			return nil, errors.New("unknown passenger Id provided")
		}
	}
	booking.PassengerID = bookreq.PassengerID
	//Getting bus info
	bus, err := usi.repo.GetBusInfo(int(bookreq.BusID))
	if err != nil {
		log.Println("Error fetching bus details, in userServiceImpl file")
		return nil, err
	}
	//Getting provider info
	provider, _ := usi.repo.GetProviderInfo(int(bus.ProviderID))
	providerBalance := provider.ProviderWallet
	scheduleID := int(bus.ScheduleID)
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
	chart, err := usi.repo.GetChart(int(bookreq.BusID), parsedDate)
	if err != nil {
		log.Println("Error fetching bus schedule, in userServiceImpl file")
		return nil, err
	}
	if chart.Status != "Active" {
		log.Println("Bus Schedule seems to be cancelled or Inactive, in userServiceImpl file")
		return nil, errors.New("schedule not in active state")
	}
	//Fetching the fare
	bFare, err := usi.repo.GetBaseFare(scheduleID)
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
		totalFare = float64(len(bookreq.PassengerID)) * booking.ActualFare * 1.2
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
		totalFare = float64(len(bookreq.PassengerID)) * booking.ActualFare
	}
	booking.ActualFare = totalFare
	discount := 0
	coupon, err := usi.repo.FindCouponByID(int(bookreq.UsedCouponID))
	if err != nil {
		log.Println("Error finding coupon, in userServiceImpl file")
		return nil, err
	}
	if coupon.IsActive {
		booking.UsedCouponID = bookreq.UsedCouponID
		discount = int(coupon.Discount)
	} else {
		log.Println("Coupon not active or valid, in userServiceImpl file")
		return nil, errors.New("coupon not active or valid")
	}
	booking.Status = "Awaiting Payment"
	booking.FarePostDiscount = booking.ActualFare * float64((100-float64(discount))/100)
	if bookreq.PreferredPaymentType == "Wallet" && booking.FarePostDiscount <= float64(userBalance) {
		userBalance = userBalance - uint(booking.FarePostDiscount)
		providerBalance = providerBalance + int(booking.FarePostDiscount)
		user.UserWallet = int(userBalance)
		provider.ProviderWallet = providerBalance
		booking.Status = "Success"
	} else if bookreq.PreferredPaymentType == "Wallet" && booking.FarePostDiscount > float64(userBalance) {
		log.Println("Insuffucient fund to make the booking using wallet,Redirecting to RazorPay.")
	}
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
	// var copyLayoutOne [][]bool
	// var copyLayoutTwo [][]bool
	// flag := true
	// flag1 := true
	json.Unmarshal(chart.DeckOneSeatLayout, &unmarshaledLayoutOne)
	copyLayoutOne := deepcopy.Copy(unmarshaledLayoutOne).(DeckOneLayoutstr)

	// copyLayoutOne = unmarshaledLayoutOne.DeckLayout
	// copyLayoutOne = append(copyLayoutOne, unmarshaledLayoutOne.DeckLayout...)
	// fmt.Print(copyLayoutOne)
	json.Unmarshal(chart.DeckTwoSeatLayout, &unmarshaledLayoutTwo)
	// copyLayoutTwo = unmarshaledLayoutTwo.DeckLayout
	copyLayoutTwo := deepcopy.Copy(unmarshaledLayoutTwo).(DeckTwoLayoutstr)

	// fmt.Print(copyLayoutOne.DeckLayout)
	// copyLayoutOne = unmarshaledLayoutOne
	// copyLayoutOne := make([][]bool, len(unmarshaledLayoutOne.DeckLayout))
	// for i := range copyLayoutOne {
	// 	copyLayoutOne[i] = make([]int, len(original[i]))
	// 	copy(copySlice[i], original[i])
	// }
	// json.Unmarshal(chart.DeckTwoSeatLayout, &unmarshaledLayoutTwo)
	// fmt.Print(unmarshaledLayoutTwo)
	// copy(te, unmarshaledLayoutTwo.DeckLayout)
	// fmt.Print(te)
	for i := 0; i < len(bookreq.SeatsReserved); i++ {
		k := bookreq.SeatsReserved[i][2]
		num, _ := strconv.Atoi(bookreq.SeatsReserved[i][:2])
		if k == 'A' || k == 'B' || k == 'C' {
			// if flag {
			// 	flag = false
			// }
			json.Unmarshal(chart.DeckOneSeatLayout, &unmarshaledLayoutOne)
			// copyLayoutOne = unmarshaledLayoutOne
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
			// copyLayoutTwo = unmarshaledLayoutTwo
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
	if _, err := usi.repo.UpdateUser(user); err != nil {
		log.Println("Could not update the user Balance, in userService file")
		return nil, err
	}
	if _, err := usi.repo.UpdateProvider(provider); err != nil {
		log.Println("Could not update the provider Balance, in userService file")
		return nil, err
	}
	booked, err := usi.repo.MakeBooking(booking)
	if err != nil {
		log.Println("Unable to make the booking, in userServiceImpl file")
		return nil, err
	}
	message := fmt.Sprintf("The seats %s of the bus %d has been booked for the day %s.", booked.SeatReserved[:], booked.BusID, booked.BookingDate)
	WhatsappNotifier(message, user.PhoneNumber)
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
func (usi *UserServiceImpl) FindBus(request *dto.BusRequest) ([]*entities.BusesResp, error) {
	depart := request.DepartureStation
	arrival := request.ArrivalStation
	buses, err := usi.repo.FindBus(depart, arrival)
	if err != nil {
		log.Println("No Buses EXISTS for this route, in userService file")
		return nil, errors.New("no Bus exists")
	}
	var outbuses []*entities.BusesResp
	if request.Duration != 0 {
		for _, bus := range buses {
			depart, _ := time.Parse("15:04:05", bus.DepartureTime)
			arrive, _ := time.Parse("15:04:05", bus.ArrivalTime)
			if depart.After(arrive) {
				arrive = arrive.Add(24 * time.Hour)
			}
			if arrive.Sub(depart) <= time.Duration(request.Duration)*time.Hour {
				b := &entities.BusesResp{}
				b.BusID = bus.BusID
				b.BusNumber = bus.BusNumber
				b.BusTypeCode = bus.BusTypeCode
				b.TotalPushBackSeats = bus.TotalPushBackSeats
				b.TotalSleeperSeats = bus.TotalSleeperSeats
				outbuses = append(outbuses, b)
			}
		}
		// buses = outbuses
		// return outbuses, nil
	}
	// if request.Price != 0 {
	// 	var outpricebuses []*entities.BusScheduleCombo
	// 	for _, bus := range buses {
	// 		baseFare, _ := usi.repo.GetBaseFare(int(bus.ScheduleID))
	// 		if request.Price <= int(baseFare.BaseFare) {
	// 			outpricebuses = append(outpricebuses, bus)
	// 		}

	// 	}
	// }
	return outbuses, nil
}

// Login function is used to login the user into the application
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

// RegisterUser function is used to register the user with the hashed password.
func (usi *UserServiceImpl) RegisterUser(user *entities.User) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Unable to hash password")
		return nil, err
	}
	user.Password = hashedPassword
	users, err := usi.repo.RegisterUser(user)
	if err != nil {
		log.Println("User not added, adminService file")
		return users, err
	}
	return users, err
}

// NewUserService function returns UserServiceImpl of type UserService Interface
func NewUserService(repo repository.UserRepository, jwt *middleware.JwtUtil) interfaces.UserService {
	return &UserServiceImpl{
		repo: repo,
		jwt:  jwt,
	}
}
