package models

import (
	"time"

	"gorm.io/gorm"
)

type Candidato struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	EleccionID uint `gorm:"not null;index" json:"eleccion_id"`

	Nombre    string `gorm:"size:255;not null" json:"nombre" validate:"required,min=3,max=255"`
	Documento string `gorm:"size:32;not null;unique" json:"documento" validate:"required,min=5,max=32"`
	Partido   string `gorm:"size:255;not null" json:"partido" validate:"required,min=2,max=255"`
	FotoURL   string `gorm:"type:text" json:"foto_url" validate:"required"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (e *Candidato) Validate() error {
	return validate.Struct(e)
}
