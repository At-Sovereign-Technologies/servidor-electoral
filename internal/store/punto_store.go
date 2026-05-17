package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type PuntoStore interface {
	Create(p *models.Punto) error
	GetByID(id uint) (*models.Punto, error)
	GetByEleccionID(eleccionID uint) ([]models.Punto, error)
	List() ([]models.Punto, error)
	Update(p *models.Punto) error
	Delete(id uint) error
}
