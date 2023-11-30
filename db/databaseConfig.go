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
	db.AutoMigrate(&entities.User{}, &entities.ServiceProvider{})
	return db
}
