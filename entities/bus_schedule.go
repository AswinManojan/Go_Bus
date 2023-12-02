package entities

import (
	"time"

	"gorm.io/gorm"
)

type BusSchedule struct {
	gorm.Model
	BusID             uint      `json:"bus_id" gorm:"not_null"`
	Day               time.Time `json:"day" gorm:"not_null"`
	DeckOneSeatLayout []byte
	DeckTwoSeatLayout []byte
}
