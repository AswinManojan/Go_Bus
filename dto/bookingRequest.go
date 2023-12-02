package dto

import "github.com/lib/pq"

type BookingRequest struct {
	UsedCouponId  uint          `json:"coupon_id"`
	BusId         uint          `json:"bus_id"`
	PassengerId   pq.Int64Array `json:"passenger_id"`
	SeatsReserved []string      `json:"seat_reserved"`
	BookingDate   string        `json:"booking_date"`
}
