package models

import (
	"time"

	"gorm.io/gorm"
)

type TipoEleccion string

const (
	EleccionPresidencial TipoEleccion = "presidencial"
	EleccionLegislativa  TipoEleccion = "legislativa"
	EleccionTerritorial  TipoEleccion = "territorial"
)

type Eleccion struct {
	ID           uint         `gorm:"primaryKey" json:"id"`
	Nombre       string       `gorm:"size:255;not null" json:"nombre" validate:"required,min=3,max=255"`
	TipoEleccion TipoEleccion `gorm:"type:varchar(32);not null" json:"tipo_eleccion" validate:"required,oneof=presidencial legislativa territorial"`
	FechaInicio  int64        `gorm:"not null" json:"fecha_inicio" validate:"required,gt=0"`
	FechaFin     int64        `gorm:"not null" json:"fecha_fin" validate:"required,gt=0,gtefield=FechaInicio"`

	Candidatos []Candidato `gorm:"constraint:OnDelete:CASCADE;" json:"candidatos"`
	Nodos      []Nodo      `gorm:"constraint:OnDelete:CASCADE;" json:"nodos"`
	Puntos     []Punto     `gorm:"constraint:OnDelete:CASCADE;" json:"puntos"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (e *Eleccion) Validate() error {
	return validate.Struct(e)
}
