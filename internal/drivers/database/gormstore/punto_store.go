package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormPuntoStore struct {
	Store
}

func (s *GormPuntoStore) Create(p *models.Punto) error {
	return translateError(s.DB.Create(p).Error)
}

func (s *GormPuntoStore) GetByID(id uint) (*models.Punto, error) {
	var p models.Punto
	err := s.DB.Preload("Jurados").Preload("Terminales").First(&p, id).Error
	return &p, translateError(err)
}

func (s *GormPuntoStore) GetByEleccionID(eleccionID uint) ([]models.Punto, error) {
	var puntos []models.Punto
	err := s.DB.Where("eleccion_id = ?", eleccionID).Preload("Jurados").Preload("Terminales").Find(&puntos).Error
	return puntos, translateError(err)
}

func (s *GormPuntoStore) List() ([]models.Punto, error) {
	var puntos []models.Punto
	err := s.DB.Preload("Jurados").Preload("Terminales").Find(&puntos).Error
	return puntos, translateError(err)
}

func (s *GormPuntoStore) Update(p *models.Punto) error {
	return translateError(s.DB.Save(p).Error)
}

func (s *GormPuntoStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.Punto{}, id))
}
