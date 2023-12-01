package di

import (
	"gobus/db"
	"gobus/handlers"
	"gobus/middleware"
	"gobus/otphandler"
	otphandlerprovider "gobus/otphandler_provider"
	"gobus/repository"
	"gobus/routes"
	"gobus/server"
	"gobus/services"
)

func Init() *server.ServerStruct {
	Loadenv()
	db := db.ConnectDB()
	jwt := middleware.NewJwtUtil()
	// entities.SeatLayoutStr(db)
	otphandler.InitRedis()
	otphandlerprovider.InitRedis()
	userRepository := repository.NewUserRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	providerRepository := repository.NewProviderRepository(db)
	userService := services.NewUserService(userRepository, jwt)
	adminService := services.NewAdminService(adminRepository, jwt)
	providerService := services.NewProviderService(providerRepository, jwt)
	userHandler := handlers.NewUserHandler(userService)
	adminHandler := handlers.NewAdminHandler(adminService)
	providerHandler := handlers.NewProviderHandler(providerService)
	otpHandler := otphandler.NewotpHandler(userService)
	otpproviderHandler := otphandlerprovider.NewotpHandler(providerService)
	server := server.NewServer()
	userRoutes := routes.NewUserRoutes(userHandler, server, jwt, otpHandler)
	adminRoutes := routes.NewAdminRoutes(adminHandler, server, jwt)
	providerRoutes := routes.NewProviderRoutes(providerHandler, server, jwt, otpproviderHandler)
	adminRoutes.Routes()
	userRoutes.URoutes()
	providerRoutes.ProRoutes()
	return server
}
