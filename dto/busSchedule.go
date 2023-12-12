package dto

// BusSchedule struct is get the input from the user inorder to fetch the chart based on the bus and the day provided.
type BusSchedule struct {
	BusID uint   `json:"bus_id" gorm:"not_null" validate:"required"`
	Day   string `json:"day" gorm:"not_null" validate:"required"`
}
