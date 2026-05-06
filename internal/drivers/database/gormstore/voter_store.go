package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormVoterStore struct {
	Store
}

func (s *GormVoterStore) Create(v *models.Voter) error {
	return translateError(s.DB.Create(v).Error)
}

func (s *GormVoterStore) GetByID(id uint) (*models.Voter, error) {
	var v models.Voter
	err := s.DB.
		Preload("Election").
		Preload("VotingPlace").
		First(&v, id).Error
	return &v, translateError(err)
}

func (s *GormVoterStore) ListByElection(electionID uint) ([]models.Voter, error) {
	var voters []models.Voter
	err := s.DB.Where("election_id = ?", electionID).Find(&voters).Error
	return voters, translateError(err)
}

func (s *GormVoterStore) Update(v *models.Voter) error {
	return translateError(s.DB.Save(v).Error)
}

func (s *GormVoterStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Voter{}, id))
}
