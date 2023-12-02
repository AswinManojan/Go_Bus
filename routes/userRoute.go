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
	as.router.R.POST("/create_user", as.otp.GenerateOTP)
	as.router.R.POST("/verify-user", as.otp.VerifyOTP)
	// as.router.R.POST("/user/register", as.user.RegisterUser)
	as.router.R.POST("/user/login", as.user.Login)
	as.router.R.GET("/user/home", as.jwt.ValidateToken("user"), as.user.Home)
	as.router.R.GET("/user/findbus", as.jwt.ValidateToken("user"), as.user.FindBus)
	as.router.R.POST("/user/addpassenger", as.jwt.ValidateToken("user"), as.user.AddPassenger)
	as.router.R.GET("/user/viewallpassenger", as.jwt.ValidateToken("user"), as.user.ViewAllPassengers)
	as.router.R.POST("/user/bookseat", as.jwt.ValidateToken("user"), as.user.BookSeat)
	as.router.R.GET("/user/coupon/view", as.jwt.ValidateToken("user"), as.user.FindCoupon)
	as.router.R.GET("/user/bookings/view", as.jwt.ValidateToken("user"), as.user.ViewBookings)
}

func NewUserRoutes(a *handlers.UserHandler, server *server.ServerStruct, jwt *middleware.JwtUtil, o *otphandler.OtpHandler) *UserRouters {
	return &UserRouters{
		router: server,
		user:   a,
		jwt:    jwt,
		otp:    o,
	}
}
