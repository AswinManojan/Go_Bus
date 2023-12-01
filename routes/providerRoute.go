package routes

import (
	"gobus/handlers"
	"gobus/middleware"
	otphandlerprovider "gobus/otphandler_provider"
	"gobus/server"
)

type ProviderRouters struct {
	router   *server.ServerStruct
	provider *handlers.ProviderHandler
	jwt      *middleware.JwtUtil
	otp      *otphandlerprovider.OtpHandler
}

func (pr *ProviderRouters) ProRoutes() {
	pr.router.R.POST("/provider/login", pr.provider.Login)
	// // pr.router.R.POST("/provider/register", pr.provider.RegisterProvider) // need otp verification
	pr.router.R.POST("/create_provider", pr.otp.GenerateOTP)
	pr.router.R.POST("/verify-provider", pr.otp.VerifyOTP)
	// pr.router.R.PUT("/provider/edit_provider", pr.jwt.ValidateToken("provider"), pr.provider.EditProvider)
	// pr.router.R.GET("/provider/station/view/:id", pr.jwt.ValidateToken("provider"), pr.provider.FindStationById)
	// pr.router.R.GET("/provider/station/view_name", pr.jwt.ValidateToken("provider"), pr.provider.FindStationByName)
	// pr.router.R.GET("/provider/station/view", pr.jwt.ValidateToken("provider"), pr.provider.FindAllStations)
	// pr.router.R.POST("/provider/bus/add", pr.jwt.ValidateToken("provider"), pr.provider.AddBus)
	// pr.router.R.GET("/provider/bus/view", pr.jwt.ValidateToken("provider"), pr.provider.FindBus)
	// pr.router.R.GET("/provider/bus/view/:id", pr.jwt.ValidateToken("provider"), pr.provider.FindBusById)
	// pr.router.R.PUT("/provider/edit_bus/:id", pr.jwt.ValidateToken("provider"), pr.provider.EditBus)
	// pr.router.R.DELETE("/provider/delete_bus/:id", pr.jwt.ValidateToken("provider"), pr.provider.DeleteBus)
	// pr.router.R.GET("/provider/coupon/view", pr.jwt.ValidateToken("provider"), pr.provider.FindCoupon)
	// pr.router.R.GET("/provider/coupon/view/:id", pr.jwt.ValidateToken("provider"), pr.provider.FindCouponById)
	// pr.router.R.POST("/provider/coupon/add", pr.jwt.ValidateToken("provider"), pr.provider.AddCoupon)
	// pr.router.R.PUT("/provider/edit_coupon/:id", pr.jwt.ValidateToken("provider"), pr.provider.EditCoupon)
	// pr.router.R.GET("/provider/deactivate_coupon/:id", pr.jwt.ValidateToken("provider"), pr.provider.DeactivateCoupon)
	// pr.router.R.GET("/provider/activate_coupon/:id", pr.jwt.ValidateToken("provider"), pr.provider.ActivateCoupon)
	// pr.router.R.GET("/provider/coupon/view_code", pr.jwt.ValidateToken("provider"), pr.provider.FindCouponByCode)
	providerGroup := pr.router.R.Group("/provider").Use(pr.jwt.ValidateToken("provider"))
	{
		providerGroup.PUT("/edit_provider", pr.provider.EditProvider)
		providerGroup.GET("/station/view/:id", pr.provider.FindStationById)
		providerGroup.GET("/station/view_name", pr.provider.FindStationByName)
		providerGroup.GET("/station/view", pr.provider.FindAllStations)
		providerGroup.POST("/bus/add", pr.provider.AddBus)
		providerGroup.GET("/bus/view", pr.provider.FindBus)
		providerGroup.GET("/bus/view/:id", pr.provider.FindBusById)
		providerGroup.PUT("/edit_bus/:id", pr.provider.EditBus)
		providerGroup.DELETE("/delete_bus/:id", pr.provider.DeleteBus)
		providerGroup.GET("/coupon/view", pr.provider.FindCoupon)
		providerGroup.GET("/coupon/view/:id", pr.provider.FindCouponById)
		providerGroup.POST("/coupon/add", pr.provider.AddCoupon)
		providerGroup.PUT("/edit_coupon/:id", pr.provider.EditCoupon)
		providerGroup.GET("/deactivate_coupon/:id", pr.provider.DeactivateCoupon)
		providerGroup.GET("/activate_coupon/:id", pr.provider.ActivateCoupon)
		providerGroup.GET("/coupon/view_code", pr.provider.FindCouponByCode)
	}
}

func NewProviderRoutes(a *handlers.ProviderHandler, server *server.ServerStruct, jwt *middleware.JwtUtil, o *otphandlerprovider.OtpHandler) *ProviderRouters {
	return &ProviderRouters{
		router:   server,
		provider: a,
		jwt:      jwt,
		otp:      o,
	}
}
