package handlers

import (
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	provider interfaces.ProviderService
}

func (ph *ProviderHandler) Login(c *gin.Context) {
	loginRequest := &dto.LoginRequest{}
	c.BindJSON(loginRequest)
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
func (ph *ProviderHandler) RegisterProvider(c *gin.Context) {
	provider := &entities.ServiceProvider{}
	c.BindJSON(provider)
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

func (ph *ProviderHandler) EditProvider(c *gin.Context) {
	provider := &entities.ServiceProvider{}
	c.BindJSON(provider)
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

func (ph *ProviderHandler) FindStationById(c *gin.Context) {
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
	station, err := ph.provider.FindStationById(stationID)
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
func (ph *ProviderHandler) FindBusById(c *gin.Context) {
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
	bus, err := ph.provider.FindBusById(busID)
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
func (ph *ProviderHandler) FindCouponById(c *gin.Context) {
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
	coupon, err := ph.provider.FindCouponById(couponID)
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
func (ph *ProviderHandler) AddCoupon(c *gin.Context) {
	coupon := &entities.Coupons{}
	c.BindJSON(coupon)
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
func (ph *ProviderHandler) AddBus(c *gin.Context) {
	bus := &entities.Buses{}
	c.BindJSON(bus)
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

func NewProviderHandler(providerService interfaces.ProviderService) *ProviderHandler {
	return &ProviderHandler{
		provider: providerService,
	}

}
