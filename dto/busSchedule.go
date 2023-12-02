package dto

type BusSchedule struct {
	BusID uint   `json:"bus_id" gorm:"not_null"`
	Day   string `json:"day" gorm:"not_null"`
}
