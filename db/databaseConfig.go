package db

import (
	"gobus/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_CONFIG")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to DB")
	}
	db.AutoMigrate(&entities.User{},
		&entities.ServiceProvider{},
		&entities.Buses{},
		&entities.BusSchedule{},
		&entities.Coupons{},
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
