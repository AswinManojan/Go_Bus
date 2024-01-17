package entities

// Stations struct is used to store the information of Bus Stations
type Stations struct {
	// Schedule    Schedule `gorm:"foreignKey:ArrivalStation;references:StationName"`
	// Schedules    Schedule `gorm:"foreignKey:DepartureStation;references:StationName"`
	StationID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	StationName string `json:"name" gorm:"unique" validate:"required"`
}
