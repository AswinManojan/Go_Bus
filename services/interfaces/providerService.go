package interfaces

import (
	"gobus/dto"
	"gobus/entities"
)

type ProviderService interface {
	Login(loginRequest *dto.LoginRequest) (map[string]string, error)
	RegisterProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	FindProviderByEmail(email string) (*entities.ServiceProvider, error)
	EditProvider(email string, provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	FindStationById(id int) (*entities.Stations, error)
	FindStationByName(name string) (*entities.Stations, error)
	FindAllStations() ([]*entities.Stations, error)
	FindBus() ([]*entities.Buses, error)
	FindBusById(id int) (*entities.Buses, error)
	AddBus(bus *entities.Buses, email string) (*entities.Buses, error)
	EditBus(id int, bus *entities.Buses) (*entities.Buses, error)
	DeleteBus(id int, email string) (*entities.Buses, error)
	FindCoupon() ([]*entities.Coupons, error)
	FindCouponById(id int) (*entities.Coupons, error)
	AddCoupon(coupon *entities.Coupons) (*entities.Coupons, error)
	EditCoupon(id int, coupon *entities.Coupons) (*entities.Coupons, error)
	DeactivateCoupon(id int) (*entities.Coupons, error)
	ActivateCoupon(id int) (*entities.Coupons, error)
	FindCouponByCode(code string) (*entities.Coupons, error)
}
