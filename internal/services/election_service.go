package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type ElectionService struct {
	ElectionStore store.ElectionStore
	AuthService   *AuthService
}

func (s *ElectionService) UpsertElection(election *models.Election) error {
	existing, err := s.ElectionStore.GetByID(election.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.createElection(election)
		}

		return fmt.Errorf("failed to check existing election: %w", err)
	}

	existing.Name = election.Name
	existing.Status = election.Status
	existing.StartDate = election.StartDate
	existing.EndDate = election.EndDate

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid election data: %w", err)
	}

	return s.ElectionStore.Update(existing)

}

func (s *ElectionService) createElection(election *models.Election) error {
	secret, err := s.AuthService.GenerateSecret()
	if err != nil {
		return fmt.Errorf("failed to generate secret when creating election: %w", err)
	}

	election.Secret = secret

	if err := election.Validate(); err != nil {
		return fmt.Errorf("invalid election data: %w", err)
	}

	return s.ElectionStore.Create(election)
}
