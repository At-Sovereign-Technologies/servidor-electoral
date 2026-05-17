package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormNodoStore struct {
	Store
}

func (s *GormNodoStore) Create(n *models.Nodo) error {
	return translateError(s.DB.Create(n).Error)
}

func (s *GormNodoStore) GetByID(id uint) (*models.Nodo, error) {
	var n models.Nodo
	err := s.DB.First(&n, id).Error
	return &n, translateError(err)
}

func (s *GormNodoStore) GetByEleccionID(eleccionID uint) ([]models.Nodo, error) {
	var nodos []models.Nodo
	err := s.DB.Where("eleccion_id = ?", eleccionID).Find(&nodos).Error
	return nodos, translateError(err)
}

func (s *GormNodoStore) List() ([]models.Nodo, error) {
	var nodos []models.Nodo
	err := s.DB.Find(&nodos).Error
	return nodos, translateError(err)
}

func (s *GormNodoStore) Update(n *models.Nodo) error {
	return translateError(s.DB.Save(n).Error)
}

func (s *GormNodoStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Nodo{}, id))
}
