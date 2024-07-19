package modelos

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Titulo string `gorm:"not null"`
	Imagen string `gorm:"not null"`
}
