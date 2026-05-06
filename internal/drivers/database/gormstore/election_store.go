package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormElectionStore struct {
	Store
}

func (s *GormElectionStore) Create(e *models.Election) error {
	return translateError(s.DB.Create(e).Error)
}

func (s *GormElectionStore) GetByID(id uint) (*models.Election, error) {
	var e models.Election
	err := s.DB.
		Preload("Candidates").
		Preload("VotingPlaces.VotingBooths").
		Preload("Voters").
		First(&e, id).Error
	return &e, translateError(err)
}

func (s *GormElectionStore) List() ([]models.Election, error) {
	var elections []models.Election
	err := s.DB.Find(&elections).Error
	return elections, translateError(err)
}

func (s *GormElectionStore) Update(e *models.Election) error {
	return translateError(s.DB.Save(e).Error)
}

func (s *GormElectionStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Election{}, id))
}
