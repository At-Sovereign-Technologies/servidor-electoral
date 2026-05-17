package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormVotanteStore struct {
	Store
}

func (s *GormVotanteStore) Create(v *models.Votante) error {
	return translateError(s.DB.Create(v).Error)
}

func (s *GormVotanteStore) GetByID(id uint) (*models.Votante, error) {
	var v models.Votante
	err := s.DB.First(&v, id).Error
	return &v, translateError(err)
}

func (s *GormVotanteStore) GetByTerminalID(terminalID uint) ([]models.Votante, error) {
	var votantes []models.Votante
	err := s.DB.Where("terminal_id = ?", terminalID).Find(&votantes).Error
	return votantes, translateError(err)
}

func (s *GormVotanteStore) List() ([]models.Votante, error) {
	var votantes []models.Votante
	err := s.DB.Find(&votantes).Error
	return votantes, translateError(err)
}

func (s *GormVotanteStore) Update(v *models.Votante) error {
	return translateError(s.DB.Save(v).Error)
}

func (s *GormVotanteStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Votante{}, id))
}
