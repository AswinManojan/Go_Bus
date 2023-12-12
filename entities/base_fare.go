package entities

import "gorm.io/gorm"

// BaseFare struct is used to make a table which stores the data related to base fare based on the schedule/bus route.
type BaseFare struct {
	gorm.Model
	ScheduleID uint `json:"schedule_id" gorm:"not null" validate:"required"`
	BaseFare   uint `json:"fare" gorm:"not null" validate:"required"`
}
