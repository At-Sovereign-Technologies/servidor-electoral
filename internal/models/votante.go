package models

import (
	"time"

	"gorm.io/gorm"
)

type Votante struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	TerminalID uint `gorm:"not null;index" json:"terminal_id"`

	Nombre    string `gorm:"size:255;not null" json:"nombre" validate:"required,min=3,max=255"`
	Documento string `gorm:"size:32;not null" json:"documento" validate:"required,min=5,max=32"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (v *Votante) Validate() error {
	return validate.Struct(v)
}
