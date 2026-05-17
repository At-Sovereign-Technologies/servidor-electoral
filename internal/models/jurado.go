package models

import (
	"time"

	"gorm.io/gorm"
)

type Jurado struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	PuntoID uint `gorm:"not null;index" json:"punto_id"`

	Nombre    string `gorm:"size:255;not null" json:"nombre" validate:"required,min=3,max=255"`
	Documento string `gorm:"size:32;not null;unique" json:"documento" validate:"required,min=5,max=32"`

	Usuario string `gorm:"size:64;not null;unique" json:"usuario" validate:"required,min=3,max=64,alphanum"`
	Hash    string `gorm:"type:text;not null" json:"hash" validate:"required"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (j *Jurado) Validate() error {
	return validate.Struct(j)
}
