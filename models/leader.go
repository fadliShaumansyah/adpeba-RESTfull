package models

import (
	"gorm.io/gorm"
)

type Leader struct {
	gorm.Model
	Name		string `json:"name" gorm:"not null"`
	Postition	string `json:"position" gorm:"not null"`
	Jamaah		string `json:"jamaah_id" gorm:"not null"`
	Periode		string `json:"periode" gorm:"not null"`
}