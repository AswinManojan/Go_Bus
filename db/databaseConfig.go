package db

import (
	"gobus/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=gobus port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to DB")
	}
	db.AutoMigrate(&entities.User{},
		&entities.ServiceProvider{},
		&entities.Buses{},
		&entities.BusSchedule{},
		&entities.Coupons{},
		&entities.BookPassenger{},
		&entities.Booking{},
		&entities.BusSeatLayout{},
		&entities.BusStatus{},
		&entities.PassengerInfo{},
		&entities.Schedule{},
		&entities.BusType{},
		&entities.Stations{},
		&entities.BaseFare{})
	return db
}
