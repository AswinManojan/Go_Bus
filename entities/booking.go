package entities

import "github.com/lib/pq"

// Booking struct is used to make a booking table to store the booking related information.
type Booking struct {
	BookingID        uint `json:"booking_id" gorm:"primaryKey; autoIncrement"`
	UserID           uint
	UsedCouponID     uint
	ActualFare       float64
	FarePostDiscount float64
	BusID            uint
	BookingDate      string         `json:"booking_date"  validate:"required"`
	PassengerID      pq.Int64Array  `gorm:"type:integer[]"  validate:"required"`
	SeatReserved     pq.StringArray `json:"seat_reserved" gorm:"type:text[]"  validate:"required"`
	Status           string
}
