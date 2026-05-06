package store

import "github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"

type ElectionWorkerStore interface {
	Create(worker *models.ElectionWorker) error
	GetByID(id uint) (*models.ElectionWorker, error)
	ListByElection(electionID uint) ([]models.ElectionWorker, error)
	Update(worker *models.ElectionWorker) error
	Delete(id uint) error
}
