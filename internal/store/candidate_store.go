package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type CandidateStore interface {
	Create(c *models.Candidate) error
	GetByID(id uint) (*models.Candidate, error)
	ListByElection(electionID uint) ([]models.Candidate, error)
	Update(c *models.Candidate) error
	Delete(id uint) error
}
