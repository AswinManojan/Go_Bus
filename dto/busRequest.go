package dto

type BusRequest struct {
	DepartureStation string `json:"depart" gorm:"not null"`
	ArrivalStation   string `json:"arrival" gorm:"not null"`
}
