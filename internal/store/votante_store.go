package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type VotanteStore interface {
	Create(v *models.Votante) error
	GetByID(id uint) (*models.Votante, error)
	GetByTerminalID(terminalID uint) ([]models.Votante, error)
	List() ([]models.Votante, error)
	Update(v *models.Votante) error
	Delete(id uint) error
}
