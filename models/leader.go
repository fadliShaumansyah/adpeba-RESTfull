package models

import (
	"gorm.io/gorm"
)

type Leader struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`                // Nama wajib
	Position string `json:"position" binding:"required"`            // Posisi wajib
	Jamaah   string `json:"jamaah" binding:"required"`           // Jamaah wajib
	Periode  string `json:"periode" binding:"required"`             // Periode wajib
}
