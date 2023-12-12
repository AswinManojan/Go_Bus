package handlers

import (
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProviderHandler struct is used to setup Provider Handler
type ProviderHandler struct {
	provider interfaces.ProviderService
}

// Login function is used for provider login purpose.
func (ph *ProviderHandler) Login(c *gin.Context) {
	loginRequest := &dto.LoginRequest{}
	c.BindJSON(loginRequest)
	if err := validate.Struct(loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	token, err := ph.provider.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to Login",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Provider Login successful",
		"data":    token,
	})
}

// RegisterProvider fucntion is used to register the provider into the application.
func (ph *ProviderHandler) RegisterProvider(c *gin.Context) {
	provider := &entities.ServiceProvider{}
	c.BindJSON(provider)
	if err := validate.Struct(provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	regProvider, err := ph.provider.RegisterProvider(provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to register the user",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "User has been registered",
		"data":    regProvider,
	})
}

// EditProvider function is used to edit the bus provider details.
func (ph *ProviderHandler) EditProvider(c *gin.Context) {
	provider := &entities.ServiceProvider{}
	c.BindJSON(provider)
	if err := validate.Struct(provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	email := c.MustGet("email").(string)
	editedProvider, err := ph.provider.EditProvider(email, provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to edit the provider",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Provider data edited successfully",
		"data":    editedProvider,
	})
}

// FindStationByID is used to find the station based on the ID passed.
func (ph *ProviderHandler) FindStationByID(c *gin.Context) {
	id := c.Param("id")
	stationID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid station ID provided",
			"data":    err.Error(),
		})
		return
	}
	station, err := ph.provider.FindStationByID(stationID)
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
		"message": "Successfully found the station",
		"data":    station,
	})
}

// FindStationByName function is used to find the station based on the name.
func (ph *ProviderHandler) FindStationByName(c *gin.Context) {
	name := c.Query("name")
	station, err := ph.provider.FindStationByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to find the station",
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
func (ph *ProviderHandler) FindAllStations(c *gin.Context) {
	stations, err := ph.provider.FindAllStations()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to fetch the station",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the stations",
		"data":    stations,
	})
}

// FindBus is used to find the bus based all the buses
func (ph *ProviderHandler) FindBus(c *gin.Context) {
	buses, err := ph.provider.FindBus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to fetch the buses",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status":  "Success",
		"message": "Successfully found the buses",
		"data":    buses,
	})
}

// FindBusByID is used to find the bus based on the ID
func (ph *ProviderHandler) FindBusByID(c *gin.Context) {
	id := c.Param("id")
	busID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid bus ID",
			"data":    err.Error(),
		})
		return
	}
	bus, err := ph.provider.FindBusByID(busID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to fetch the bus info.",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully found the bus",
		"data":    bus,
	})
}

// EditBus function is used to edit the bus information.
func (ph *ProviderHandler) EditBus(c *gin.Context) {
	id := c.Param("id")
	busID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Bus ID",
			"data":    err.Error(),
		})
		return
	}
	bus := &entities.Buses{}
	c.BindJSON(bus)
	if err := validate.Struct(bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	editedBus, err := ph.provider.EditBus(busID, bus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to edit the bus info",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Successfully edited the bus info",
		"data":    editedBus,
	})
}

// DeleteBus function is used to delete the bus
func (ph *ProviderHandler) DeleteBus(c *gin.Context) {
	id := c.Param("id")
	busID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Bus ID",
			"data":    err.Error(),
		})
		return
	}
	email := c.MustGet("email").(string)
	deletedBus, err := ph.provider.DeleteBus(busID, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to delete the bus",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Successfully deleted the bus",
		"data":    deletedBus,
	})
}

// FindCoupon is used to find all the coupons
func (ph *ProviderHandler) FindCoupon(c *gin.Context) {
	coupons, err := ph.provider.FindCoupon()
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

// FindCouponByID is used to Find the coupons based on the ID
func (ph *ProviderHandler) FindCouponByID(c *gin.Context) {
	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid station ID",
			"data":    err.Error(),
		})
		return
	}
	coupon, err := ph.provider.FindCouponByID(couponID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to find the coupon",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully found the coupon",
		"data":    coupon,
	})
}

// AddCoupon function is used to add new coupon.
func (ph *ProviderHandler) AddCoupon(c *gin.Context) {
	coupon := &entities.Coupons{}
	c.BindJSON(coupon)
	if err := validate.Struct(coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	coupon, err := ph.provider.AddCoupon(coupon)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add a new coupon",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added a new coupon",
		"data":    coupon,
	})
}

// AddBus function is used to add new bus
func (ph *ProviderHandler) AddBus(c *gin.Context) {
	bus := &entities.Buses{}
	c.BindJSON(bus)
	if err := validate.Struct(bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	email := c.MustGet("email").(string)
	bus, err := ph.provider.AddBus(bus, email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to add a new bus",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the bus",
		"data":    bus,
	})
}

// EditCoupon function is used edit the coupon based on the id
func (ph *ProviderHandler) EditCoupon(c *gin.Context) {

	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Coupon ID",
			"data":    err.Error(),
		})
		return
	}
	coupon := &entities.Coupons{}
	c.BindJSON(coupon)
	if err := validate.Struct(coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Please fill all the mandatory fields.",
			"data":    err.Error(),
		})
	}
	editedCoupon, err := ph.provider.EditCoupon(couponID, coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to edit the coupon info",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Successfully edited the coupon info",
		"data":    editedCoupon,
	})
}

// DeactivateCoupon function is used to set the coupon to inactive state
func (ph *ProviderHandler) DeactivateCoupon(c *gin.Context) {
	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Coupon ID",
			"data":    err.Error(),
		})
		return
	}
	deletedCoupon, err := ph.provider.DeactivateCoupon(couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to deactivate the coupon",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Successfully deactivated the coupon",
		"data":    deletedCoupon,
	})
}

// ActivateCoupon function is used to set the coupon to active state
func (ph *ProviderHandler) ActivateCoupon(c *gin.Context) {
	id := c.Param("id")
	couponID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Invalid Coupon ID",
			"data":    err.Error(),
		})
		return
	}
	deletedCoupon, err := ph.provider.ActivateCoupon(couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Unable to activate the coupon",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "Success",
		"message": "Successfully activated the coupon",
		"data":    deletedCoupon,
	})
}

// FindCouponByCode function is used to find the coupon based on the code
func (ph *ProviderHandler) FindCouponByCode(c *gin.Context) {
	code := c.Query("code")
	coupon, err := ph.provider.FindCouponByCode(code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to find the coupon",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully found the coupon",
		"data":    coupon,
	})
}

//AddSubStations function is used to add the sub stations
func (ph *ProviderHandler) AddSubStations(c *gin.Context) {
	subStation := &entities.SubStation{}
	c.BindJSON(subStation)
	station, err := ph.provider.AddSubStations(subStation)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Failed to add the sub station.",
			"data":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "Success",
		"message": "Successfully added the substation.",
		"data":    station,
	})
}

// NewProviderHandler is used to initialize the ProviderHandler
func NewProviderHandler(providerService interfaces.ProviderService) *ProviderHandler {
	return &ProviderHandler{
		provider: providerService,
	}

}
