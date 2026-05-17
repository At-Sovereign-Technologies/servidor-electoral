package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type CandidatoStore interface {
	Create(c *models.Candidato) error
	GetByID(id uint) (*models.Candidato, error)
	List() ([]models.Candidato, error)
	Update(c *models.Candidato) error
	Delete(id uint) error
}
