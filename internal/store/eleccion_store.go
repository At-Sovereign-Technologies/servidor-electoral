package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type EleccionStore interface {
	Create(e *models.Eleccion) error
	GetByID(id uint) (*models.Eleccion, error)
	List() ([]models.Eleccion, error)
	Update(e *models.Eleccion) error
	Delete(id uint) error
}
