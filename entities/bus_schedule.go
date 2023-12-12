package entities

import (
	"time"

	"gorm.io/gorm"
)

// BusSchedule struct is used to store the information related to the chart based on the busID and day.
type BusSchedule struct {
	gorm.Model
	BusID             uint      `json:"bus_id" gorm:"not_null" validate:"required"`
	Day               time.Time `json:"day" gorm:"not_null" validate:"required"`
	DeckOneSeatLayout []byte
	DeckTwoSeatLayout []byte
	Status            string `json:"status" gorm:"default: Active" validate:"required"`
}
