package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	"gobus/otphandler"
	"gobus/server"
)

type UserRouters struct {
	router *server.ServerStruct
	user   *handlers.UserHandler
	jwt    *middleware.JwtUtil
	otp    *otphandler.OtpHandler
}

func (as *UserRouters) URoutes() {
	as.router.R.POST("/generate-otp", as.otp.GenerateOTP)
	as.router.R.POST("/verify-otp", as.otp.VerifyOTP)
	// as.router.R.POST("/user/register", as.user.RegisterUser)
	as.router.R.POST("/user/login", as.user.Login)
	as.router.R.GET("/user/home", as.jwt.ValidateToken("user"), as.user.Home)
}

func NewUserRoutes(a *handlers.UserHandler, server *server.ServerStruct, jwt *middleware.JwtUtil, o *otphandler.OtpHandler) *UserRouters {
	return &UserRouters{
		router: server,
		user:   a,
		jwt:    jwt,
		otp:    o,
	}
}
