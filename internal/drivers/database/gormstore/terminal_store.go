package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormTerminalStore struct {
	Store
}

func (s *GormTerminalStore) Create(t *models.Terminal) error {
	return translateError(s.DB.Create(t).Error)
}

func (s *GormTerminalStore) GetByID(id uint) (*models.Terminal, error) {
	var t models.Terminal
	err := s.DB.Preload("Votantes").First(&t, id).Error
	return &t, translateError(err)
}

func (s *GormTerminalStore) List() ([]models.Terminal, error) {
	var terminales []models.Terminal
	err := s.DB.Preload("Votantes").Find(&terminales).Error
	return terminales, translateError(err)
}

func (s *GormTerminalStore) Update(t *models.Terminal) error {
	return translateError(s.DB.Save(t).Error)
}

func (s *GormTerminalStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Terminal{}, id))
}
