package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type CandidatoService struct {
	Store store.CandidatoStore
}

func (s *CandidatoService) Upsert(candidato *models.Candidato) error {
	if err := candidato.Validate(); err != nil {
		return fmt.Errorf("invalid candidato data: %w", err)
	}

	if candidato.ID == 0 {
		return s.Store.Create(candidato)
	}

	existing, err := s.Store.GetByID(candidato.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(candidato)
		}
		return fmt.Errorf("failed to check existing candidato: %w", err)
	}

	existing.Nombre = candidato.Nombre
	existing.Documento = candidato.Documento
	existing.Partido = candidato.Partido
	existing.FotoURL = candidato.FotoURL

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid candidato data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *CandidatoService) GetByID(id uint) (*models.Candidato, error) {
	return s.Store.GetByID(id)
}

func (s *CandidatoService) List() ([]models.Candidato, error) {
	return s.Store.List()
}

func (s *CandidatoService) Delete(id uint) error {
	return s.Store.Delete(id)
}
