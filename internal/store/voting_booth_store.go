package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type VotingBoothStore interface {
	Create(vb *models.VotingBooth) error
	GetByID(id uint) (*models.VotingBooth, error)
	ListByPlace(placeID uint) ([]models.VotingBooth, error)
	Update(vb *models.VotingBooth) error
	Delete(id uint) error
}
