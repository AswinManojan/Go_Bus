package entities

type Buses struct {
	BusId     uint   `json:"bus_id" gorm:"primaryKey; autoIncrement"`
	BusNumber string `json:"bus_number" gorm:"unique"`
	// TotalSleeperSeats  uint   `json:"sl_seats"`
	// TotalPushBackSeats uint   `json:"pb_seats"`
	BusTypeCode string `json:"bus_type" gorm:"not null"`
	// SeatId       uint   `json:"seat_id" gorm:"not null"`
	ProviderId uint `json:"provider_id" gorm:"not null"`
	ScheduleId uint `json:"schedule_id" gorm:"not null"`
	// BusStationId uint   `json:"station_id" gorm:"not null"`
}

