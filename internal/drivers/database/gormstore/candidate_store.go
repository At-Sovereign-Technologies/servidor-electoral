package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormCandidateStore struct {
	Store
}

func (s *GormCandidateStore) Create(c *models.Candidate) error {
	return translateError(s.DB.Create(c).Error)
}

func (s *GormCandidateStore) GetByID(id uint) (*models.Candidate, error) {
	var c models.Candidate
	err := s.DB.Preload("Election").First(&c, id).Error
	return &c, translateError(err)
}

func (s *GormCandidateStore) ListByElection(electionID uint) ([]models.Candidate, error) {
	var candidates []models.Candidate
	err := s.DB.Where("election_id = ?", electionID).Find(&candidates).Error
	return candidates, translateError(err)
}

func (s *GormCandidateStore) Update(c *models.Candidate) error {
	return translateError(s.DB.Save(c).Error)
}

func (s *GormCandidateStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Candidate{}, id))
}
