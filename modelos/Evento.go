package modelos

import "gorm.io/gorm"

type Evento struct {
	gorm.Model
	Nombre       string `gorm:"not null"`
	Descripcion  string `gorm:"not null"`
	FechaTermino string `gorm:"not null"`
}
