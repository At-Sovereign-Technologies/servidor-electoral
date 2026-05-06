package gormstore

import "github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"

type GormWorkerStore struct {
	Store
}

func (s *GormWorkerStore) Create(worker *models.ElectionWorker) error {
	return translateError(s.DB.Create(worker).Error)
}

func (s *GormWorkerStore) GetByID(id uint) (*models.ElectionWorker, error) {
	var worker models.ElectionWorker
	err := s.DB.First(&worker, id).Error
	return &worker, translateError(err)
}

func (s *GormWorkerStore) ListByElection(electionID uint) ([]models.ElectionWorker, error) {
	var workers []models.ElectionWorker
	err := s.DB.Where("election_id = ?", electionID).Find(&workers).Error
	return workers, translateError(err)
}

func (s *GormWorkerStore) Update(worker *models.ElectionWorker) error {
	return translateError(s.DB.Save(worker).Error)
}

func (s *GormWorkerStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.ElectionWorker{}, id))
}
