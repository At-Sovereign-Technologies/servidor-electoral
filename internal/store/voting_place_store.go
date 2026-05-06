package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type VotingPlaceStore interface {
	Create(vp *models.VotingPlace) error
	GetByID(id uint) (*models.VotingPlace, error)
	ListByElection(electionID uint) ([]models.VotingPlace, error)
	Update(vp *models.VotingPlace) error
	Delete(id uint) error
}
