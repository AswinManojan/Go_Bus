package entities

type PassengerInfo struct {
	PassengerId uint   `json:"passenger_id" gorm:"primaryKey; autoIncrement"`
	Name        string `json:"passenger_name" gorm:"not null"`
	Age         uint   `json:"age" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	UserId      uint   `json:"user_id"`
}
