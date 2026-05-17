package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormEleccionStore struct {
	Store
}

func (s *GormEleccionStore) Create(e *models.Eleccion) error {
	return translateError(s.DB.Create(e).Error)
}

func (s *GormEleccionStore) GetByID(id uint) (*models.Eleccion, error) {
	var e models.Eleccion
	err := s.DB.Preload("Candidatos").Preload("Nodos").Preload("Puntos").First(&e, id).Error
	return &e, translateError(err)
}

func (s *GormEleccionStore) List() ([]models.Eleccion, error) {
	var elecciones []models.Eleccion
	err := s.DB.Preload("Candidatos").Preload("Nodos").Preload("Puntos").Find(&elecciones).Error
	return elecciones, translateError(err)
}

func (s *GormEleccionStore) Update(e *models.Eleccion) error {
	return translateError(s.DB.Save(e).Error)
}

func (s *GormEleccionStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Eleccion{}, id))
}
