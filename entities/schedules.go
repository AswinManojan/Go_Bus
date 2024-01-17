package entities

// Schedule struct is used to store the informations related to the schedule of buses.
type Schedule struct {
	// Buses            Buses    `gorm:"foreignKey:ScheduleID;references:ScheduleID"`
	BaseFare         BaseFare `gorm:"foreignKey:ScheduleID;references:ScheduleID"`
	ScheduleID       uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	DepartureStation string   `json:"depart" gorm:"not null" validate:"required"`
	ArrivalStation   string   `json:"arrive" gorm:"not null" validate:"required"`
	DepartureTime    string
	ArrivalTime      string
}

// BusScheduleCombo struct is combination of Schedule and buses struct, used for providing response.
type BusScheduleCombo struct {
	Schedule
	Buses
}

//bus_schedule - BusScheduleCombo
