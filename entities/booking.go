package entities

import "github.com/lib/pq"

type Booking struct {
	BookingId        uint `json:"booking_id" gorm:"primaryKey; autoIncrement"`
	UserId           uint
	UsedCouponId     uint
	ActualFare       float64
	FarePostDiscount float64
	BusId            uint
	PassengerId      pq.Int64Array  `gorm:"type:integer[]"`
	SeatReserved     pq.StringArray `json:"seat_reserved" gorm:"type:text[]"`
	Status           string         `gorm:"default: Payment Pending"`
}
