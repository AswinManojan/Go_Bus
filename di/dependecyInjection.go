package di

import (
	"gobus/db"
	"gobus/handlers"
	"gobus/middleware"
	"gobus/otphandler"
	"gobus/repository"
	"gobus/routes"
	"gobus/server"
	"gobus/services"
)

func Init() *server.ServerStruct {
	Loadenv()
	db := db.ConnectDB()
	jwt := middleware.NewJwtUtil()
	otphandler.InitRedis()
	userRepository := repository.NewUserRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	userService := services.NewUserService(userRepository, jwt)
	adminService := services.NewAdminService(adminRepository, jwt)
	userHandler := handlers.NewUserHandler(userService)
	adminHandler := handlers.NewAdminHandler(adminService)
	otpHandler := otphandler.NewotpHandler(userService)
	server := server.NewServer()
	userRoutes := routes.NewUserRoutes(userHandler, server, jwt, otpHandler)
	adminRoutes := routes.NewAdminRoutes(adminHandler, server, jwt)
	adminRoutes.Routes()

	userRoutes.URoutes()
	return server
}
