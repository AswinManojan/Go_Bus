package db

import (
	"fmt"
	"gobus/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB is used to configure the DB connections and Automigrate the table to DB
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
		&entities.Booking{},
		&entities.BusSeatLayout{},
		&entities.PassengerInfo{},
		&entities.Schedule{},
		&entities.BusType{},
		&entities.Stations{},
		&entities.BaseFare{},
		&entities.RazorPay{},
		&entities.SubStation{},
	)
	return db
}

//DropTables function can be used to when we want to drop the entire tables
func DropTables() {
	dsn := "host=localhost user=postgres password=12345 dbname=gobus port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to DB")
	}
	if err := db.Migrator().DropTable(&entities.User{},
		&entities.ServiceProvider{},
		&entities.Buses{},
		&entities.BusSchedule{},
		&entities.Coupons{},
		&entities.Booking{},
		&entities.BusSeatLayout{},
		&entities.PassengerInfo{},
		&entities.Schedule{},
		&entities.BusType{},
		&entities.Stations{},
		&entities.BaseFare{},
		&entities.RazorPay{},
		&entities.SubStation{}); err != nil {
		fmt.Println("Error dropping the table:", err)
		return
	}
	fmt.Println("Table dropped successfully.")
}
