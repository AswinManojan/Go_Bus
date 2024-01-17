package entities

// Buses struct is used to store the information related to the buses.
type Buses struct {
	// BusSchedule        BusSchedule `gorm:"foreignKey:BusID;references:BusID"`
	BusID              uint        `json:"bus_id" gorm:"primaryKey; autoIncrement"`
	BusNumber          string      `json:"bus_number" gorm:"unique" validate:"required"`
	TotalSleeperSeats  uint
	TotalPushBackSeats uint
	BusTypeCode        string `json:"bus_type" gorm:"not null" validate:"required"`
	// SeatId       uint   `json:"seat_id" gorm:"not null"`
	ProviderID uint `json:"provider_id" gorm:"not null" validate:"required"`
	ScheduleID uint `json:"schedule_id" gorm:"not null" validate:"required"`
	// BusStationId uint   `json:"station_id" gorm:"not null"`
}

type BusesResp struct {
	// BusSchedule        BusSchedule `gorm:"foreignKey:BusID;references:BusID"`
	BusID              uint   `json:"bus_id" gorm:"primaryKey; autoIncrement"`
	BusNumber          string `json:"bus_number" gorm:"unique" validate:"required"`
	TotalSleeperSeats  uint
	TotalPushBackSeats uint
	BusTypeCode        string `json:"bus_type" gorm:"not null" validate:"required"`
	// SeatId       uint   `json:"seat_id" gorm:"not null"`
	// ProviderID uint `json:"provider_id" gorm:"not null" validate:"required"`
	// ScheduleID uint `json:"schedule_id" gorm:"not null" validate:"required"`
	// BusStationId uint   `json:"station_id" gorm:"not null"`
}
