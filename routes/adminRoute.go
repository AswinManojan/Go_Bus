package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	"gobus/server"
)

type AdminRouters struct {
	router *server.ServerStruct
	admin  *handlers.AdminHandler
	jwt    *middleware.JwtUtil
}

func (ar *AdminRouters) Routes() {
	ar.router.R.POST("/admin/login", ar.admin.Login)
	// ar.router.R.POST("/admin/stations/add", ar.jwt.ValidateToken("admin"), ar.admin.AddStation)
	// ar.router.R.GET("/admin/user_management/view/:id", ar.jwt.ValidateToken("admin"), ar.admin.FindUser)
	// ar.router.R.GET("/admin/user_management/view", ar.jwt.ValidateToken("admin"), ar.admin.FindAllUsers)
	// ar.router.R.GET("/admin/user_management/block/:id", ar.jwt.ValidateToken("admin"), ar.admin.BlockUser)
	// ar.router.R.GET("/admin/user_management/unblock/:id", ar.jwt.ValidateToken("admin"), ar.admin.UnBlockUser)
	// ar.router.R.GET("/admin/provider_management/view/:id", ar.jwt.ValidateToken("admin"), ar.admin.FindProvider)
	// ar.router.R.GET("/admin/provider_management/view", ar.jwt.ValidateToken("admin"), ar.admin.FindAllProvider)
	// ar.router.R.GET("/admin/provider_management/block/:id", ar.jwt.ValidateToken("admin"), ar.admin.BlockProvider)
	// ar.router.R.GET("/admin/provider_management/unblock/:id", ar.jwt.ValidateToken("admin"), ar.admin.UnBlockProvider)
	// ar.router.R.GET("/api/stations/view/:id", ar.jwt.ValidateToken("admin"), ar.admin.FindStation)
	// ar.router.R.GET("/api/stations/viewbyname", ar.jwt.ValidateToken("admin"), ar.admin.FindStationByName)
	// ar.router.R.GET("/api/stations/view", ar.jwt.ValidateToken("admin"), ar.admin.FindAllStations)
	// ar.router.R.PUT("/admin/user_management/edit/:id", ar.jwt.ValidateToken("admin"), ar.admin.UpdateUser)
	// ar.router.R.PUT("/admin/stations/edit/:id", ar.jwt.ValidateToken("admin"), ar.admin.UpdateStation)
	// ar.router.R.PUT("/admin/provider_management/edit/:id", ar.jwt.ValidateToken("admin"), ar.admin.UpdateProvider)
	// ar.router.R.DELETE("/admin/user_management/remove/:id", ar.jwt.ValidateToken("admin"), ar.admin.DeleteUser)
	// ar.router.R.DELETE("/admin/provider_management/remove/:id", ar.jwt.ValidateToken("admin"), ar.admin.DeleteProvider)
	// ar.router.R.DELETE("/admin/stations/remove/:id", ar.jwt.ValidateToken("admin"), ar.admin.DeleteStation)

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
	}
	// adminGroup.POST("/login", ar.admin.Login)
}

func NewAdminRoutes(a *handlers.AdminHandler, r *server.ServerStruct, jwt *middleware.JwtUtil) *AdminRouters {
	return &AdminRouters{
		router: r,
		admin:  a,
		jwt:    jwt,
	}
}
