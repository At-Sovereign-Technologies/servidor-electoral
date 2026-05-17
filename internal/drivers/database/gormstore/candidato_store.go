package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormCandidatoStore struct {
	Store
}

func (s *GormCandidatoStore) Create(c *models.Candidato) error {
	return translateError(s.DB.Create(c).Error)
}

func (s *GormCandidatoStore) GetByID(id uint) (*models.Candidato, error) {
	var c models.Candidato
	err := s.DB.First(&c, id).Error
	return &c, translateError(err)
}

func (s *GormCandidatoStore) List() ([]models.Candidato, error) {
	var candidatos []models.Candidato
	err := s.DB.Find(&candidatos).Error
	return candidatos, translateError(err)
}

func (s *GormCandidatoStore) Update(c *models.Candidato) error {
	return translateError(s.DB.Save(c).Error)
}

func (s *GormCandidatoStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Candidato{}, id))
}
