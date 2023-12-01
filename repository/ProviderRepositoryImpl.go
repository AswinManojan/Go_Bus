package repository

import (
	"errors"
	"gobus/entities"
	"gobus/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type ProviderRepositoryImpl struct {
	DB *gorm.DB
}

// FindProviderById implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindProviderByEmail(email string) (*entities.ServiceProvider, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	provider := &entities.ServiceProvider{}
	result := pr.DB.Where("email=?", email).First(provider)
	if result.Error != nil {
		log.Println("Coupon doesn't exist")
		return nil, errors.New("no coupon found with this name")
	}
	return provider, nil
}

// FindCouponByCode implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindCouponByCode(code string) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	coupon := &entities.Coupons{}
	result := pr.DB.Where("coupon_code=?", code).First(coupon)
	if result.Error != nil {
		log.Println("Coupon doesn't exist")
		return nil, errors.New("no coupon found with this name")
	}
	return coupon, nil
}

// AddCoupon implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) AddCoupon(coupon *entities.Coupons) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	_, err := pr.FindCouponByCode(coupon.CouponCode)
	if err == nil {
		log.Println("Coupon ALREADY EXISTS")
		return nil, errors.New("coupon exists in db")
	}

	result := pr.DB.Create(&coupon)
	if result.Error != nil {
		log.Println("Unable to add coupon, AdminRepositoryImpl package")
		return nil, errors.New("coupon not added to db")
	}
	return coupon, nil
}

func (pr *ProviderRepositoryImpl) AddBus(bus *entities.Buses, email string) (*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	spr, err := pr.FindProviderByEmail(email)
	if bus.ProviderId == 0 {
		bus.ProviderId = spr.ProviderID
	}
	if spr.ProviderID != bus.ProviderId {
		log.Println("Unable to modify the details of a different provider")
		return nil, errors.New("access Restricted")
	}
	if err != nil {
		log.Println("Unable to find this provider")
		return nil, err
	}
	if _, err := pr.FindBusByNumber(bus.BusNumber); err == nil {
		log.Println("Bus ALREADY EXISTS")
		return nil, errors.New("bus exists in db")
	}

	result := pr.DB.Create(&bus)
	if result.Error != nil {
		log.Println("Unable to add bus, AdminRepositoryImpl package", result.Error)
		return nil, errors.New("bus not added to db")
	}
	return bus, nil
}

// DeleteBus implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) DeleteBus(id int, email string) (*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	spr, _ := pr.FindProviderByEmail(email)
	bus, err := pr.FindBusById(id)
	if spr.ProviderID != bus.ProviderId {
		log.Println("Unable to modify the details of a different provider")
		return nil, errors.New("access Restricted")
	}
	if err != nil {
		log.Println("Bus Not Found")
		return nil, errors.New("bus doesnot exists in db")
	}
	pr.DB.Delete(bus).Where("bus_id=?", id)
	return bus, nil
}

// DeactivateCoupon implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) DeactivateCoupon(id int) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundCoupon, err := pr.FindCouponById(id)
	if err != nil {
		log.Println("Coupon Not Found")
		return nil, errors.New("coupon doesnot exists in db")
	}
	foundCoupon.IsActive = false
	result := pr.DB.Save(&foundCoupon)
	if result.Error != nil {
		log.Println("Coupon Not Updated, AdminRepositoryImpl package")
		return nil, errors.New("coupon not updated")
	}
	return foundCoupon, nil
}
func (pr *ProviderRepositoryImpl) ActivateCoupon(id int) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundCoupon, err := pr.FindCouponById(id)
	if err != nil {
		log.Println("Coupon Not Found")
		return nil, errors.New("coupon doesnot exists in db")
	}
	foundCoupon.IsActive = true
	result := pr.DB.Save(&foundCoupon)
	if result.Error != nil {
		log.Println("Coupon Not Updated, AdminRepositoryImpl package")
		return nil, errors.New("coupon not updated")
	}
	return foundCoupon, nil
}

// EditBus implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) EditBus(id int, bus *entities.Buses) (*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundBus, err := pr.FindBusById(id)
	if err != nil {
		log.Println("Bus Not Found")
		return nil, errors.New("bus doesnot exists in db")
	}
	if bus.BusNumber != "" {
		foundBus.BusNumber = bus.BusNumber
	}
	if bus.BusTypeCode != "" {
		foundBus.BusTypeCode = bus.BusTypeCode
	}
	// if bus.BusStationId != 0 {
	// 	foundBus.BusStationId = bus.BusStationId
	// }
	if bus.ProviderId != 0 {
		foundBus.ProviderId = bus.ProviderId
	}
	if bus.ScheduleId != 0 {
		foundBus.ScheduleId = bus.ScheduleId
	}
	// if bus.SeatId != 0 {
	// 	foundBus.SeatId = bus.SeatId
	// }
	// if bus.TotalPushBackSeats != 0 {
	// 	foundBus.TotalPushBackSeats = bus.TotalPushBackSeats
	// }
	// if bus.TotalSleeperSeats != 0 {
	// 	foundBus.TotalSleeperSeats = bus.TotalSleeperSeats
	// }
	result := pr.DB.Save(&foundBus)
	if result.Error != nil {
		log.Println("Bus Not Updated, AdminRepositoryImpl package")
		return nil, errors.New("bus not updated")
	}
	return foundBus, nil
}

// EditCoupon implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) EditCoupon(id int, coupon *entities.Coupons) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundCoupon, err := pr.FindCouponById(id)
	if err != nil {
		log.Println("Coupon Not Found")
		return nil, errors.New("coupon doesnot exists in db")
	}
	if coupon.CouponCode != "" {
		foundCoupon.CouponCode = coupon.CouponCode
	}
	if coupon.ValidFrom != "" {
		foundCoupon.ValidFrom = coupon.ValidFrom
	}
	if coupon.ValidUpto != "" {
		foundCoupon.ValidUpto = coupon.ValidUpto
	}
	if coupon.Discount != 0 {
		foundCoupon.Discount = coupon.Discount
	}
	result := pr.DB.Save(&foundCoupon)
	if result.Error != nil {
		log.Println("Coupon Not Updated, AdminRepositoryImpl package")
		return nil, errors.New("coupon not updated")
	}
	return foundCoupon, nil

}

// EditProvider implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) EditProvider(email string, provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	foundProvider, _ := pr.FindProviderByEmail(email)
	if provider.Address != "" {
		foundProvider.Address = provider.Address
	}
	if provider.CompanyName != "" {
		foundProvider.CompanyName = provider.CompanyName
	}
	if provider.Email != "" {
		foundProvider.Email = provider.Email
	}
	if provider.Password != "" {
		foundProvider.Password = provider.Password
	}
	if provider.PhoneNumber != "" {
		foundProvider.PhoneNumber = provider.PhoneNumber
	}
	if provider.BusCount != 0 {
		foundProvider.BusCount = provider.BusCount
	}
	result := pr.DB.Save(&foundProvider)
	if result.Error != nil {
		log.Println("Coupon Not Updated, AdminRepositoryImpl package")
		return nil, errors.New("coupon not updated")
	}
	return foundProvider, nil
}

// FindAllStations implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindAllStations() ([]*entities.Stations, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	stations := []*entities.Stations{}
	result := pr.DB.Find(&stations)
	if result.Error != nil {
		return nil, result.Error
	}
	return stations, nil
}

// FindBus implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindBus() ([]*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	buses := []*entities.Buses{}
	result := pr.DB.Find(&buses)
	if result.Error != nil {
		return nil, result.Error
	}
	return buses, nil
}

// FindCoupon implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindCoupon() ([]*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	coupons := []*entities.Coupons{}
	result := pr.DB.Find(&coupons)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupons, nil
}

// FindCouponById implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindCouponById(id int) (*entities.Coupons, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	coupon := &entities.Coupons{}
	result := pr.DB.Where("coupon_id", id).First(coupon)
	if result.Error != nil {
		return nil, result.Error
	}
	return coupon, nil
}

// FindStationById implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindStationById(id int) (*entities.Stations, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station := &entities.Stations{}
	result := pr.DB.Where("station_id", id).First(station)
	if result.Error != nil {
		return nil, result.Error
	}
	return station, nil
}

// FindStationByName implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindStationByName(name string) (*entities.Stations, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	station := &entities.Stations{}
	result := pr.DB.Where("station_name", name).First(station)
	if result.Error != nil {
		return nil, result.Error
	}
	return station, nil
}

func (pr *ProviderRepositoryImpl) FindBusByNumber(number string) (*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bus := &entities.Buses{}
	result := pr.DB.Where("bus_number", number).First(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

// FundBusById implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) FindBusById(id int) (*entities.Buses, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB")
		return nil, errors.New("error connecting database")
	}
	bus := &entities.Buses{}
	result := pr.DB.Where("bus_id", id).First(bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return bus, nil
}

// RegisterProvider implements interfaces.ProviderRepository.
func (pr *ProviderRepositoryImpl) RegisterProvider(provider *entities.ServiceProvider) (*entities.ServiceProvider, error) {
	if pr.DB == nil {
		log.Println("Error connecting DB, ProviderRepositoryImpl package")
		return nil, errors.New("error Connecting Database")
	}

	result := pr.DB.Create(&provider)
	if result.Error != nil {
		log.Println("Unable to add user, ProviderRepositoryImpl package")
		return provider, errors.New("user already exists")
	}

	return provider, nil
}

func NewProviderRepository(db *gorm.DB) interfaces.ProviderRepository {
	return &ProviderRepositoryImpl{
		DB: db,
	}
}
