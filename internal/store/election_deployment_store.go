package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type ElectionDeploymentStore interface {
	Create(ed *models.ElectionDeployment) error
	GetByElection(electionID uint) ([]models.ElectionDeployment, error)
	Update(ed *models.ElectionDeployment) error
	Delete(id uint) error
}
