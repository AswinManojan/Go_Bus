package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	"gobus/otphandler"
	"gobus/server"
)

// UserRouters struct is used to  intialize the user router
type UserRouters struct {
	router *server.Serverstruct
	user   *handlers.UserHandler
	jwt    *middleware.JwtUtil
	otp    *otphandler.OtpHandler
}

// URoutes function defines the user routes.
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
	as.router.R.GET("/user/payment/:bookid", as.user.MakePayment)
	as.router.R.GET("/user/payment/success", as.user.PaymentSuccess)
	as.router.R.GET("/user/coupon/view", as.jwt.ValidateToken("user"), as.user.FindCoupon)
	as.router.R.GET("/user/bookings/view", as.jwt.ValidateToken("user"), as.user.ViewBookings)
	as.router.R.POST("/user/bookings/cancel/:id", as.jwt.ValidateToken("user"), as.user.CancelBooking)
	as.router.R.GET("/user/seatstatus", as.jwt.ValidateToken("user"), as.user.SeatStatus)
	as.router.R.GET("/success", as.user.SuccessPage)
	as.router.R.GET("/user/getsubstationlist", as.user.SubStationsDetails)
	as.router.R.GET("/", as.user.IndexPage)
}

// NewUserRoutes function will return the pointer of UserRouters
func NewUserRoutes(a *handlers.UserHandler, server *server.Serverstruct, jwt *middleware.JwtUtil, o *otphandler.OtpHandler) *UserRouters {
	return &UserRouters{
		router: server,
		user:   a,
		jwt:    jwt,
		otp:    o,
	}
}
