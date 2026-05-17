package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type TerminalStore interface {
	Create(t *models.Terminal) error
	GetByID(id uint) (*models.Terminal, error)
	GetByPuntoID(puntoID uint) ([]models.Terminal, error)
	List() ([]models.Terminal, error)
	Update(t *models.Terminal) error
	Delete(id uint) error
}
