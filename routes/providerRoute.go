package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	otphandlerprovider "gobus/otphandler_provider"
	"gobus/server"
)

// ProviderRouters struct is used to  intialize the provider router
type ProviderRouters struct {
	router   *server.Serverstruct
	provider *handlers.ProviderHandler
	jwt      *middleware.JwtUtil
	otp      *otphandlerprovider.OtpHandler
}

// ProRoutes function defines the provider routes.
func (pr *ProviderRouters) ProRoutes() {
	pr.router.R.POST("/provider/login", pr.provider.Login)
	// // pr.router.R.POST("/provider/register", pr.provider.RegisterProvider) // need otp verification
	pr.router.R.POST("/create_provider", pr.otp.GenerateOTP)
	pr.router.R.POST("/verify-provider", pr.otp.VerifyOTP)
	providerGroup := pr.router.R.Group("/provider").Use(pr.jwt.ValidateToken("provider"))
	{
		providerGroup.PUT("/edit_provider", pr.provider.EditProvider)
		providerGroup.GET("/station/view/:id", pr.provider.FindStationByID)
		providerGroup.GET("/station/view_name", pr.provider.FindStationByName)
		providerGroup.GET("/station/view", pr.provider.FindAllStations)
		providerGroup.POST("/bus/add", pr.provider.AddBus)
		providerGroup.GET("/bus/view", pr.provider.FindBus)
		providerGroup.GET("/bus/view/:id", pr.provider.FindBusByID)
		providerGroup.PUT("/edit_bus/:id", pr.provider.EditBus)
		providerGroup.DELETE("/delete_bus/:id", pr.provider.DeleteBus)
		providerGroup.GET("/coupon/view", pr.provider.FindCoupon)
		providerGroup.GET("/coupon/view/:id", pr.provider.FindCouponByID)
		providerGroup.POST("/coupon/add", pr.provider.AddCoupon)
		providerGroup.PUT("/edit_coupon/:id", pr.provider.EditCoupon)
		providerGroup.GET("/deactivate_coupon/:id", pr.provider.DeactivateCoupon)
		providerGroup.GET("/activate_coupon/:id", pr.provider.ActivateCoupon)
		providerGroup.GET("/coupon/view_code", pr.provider.FindCouponByCode)
		providerGroup.POST("station/add_sub_station", pr.provider.AddSubStations)
	}
}

// NewProviderRoutes function is used to instatiate the Provider Router
func NewProviderRoutes(a *handlers.ProviderHandler, server *server.Serverstruct, jwt *middleware.JwtUtil, o *otphandlerprovider.OtpHandler) *ProviderRouters {
	return &ProviderRouters{
		router:   server,
		provider: a,
		jwt:      jwt,
		otp:      o,
	}
}
