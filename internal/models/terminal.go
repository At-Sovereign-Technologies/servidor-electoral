package models

import (
	"time"

	"gorm.io/gorm"
)

type Terminal struct {
	ID      uint `gorm:"primaryKey" json:"id"`
	PuntoID uint `gorm:"not null;index" json:"punto_id"`
	Activo  bool `gorm:"not null" json:"activo"`

	// Secreto bearer token de la terminal
	Secreto string `gorm:"type:text;not null" json:"secreto" validate:"required_if=Activo true"`

	// ClavePublica Ed25519
	ClavePublica string `gorm:"type:text;not null" json:"clave_publica" validate:"required_if=Activo true"`

	Votantes []Votante `gorm:"constraint:OnDelete:CASCADE;" json:"votantes"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Terminal) Validate() error {
	return validate.Struct(t)
}
