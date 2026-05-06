package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type VoterStore interface {
	Create(v *models.Voter) error
	GetByID(id uint) (*models.Voter, error)
	ListByElection(electionID uint) ([]models.Voter, error)
	Update(v *models.Voter) error
	Delete(id uint) error
}
