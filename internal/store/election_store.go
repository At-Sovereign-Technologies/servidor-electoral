package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type ElectionStore interface {
	Create(e *models.Election) error
	GetByID(id uint) (*models.Election, error)
	List() ([]models.Election, error)
	Update(e *models.Election) error
	Delete(id uint) error
}
