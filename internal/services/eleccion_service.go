package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type EleccionService struct {
	Store store.EleccionStore
}

func (s *EleccionService) Upsert(eleccion *models.Eleccion) error {
	if err := eleccion.Validate(); err != nil {
		return fmt.Errorf("invalid eleccion data: %w", err)
	}

	if eleccion.ID == 0 {
		return s.Store.Create(eleccion)
	}

	existing, err := s.Store.GetByID(eleccion.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(eleccion)
		}
		return fmt.Errorf("failed to check existing eleccion: %w", err)
	}

	existing.Nombre = eleccion.Nombre
	existing.TipoEleccion = eleccion.TipoEleccion
	existing.FechaInicio = eleccion.FechaInicio
	existing.FechaFin = eleccion.FechaFin

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid eleccion data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *EleccionService) GetByID(id uint) (*models.Eleccion, error) {
	return s.Store.GetByID(id)
}

func (s *EleccionService) List() ([]models.Eleccion, error) {
	return s.Store.List()
}

func (s *EleccionService) Delete(id uint) error {
	return s.Store.Delete(id)
}
