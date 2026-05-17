package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type JuradoService struct {
	Store store.JuradoStore
}

func (s *JuradoService) Upsert(jurado *models.Jurado) error {
	if err := jurado.Validate(); err != nil {
		return fmt.Errorf("invalid jurado data: %w", err)
	}

	if jurado.ID == 0 {
		return s.Store.Create(jurado)
	}

	existing, err := s.Store.GetByID(jurado.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(jurado)
		}
		return fmt.Errorf("failed to check existing jurado: %w", err)
	}

	existing.Nombre = jurado.Nombre
	existing.Documento = jurado.Documento
	existing.Usuario = jurado.Usuario
	existing.Hash = jurado.Hash

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid jurado data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *JuradoService) GetByID(id uint) (*models.Jurado, error) {
	return s.Store.GetByID(id)
}

func (s *JuradoService) List() ([]models.Jurado, error) {
	return s.Store.List()
}

func (s *JuradoService) Delete(id uint) error {
	return s.Store.Delete(id)
}
