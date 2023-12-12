package handlers

import (
	"fmt"
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// AdminHandler struct is used to setup Admin Handler
type AdminHandler struct {
	admin interfaces.AdminService
}

// Login function is used for admin login purpose.
func (ah *AdminHandler) Login(c *gin.Context) {
	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)
	if err := validate.Struct(LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	token, err := ah.admin.Login(LoginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to login",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Admin logged in successfully",
		"data":    token,
	})
}

// FindUser function is used to find the user based on ID
func (ah *AdminHandler) FindUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid user ID",
			"data":    err.Error(),
		})
		return
	}
	user, err := ah.admin.FindUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the user",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "User has been found",
		"data":    user,
	})
}

// FindAllUsers function is used to find all the users of the application.
func (ah *AdminHandler) FindAllUsers(c *gin.Context) {
	users, err := ah.admin.FindAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to fetch the users",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Users has been found",
		"data":    users,
	})
}

// UpdateUser is used to update the user information.
func (ah *AdminHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid user ID provided",
			"data":    err.Error(),
		})
		return
	}
	user := &entities.User{}
	err = c.BindJSON(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind the user data",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	user, err = ah.admin.UpdateUser(idInt, *user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to update the user",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser function is used to delete a specific user from the application.
func (ah *AdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "",
			"data":    err.Error(),
		})
		return
	}
	user, err := ah.admin.DeleteUser(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to delete the user",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "User has been deleted successfully",
		"data":    user,
	})
}

// BlockUser function is used to restrict a users access to the application
func (ah *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid user ID provided",
			"data":    err.Error(),
		})
		return
	}
	user, err := ah.admin.BlockUser(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to block the user",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "User has been blocked",
		"data":    user,
	})
}

// UnBlockUser function is used to allow a restricted user to access into the application
func (ah *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid user ID provided",
			"data":    err.Error(),
		})
		return
	}
	user, err := ah.admin.UnBlockUser(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to unblock the user",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "User has been unblocked successfully",
		"data":    user,
	})
}

// FindProvider function is used to find the provider based on the id passed.
func (ah *AdminHandler) FindProvider(c *gin.Context) {
	id := c.Param("id")
	providerID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid provider ID provided",
			"data":    err.Error(),
		})
		return
	}
	provider, err := ah.admin.FindProvider(providerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to fetch the providers",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Provider has been found successfully",
		"data":    provider,
	})
}

// FindAllProvider is used to find the details of all the providers.
func (ah *AdminHandler) FindAllProvider(c *gin.Context) {
	providers, err := ah.admin.FindAllProvider()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the providers",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the providers",
		"data":    providers,
	})
}

// UpdateProvider function is used to update the provider information based on the id passed.
func (ah *AdminHandler) UpdateProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid provider ID provided",
			"data":    err.Error(),
		})
		return
	}
	provider := &entities.ServiceProvider{}
	err = c.BindJSON(provider)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind the provider info",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}

	provider, err = ah.admin.UpdateProvider(idInt, *provider)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to update the provider info",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully updated the provider info",
		"data":    provider,
	})
}

// DeleteProvider function is used to delete the provider
func (ah *AdminHandler) DeleteProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid provider ID provided",
			"data":    err.Error(),
		})
		return
	}
	provider, err := ah.admin.DeleteProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to delete the provider",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Provider deleted successfuly",
		"data":    provider,
	})
}

// BlockProvider function is used to block the provider from accessing the application.
func (ah *AdminHandler) BlockProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid provider ID provided",
			"data":    err.Error(),
		})
		return
	}
	provider, err := ah.admin.BlockProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to block the provider",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Succussfully blocked the provider",
		"data":    provider,
	})
}

// UnBlockProvider function is used to allow the provider to access the application.
func (ah *AdminHandler) UnBlockProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid provider ID provided",
			"data":    err.Error(),
		})
		return
	}
	provider, err := ah.admin.UnBlockProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to unblock the provider",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully unblocked the provider",
		"data":    provider,
	})
}

// FindStation function is used to find the station based on the ID
func (ah *AdminHandler) FindStation(c *gin.Context) {
	id := c.Param("id")
	stationID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid station ID",
			"data":    err.Error(),
		})
		return
	}
	station, err := ah.admin.FindStation(stationID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the station",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully found the stations",
		"data":    station,
	})
}

// FindStationByName function is used to find the station based on the name
func (ah *AdminHandler) FindStationByName(c *gin.Context) {
	name := c.Query("name")
	fmt.Print(name)
	station, err := ah.admin.FindStationByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unavble to find the station",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the station",
		"data":    station,
	})
}

// FindAllStations function is used to find all the stations
func (ah *AdminHandler) FindAllStations(c *gin.Context) {
	stations, err := ah.admin.FindAllStations()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the stations",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully fetched the stations",
		"data":    stations,
	})
}

// UpdateStation function is used to update the station
func (ah *AdminHandler) UpdateStation(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Station ID provided",
			"data":    err.Error(),
		})
		return
	}
	station := &entities.Stations{}
	err = c.BindJSON(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind station info",
			"data":    err.Error(),
		})
		return
	}

	if err := validate.Struct(station); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	station, err = ah.admin.UpdateStation(idInt, *station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to update the station",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully updated the station info",
		"data":    station,
	})
}

// DeleteStation function is used to delete a station
func (ah *AdminHandler) DeleteStation(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid station ID provided",
			"data":    err.Error(),
		})
		return
	}
	station, err := ah.admin.DeleteStation(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to delete the station",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Deleted the station successfully",
		"data":    station,
	})
}

// AddStation function is used to add new station to application.
func (ah *AdminHandler) AddStation(c *gin.Context) {
	station := &entities.Stations{}
	err := c.BindJSON(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind the station",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(station); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	addedStation, err := ah.admin.AddStation(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add station",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the station",
		"data":    addedStation,
	})

}

// AddBaseFare function is used to add the base fare for a route
func (ah *AdminHandler) AddBaseFare(c *gin.Context) {
	baseFare := &entities.BaseFare{}
	err := c.BindJSON(baseFare)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind the baseFare",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(baseFare); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	addedFare, err := ah.admin.AddFareForRoute(baseFare)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add baseFare",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the baseFare",
		"data":    addedFare,
	})

}

// AddBusSchedule function is used to add bus schedule for the bus
func (ah *AdminHandler) AddBusSchedule(c *gin.Context) {
	schedule := &dto.BusSchedule{}
	err := c.BindJSON(schedule)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to bind the schedule",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	addedschedule, err := ah.admin.AddBusSchedule(schedule)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add schedule",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the schedule",
		"data":    addedschedule,
	})
}

// ViewAllBookings function is used to list all the bookings
func (ah *AdminHandler) ViewAllBookings(c *gin.Context) {
	bookings, err := ah.admin.ViewAllBookings()
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
		"message": "Successfully fetched the bookings",
		"data":    bookings,
	})
}

// ViewBookingsPerBus function is used to list all bookings based on the bus id passed.
func (ah *AdminHandler) ViewBookingsPerBus(c *gin.Context) {
	schedule := &dto.BusSchedule{}
	c.BindJSON(schedule)
	bookings, err := ah.admin.ViewBookingsPerBus(int(schedule.BusID), schedule.Day)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the bookings",
			"data":    err.Error(),
		})
		return
	}
	if err := validate.Struct(schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully fetched the bookings",
		"data":    bookings,
	})
}

// CancelBus function is used to cancel a bus on a specific day.
func (ah *AdminHandler) CancelBus(c *gin.Context) {
	schedule := &dto.BusSchedule{}
	c.BindJSON(schedule)
	result, err := ah.admin.CancelBus(int(schedule.BusID), schedule.Day)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to cancel the bus",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully cancelled the bus",
		"data":    result,
	})
}

// NewAdminHandler is used to initialize the AdminHandler
func NewAdminHandler(adminService interfaces.AdminService) *AdminHandler {
	return &AdminHandler{
		admin: adminService,
	}
}
