package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	"gobus/server"
)

// AdminRouters struct is used to define the admin router.
type AdminRouters struct {
	router *server.Serverstruct
	admin  *handlers.AdminHandler
	jwt    *middleware.JwtUtil
}

// Routes function is used to define the admin routes.
func (ar *AdminRouters) Routes() {
	ar.router.R.POST("/admin/login", ar.admin.Login)
	adminGroup := ar.router.R.Group("/admin").Use(ar.jwt.ValidateToken("admin"))
	{
		adminGroup.POST("/stations/add", ar.admin.AddStation)
		adminGroup.GET("/user_management/view/:id", ar.admin.FindUser)
		adminGroup.GET("/user_management/view", ar.admin.FindAllUsers)
		adminGroup.GET("/user_management/block/:id", ar.admin.BlockUser)
		adminGroup.GET("/user_management/unblock/:id", ar.admin.UnBlockUser)
		adminGroup.GET("/provider_management/view/:id", ar.admin.FindProvider)
		adminGroup.GET("/provider_management/view", ar.admin.FindAllProvider)
		adminGroup.GET("/provider_management/block/:id", ar.admin.BlockProvider)
		adminGroup.GET("/provider_management/unblock/:id", ar.admin.UnBlockProvider)
		adminGroup.GET("/api/stations/view/:id", ar.admin.FindStation)
		adminGroup.GET("/api/stations/viewbyname", ar.admin.FindStationByName)
		adminGroup.GET("/api/stations/view", ar.admin.FindAllStations)
		adminGroup.PUT("/user_management/edit/:id", ar.admin.UpdateUser)
		adminGroup.PUT("/stations/edit/:id", ar.admin.UpdateStation)
		adminGroup.PUT("/provider_management/edit/:id", ar.admin.UpdateProvider)
		adminGroup.DELETE("/user_management/remove/:id", ar.admin.DeleteUser)
		adminGroup.DELETE("/provider_management/remove/:id", ar.admin.DeleteProvider)
		adminGroup.DELETE("/stations/remove/:id", ar.admin.DeleteStation)
		adminGroup.POST("/busschedule/addtochart", ar.admin.AddBusSchedule)
		adminGroup.POST("/busschedule/addbasefare", ar.admin.AddBaseFare)
		adminGroup.GET("/bookings/view", ar.admin.ViewAllBookings)
		adminGroup.GET("/bookings/viewbybus", ar.admin.ViewBookingsPerBus)
		adminGroup.POST("/bookings/cancelbus", ar.admin.CancelBus)
	}
	// adminGroup.POST("/login", ar.admin.Login)
}

// NewAdminRoutes function is used to instantiate Admin Routers.
func NewAdminRoutes(a *handlers.AdminHandler, r *server.Serverstruct, jwt *middleware.JwtUtil) *AdminRouters {
	return &AdminRouters{
		router: r,
		admin:  a,
		jwt:    jwt,
	}
}
