package handlers

import (
	"gobus/dto"
	"gobus/entities"
	"gobus/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	admin interfaces.AdminService
}

func (ah *AdminHandler) Login(c *gin.Context) {
	LoginRequest := &dto.LoginRequest{}
	c.BindJSON(LoginRequest)
	token, err := ah.admin.Login(LoginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
}
func (ah *AdminHandler) FindUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	user, err := ah.admin.FindUser(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
func (ah *AdminHandler) FindAllUsers(c *gin.Context) {
	users, err := ah.admin.FindAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": users,
	})
}
func (ah *AdminHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := &entities.User{}
	err = c.BindJSON(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = ah.admin.UpdateUser(idInt, *user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
func (ah *AdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := ah.admin.DeleteUser(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": user,
	})
}
func (ah *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := ah.admin.BlockUser(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
func (ah *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := ah.admin.UnBlockProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
func (ah *AdminHandler) FindProvider(c *gin.Context) {
	id := c.Param("id")
	providerID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid provider ID",
		})
		return
	}
	provider, err := ah.admin.FindProvider(providerID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": provider,
	})
}
func (ah *AdminHandler) FindAllProvider(c *gin.Context) {
	providers, err := ah.admin.FindAllProvider()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": providers,
	})
}
func (ah *AdminHandler) UpdateProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	provider := &entities.ServiceProvider{}
	err = c.BindJSON(provider)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	provider, err = ah.admin.UpdateProvider(idInt, *provider)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": provider,
	})
}
func (ah *AdminHandler) DeleteProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	provider, err := ah.admin.DeleteProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": provider,
	})
}
func (ah *AdminHandler) BlockProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	provider, err := ah.admin.BlockProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": provider,
	})
}
func (ah *AdminHandler) UnBlockProvider(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	provider, err := ah.admin.UnBlockProvider(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": provider,
	})
}
func (ah *AdminHandler) FindStation(c *gin.Context) {
	id := c.Param("id")
	stationID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid station ID",
		})
		return
	}
	station, err := ah.admin.FindStation(stationID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": station,
	})
}
func (ah *AdminHandler) FindStationByName(c *gin.Context) {
	name := c.Query("name")
	station, err := ah.admin.FindStationByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": station,
	})
}
func (ah *AdminHandler) FindAllStations(c *gin.Context) {
	stations, err := ah.admin.FindAllStations()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": stations,
	})
}
func (ah *AdminHandler) UpdateStation(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	station := &entities.Stations{}
	err = c.BindJSON(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	station, err = ah.admin.UpdateStation(idInt, *station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": station,
	})
}
func (ah *AdminHandler) DeleteStation(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	station, err := ah.admin.DeleteStation(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": station,
	})
}
func (ah *AdminHandler) AddStation(c *gin.Context) {
	station := &entities.Stations{}
	err := c.BindJSON(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	addedStation, err := ah.admin.AddStation(station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": addedStation,
	})

}

func NewAdminHandler(adminService interfaces.AdminService) *AdminHandler {
	return &AdminHandler{
		admin: adminService,
	}
}
