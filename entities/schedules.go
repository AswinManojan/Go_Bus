package entities

type Schedule struct {
	ScheduleId       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	DepartureStation string `json:"depart" gorm:"not null"`
	ArrivalStation   string `json:"arrive" gorm:"not null"`
	DepartureTime    string
	ArrivalTime      string
}

type Bus_schedule struct {
	Schedule
	Buses
}
