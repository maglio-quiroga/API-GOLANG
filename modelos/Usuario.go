package modelos

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string `gorm:"not null"`
	Apellido string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Clave    string `gorm:"not null"`
	Permisos bool   `gorm:"not null"`
}
