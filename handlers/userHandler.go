package handlers

import (
	"fmt"
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler struct is used to setup User Handler
type UserHandler struct {
	user interfaces.UserService
}

// RegisterUser function is used to register the user
func (uh *UserHandler) RegisterUser(c *gin.Context) {
	user := &entities.User{}
	c.BindJSON(user)
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	user, err := uh.user.RegisterUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to register the user",
			"data":    err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "User registered successfully",
		"data":    user,
	})
}

// Home function is used to retrieve the homepage.
func (uh *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Welcome to home page",
	})
}

// Login function is used to log the user into the application
func (uh *UserHandler) Login(c *gin.Context) {
	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)
	if err := validate.Struct(LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	token, err := uh.user.Login(LoginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "User login failed",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "User logged in successfully",
		"status":  "Success",
		"data":    token,
	})
}

// FindBus function is used to find the bus
func (uh *UserHandler) FindBus(c *gin.Context) {
	BusRequest := &dto.BusRequest{}
	c.BindJSON(BusRequest)
	if err := validate.Struct(BusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	buses, err := uh.user.FindBus(BusRequest)
	if len(buses) == 0 {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "No Bus has been found",
			"status":  "Success",
			"data":    buses,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bus not found for this route",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Buses has been found",
		"status":  "Success",
		"data":    buses,
	})
}

// AddPassenger function is used to add the passenger.
func (uh *UserHandler) AddPassenger(c *gin.Context) {
	pass := &entities.PassengerInfo{}
	c.BindJSON(pass)
	if err := validate.Struct(pass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	email := c.MustGet("email").(string)
	passenger, err := uh.user.AddPassenger(pass, email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add a new passenger",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the passenger",
		"data":    passenger,
	})
}

// ViewAllPassengers is used to view all the passengers
func (uh *UserHandler) ViewAllPassengers(c *gin.Context) {
	email := c.MustGet("email").(string)
	pass, err := uh.user.ViewAllPassengers(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the passengers",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully fetched the passengers",
		"data":    pass,
	})
}

// BookSeat function is used to book the Seat
func (uh *UserHandler) BookSeat(c *gin.Context) {
	bookreq := &dto.BookingRequest{}
	c.BindJSON(bookreq)
	if err := validate.Struct(bookreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	email := c.MustGet("email").(string)
	booking, err := uh.user.BookSeat(bookreq, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to book the seat",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Seat has been booked.",
		"status":  "Success",
		"data":    booking,
	})
}

// FindCoupon function is used to find the coupons.
func (uh *UserHandler) FindCoupon(c *gin.Context) {
	coupons, err := uh.user.FindCoupon()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the coupon",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the coupons",
		"data":    coupons,
	})
}

// ViewBookings function is used to view all booking of that user.
func (uh *UserHandler) ViewBookings(c *gin.Context) {
	email := c.MustGet("email").(string)
	bookings, err := uh.user.ViewBookings(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the bookings",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the bookings",
		"data":    bookings,
	})
}

// CancelBooking function is used to cancel the booking
func (uh *UserHandler) CancelBooking(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	bookings, err := uh.user.CancelBooking(intID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the bookings",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the bookings",
		"data":    bookings,
	})
}

// SeatStatus is used to get the seat availability details.
func (uh *UserHandler) SeatStatus(c *gin.Context) {
	seatReq := &dto.SeatAvailabilityRequest{}
	c.BindJSON(seatReq)
	if err := validate.Struct(seatReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	seatResp, err := uh.user.SeatAvailabilityChecker(seatReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to check the seat availability",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the seats Status",
		"data":    seatResp,
	})
}

// MakePayment function is used to make the payment
func (uh *UserHandler) MakePayment(c *gin.Context) {
	ID := c.Param("bookid")
	bookID, _ := strconv.Atoi(ID)
	book, err := uh.user.FindBookingByID(bookID)
	if err != nil {
		log.Println("Error fetching the booking")
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Failed",
			"message": "Error fetching the booking Info",
			"data":    nil,
		})
		return
	}
	if book.Status == "Success" {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "Failed",
			"message": "Booking already in Success state",
			"data":    nil,
		})
		return
	}
	// fmt.Print(bookID)
	paymentResp, err := uh.user.MakePayment(bookID)
	if err != nil {
		fmt.Printf("Problem getting repositorys information: %v\n", err)

	}
	c.HTML(http.StatusOK, "app.html", gin.H{
		"bookID":      paymentResp.BookingID,
		"totalPrice":  paymentResp.AmountInRupees,
		"total":       paymentResp.AmountInRupees * 100,
		"orderID":     paymentResp.OrderID,
		"email":       paymentResp.Email,
		"phoneNumber": paymentResp.PhoneNumber,
	})
}

// PaymentSuccess function is used to implement the logic once the payment is successful.
func (uh *UserHandler) PaymentSuccess(c *gin.Context) {
	bookIDstr := c.Query("bookID")
	// fmt.Println(bookIDstr,"****************************************************************")
	BookID, _ := strconv.Atoi(bookIDstr)
	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	signature := c.Query("signature")
	paymentAmount := c.Query("total")
	amount, _ := strconv.Atoi(paymentAmount)

	rPay := &entities.RazorPay{
		BookID:          uint(BookID),
		RazorPaymentID:  paymentID,
		Signature:       signature,
		RazorPayOrderID: orderID,
		AmountPaid:      float64(amount),
	}
	err := uh.user.PaymentSuccess(rPay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Error while updating data into DB",
			"data":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true})
}

// SuccessPage function is used to display the success page
func (uh *UserHandler) SuccessPage(c *gin.Context) {
	pID := c.Query("id")
	userID := c.Query("bookID")
	// fmt.Println(pID)
	// fmt.Println("Fully successful")

	c.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": pID,
		"userID":    userID,
	})
}

//SubStationsDetails function is used to fetch the details of the sub station.
func (uh *UserHandler) SubStationsDetails(c *gin.Context) {
	parent := c.Query("location")
	substations, err := uh.user.SubStationDetails(parent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Error while fetching sub stations.",
			"data":    err,
		})
		return
	}
	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Sub Stations retrieved successfully.",
		"data":    substations,
	})
}

//IndexPage function is used to display the index page
func (uh *UserHandler) IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// NewUserHandler is used to initialize the UserHandler
func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		user: userService,
	}
}
