package modelos

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	Nombre      string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	Imagen      string `gorm:"not null"`
}
