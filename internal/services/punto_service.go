package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type PuntoService struct {
	Store store.PuntoStore
}

func (s *PuntoService) Upsert(punto *models.Punto) error {
	if err := punto.Validate(); err != nil {
		return fmt.Errorf("invalid punto data: %w", err)
	}

	if punto.ID == 0 {
		return s.Store.Create(punto)
	}

	existing, err := s.Store.GetByID(punto.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(punto)
		}
		return fmt.Errorf("failed to check existing punto: %w", err)
	}

	existing.Nombre = punto.Nombre
	existing.Latitud = punto.Latitud
	existing.Longitud = punto.Longitud
	existing.Activo = punto.Activo
	existing.Secreto = punto.Secreto

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid punto data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *PuntoService) GetByID(id uint) (*models.Punto, error) {
	return s.Store.GetByID(id)
}

func (s *PuntoService) GetByEleccionID(eleccionID uint) ([]models.Punto, error) {
	return s.Store.GetByEleccionID(eleccionID)
}

func (s *PuntoService) List() ([]models.Punto, error) {
	return s.Store.List()
}

func (s *PuntoService) Delete(id uint) error {
	return s.Store.Delete(id)
}
