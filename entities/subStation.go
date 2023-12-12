package entities

import "gorm.io/gorm"

//SubStation struct is used to add sub stations which is connected with other main stations
type SubStation struct {
	gorm.Model
	ParentID       uint
	ParentLocation string `json:"parent" gorm:"not null" validate:"required"`
	SubStation     string `json:"sub" gorm:"not null" validate:"required"`
}
