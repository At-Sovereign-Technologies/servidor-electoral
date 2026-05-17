package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type VotanteService struct {
	Store store.VotanteStore
}

func (s *VotanteService) Upsert(votante *models.Votante) error {
	if err := votante.Validate(); err != nil {
		return fmt.Errorf("invalid votante data: %w", err)
	}

	if votante.ID == 0 {
		return s.Store.Create(votante)
	}

	existing, err := s.Store.GetByID(votante.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(votante)
		}
		return fmt.Errorf("failed to check existing votante: %w", err)
	}

	existing.Nombre = votante.Nombre
	existing.Documento = votante.Documento

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid votante data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *VotanteService) GetByID(id uint) (*models.Votante, error) {
	return s.Store.GetByID(id)
}

func (s *VotanteService) GetByTerminalID(terminalID uint) ([]models.Votante, error) {
	return s.Store.GetByTerminalID(terminalID)
}

func (s *VotanteService) List() ([]models.Votante, error) {
	return s.Store.List()
}

func (s *VotanteService) Delete(id uint) error {
	return s.Store.Delete(id)
}
