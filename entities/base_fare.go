package entities

import "gorm.io/gorm"

type BaseFare struct {
	gorm.Model
	ScheduleId uint `json:"schedule_id" gorm:"not null"`
	BaseFare   uint `json:"fare" gorm:"not null"`
}
