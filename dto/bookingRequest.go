package dto

import "github.com/lib/pq"

// BookingRequest struct is used to fetch the information for booking a seat from the user
type BookingRequest struct {
	UsedCouponID         uint          `json:"coupon_id"`
	BusID                uint          `json:"bus_id" gorm:"not null" validate:"required"`
	PassengerID          pq.Int64Array `json:"passenger_id" gorm:"not null" validate:"required"`
	SeatsReserved        []string      `json:"seat_reserved" gorm:"not null" validate:"required"`
	BookingDate          string        `json:"booking_date" gorm:"not null" validate:"required"`
	PreferredPaymentType string        `json:"payment_type" gorm:"default: Wallet"`
}
