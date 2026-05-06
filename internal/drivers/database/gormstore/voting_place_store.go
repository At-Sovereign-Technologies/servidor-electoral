package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormVotingPlaceStore struct {
	Store
}

func (s *GormVotingPlaceStore) Create(vp *models.VotingPlace) error {
	return translateError(s.DB.Create(vp).Error)
}

func (s *GormVotingPlaceStore) GetByID(id uint) (*models.VotingPlace, error) {
	var vp models.VotingPlace
	err := s.DB.
		Preload("VotingBooths").
		Preload("Voters").
		First(&vp, id).Error
	return &vp, translateError(err)
}

func (s *GormVotingPlaceStore) ListByElection(electionID uint) ([]models.VotingPlace, error) {
	var places []models.VotingPlace
	err := s.DB.Where("election_id = ?", electionID).Find(&places).Error
	return places, translateError(err)
}

func (s *GormVotingPlaceStore) Update(vp *models.VotingPlace) error {
	return translateError(s.DB.Save(vp).Error)
}

func (s *GormVotingPlaceStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.VotingPlace{}, id))
}
