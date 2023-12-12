package interfaces

import "gobus/entities"

// ProviderRepository interface is the interface used for provider repository
type ProviderRepository interface {
	RegisterProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	FindProviderByEmail(email string) (*entities.ServiceProvider, error)
	EditProvider(email string, provider *entities.ServiceProvider) (*entities.ServiceProvider, error)
	FindStationByID(id int) (*entities.Stations, error)
	FindStationByName(name string) (*entities.Stations, error)
	FindAllStations() ([]*entities.Stations, error)
	FindBus() ([]*entities.Buses, error)
	FindBusByID(id int) (*entities.Buses, error)
	EditBus(id int, bus *entities.Buses) (*entities.Buses, error)
	DeleteBus(id int, email string) (*entities.Buses, error)
	FindCoupon() ([]*entities.Coupons, error)
	FindCouponByID(id int) (*entities.Coupons, error)
	AddCoupon(coupon *entities.Coupons) (*entities.Coupons, error)
	EditCoupon(id int, coupon *entities.Coupons) (*entities.Coupons, error)
	DeactivateCoupon(id int) (*entities.Coupons, error)
	ActivateCoupon(id int) (*entities.Coupons, error)
	FindCouponByCode(code string) (*entities.Coupons, error)
	AddBus(bus *entities.Buses, email string) (*entities.Buses, error)
	FindBusByNumber(number string) (*entities.Buses, error)
	AddSubStations(station *entities.SubStation) (*entities.SubStation, error)
}
