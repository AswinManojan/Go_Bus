package dto

// BusRequest struct is used to store the information shared by the user to fetch the available buses for that route
type BusRequest struct {
	DepartureStation string `json:"depart" gorm:"not null" validate:"required"`
	ArrivalStation   string `json:"arrival" gorm:"not null" validate:"required"`
	Price            int
	Duration         int `json:"duration"`
}
