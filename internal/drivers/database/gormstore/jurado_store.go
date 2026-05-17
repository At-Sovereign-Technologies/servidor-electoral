package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormJuradoStore struct {
	Store
}

func (s *GormJuradoStore) Create(j *models.Jurado) error {
	return translateError(s.DB.Create(j).Error)
}

func (s *GormJuradoStore) GetByID(id uint) (*models.Jurado, error) {
	var j models.Jurado
	err := s.DB.First(&j, id).Error
	return &j, translateError(err)
}

func (s *GormJuradoStore) List() ([]models.Jurado, error) {
	var jurados []models.Jurado
	err := s.DB.Find(&jurados).Error
	return jurados, translateError(err)
}

func (s *GormJuradoStore) Update(j *models.Jurado) error {
	return translateError(s.DB.Save(j).Error)
}

func (s *GormJuradoStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Jurado{}, id))
}
