package entities

// PassengerInfo struct is used to store the informations related to passengers.
type PassengerInfo struct {
	PassengerID uint   `json:"passenger_id" gorm:"primaryKey; autoIncrement"`
	Name        string `json:"passenger_name" gorm:"not null" validate:"required"`
	Age         uint   `json:"age" gorm:"not null" validate:"required"`
	Gender      string `json:"gender" gorm:"not null" validate:"required"`
	UserID      uint   `json:"user_id"`
}
