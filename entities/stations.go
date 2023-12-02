package entities

type Stations struct {
	StationID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	StationName string `json:"name" gorm:"unique"`
}