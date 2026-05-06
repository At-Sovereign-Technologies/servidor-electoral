package services

import (
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type VotingPlaceService struct {
	VotingPlaceStore store.VotingPlaceStore
	AuthService      *AuthService
}

func (s *VotingPlaceService) CreateVotingPlace(vp *models.VotingPlace) (string, error) {
	secret, err := s.AuthService.GenerateSecret()
	if err != nil {
		return "", fmt.Errorf("failed to generate voting place secret: %w", err)
	}

	vp.Secret = secret
	vp.Disabled = false

	if err := vp.Validate(); err != nil {
		return "", fmt.Errorf("invalid voting place data: %w", err)
	}

	if err := s.VotingPlaceStore.Create(vp); err != nil {
		return "", err
	}

	return secret, nil
}

func (s *VotingPlaceService) DisableVotingPlace(id uint) error {
	vp, err := s.VotingPlaceStore.GetByID(id)
	if err != nil {
		return err
	}

	if vp.Disabled {
		return nil
	}

	vp.Disabled = true
	return s.VotingPlaceStore.Update(vp)
}

func (s *VotingPlaceService) EnableVotingPlace(id uint) error {
	vp, err := s.VotingPlaceStore.GetByID(id)
	if err != nil {
		return err
	}

	if !vp.Disabled {
		return nil
	}

	vp.Disabled = false
	return s.VotingPlaceStore.Update(vp)
}

func (s *VotingPlaceService) RegenerateVotingPlaceSecret(id uint) (string, error) {
	vp, err := s.VotingPlaceStore.GetByID(id)
	if err != nil {
		return "", err
	}

	secret, err := s.AuthService.GenerateSecret()
	if err != nil {
		return "", fmt.Errorf("failed to regenerate voting place secret: %w", err)
	}

	vp.Secret = secret
	if err := s.VotingPlaceStore.Update(vp); err != nil {
		return "", err
	}

	return secret, nil
}

func (s *VotingPlaceService) GetByElection(electionID uint) ([]models.VotingPlace, error) {
	return s.VotingPlaceStore.ListByElection(electionID)
}

func (s *VotingPlaceService) GetByID(id uint) (*models.VotingPlace, error) {
	return s.VotingPlaceStore.GetByID(id)
}
