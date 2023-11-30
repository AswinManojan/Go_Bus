package entities

type Stations struct {
	StationID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	StationName string `json:"name" gorm:"unique"`
	StationCode string `json:"code" gorm:"not null"`
	Location    string `json:"location" gorm:"not null"`
}
