package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormVotingBoothStore struct {
	Store
}

func (s *GormVotingBoothStore) Create(vb *models.VotingBooth) error {
	return translateError(s.DB.Create(vb).Error)
}

func (s *GormVotingBoothStore) GetByID(id uint) (*models.VotingBooth, error) {
	var vb models.VotingBooth
	err := s.DB.Preload("VotingPlace").First(&vb, id).Error
	return &vb, translateError(err)
}

func (s *GormVotingBoothStore) ListByPlace(placeID uint) ([]models.VotingBooth, error) {
	var booths []models.VotingBooth
	err := s.DB.Where("voting_place_id = ?", placeID).Find(&booths).Error
	return booths, translateError(err)
}

func (s *GormVotingBoothStore) Update(vb *models.VotingBooth) error {
	return translateError(s.DB.Save(vb).Error)
}

func (s *GormVotingBoothStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.VotingBooth{}, id))
}
