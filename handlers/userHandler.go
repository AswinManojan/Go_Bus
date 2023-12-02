package handlers

import (
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user interfaces.UserService
}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	user := &entities.User{}
	c.BindJSON(user)
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

func (uh *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Welcome to home page",
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)

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

func (uh *UserHandler) FindBus(c *gin.Context) {
	BusRequest := &dto.BusRequest{}
	c.BindJSON(BusRequest)
	buses, err := uh.user.FindBus(BusRequest)
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

func (uh *UserHandler) AddPassenger(c *gin.Context) {
	pass := &entities.PassengerInfo{}
	c.BindJSON(pass)
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
func (uh *UserHandler) BookSeat(c *gin.Context) {
	bookreq := &dto.BookingRequest{}
	c.BindJSON(bookreq)
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
		"message": "Booked in progress,awaiting payment",
		"status":  "Success",
		"data":    booking,
	})
}
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

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		user: userService,
	}
}
