package models

import (
	"time"

	"gorm.io/gorm"
)

type Punto struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	EleccionID uint `gorm:"not null;index" json:"eleccion_id"`

	Nombre   string  `gorm:"size:255;not null" json:"nombre" validate:"required,min=3,max=255"`
	Latitud  float64 `gorm:"not null" json:"latitud" validate:"required,gte=-90,lte=90"`
	Longitud float64 `gorm:"not null" json:"longitud" validate:"required,gte=-180,lte=180"`
	Activo   bool    `gorm:"not null" json:"activo"`

	// Secreto del puesto
	Secreto string `gorm:"type:text;not null" json:"secreto" validate:"required_if=Activo true"`

	Jurados    []Jurado   `gorm:"constraint:OnDelete:CASCADE;" json:"jurados"`
	Terminales []Terminal `gorm:"constraint:OnDelete:CASCADE;" json:"terminales"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Punto) Validate() error {
	return validate.Struct(p)
}
