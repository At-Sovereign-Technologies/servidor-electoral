package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type TerminalService struct {
	Store store.TerminalStore
}

func (s *TerminalService) Upsert(terminal *models.Terminal) error {
	if err := terminal.Validate(); err != nil {
		return fmt.Errorf("invalid terminal data: %w", err)
	}

	if terminal.ID == 0 {
		return s.Store.Create(terminal)
	}

	existing, err := s.Store.GetByID(terminal.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(terminal)
		}
		return fmt.Errorf("failed to check existing terminal: %w", err)
	}

	existing.Activo = terminal.Activo
	existing.Secreto = terminal.Secreto
	existing.ClavePublica = terminal.ClavePublica

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid terminal data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *TerminalService) GetByID(id uint) (*models.Terminal, error) {
	return s.Store.GetByID(id)
}

func (s *TerminalService) GetByPuntoID(puntoID uint) ([]models.Terminal, error) {
	return s.Store.GetByPuntoID(puntoID)
}

func (s *TerminalService) List() ([]models.Terminal, error) {
	return s.Store.List()
}

func (s *TerminalService) Delete(id uint) error {
	return s.Store.Delete(id)
}
