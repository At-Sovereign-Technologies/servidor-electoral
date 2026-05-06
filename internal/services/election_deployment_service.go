package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type ElectionDeploymentService struct {
	ElectionDeploymentStore store.ElectionDeploymentStore
}

func (s *ElectionDeploymentService) GetByID(id uint) (*models.ElectionDeployment, error) {
	return s.ElectionDeploymentStore.GetByElection(id)
}

func (s *ElectionDeploymentService) UpsertDeployment(deployment *models.ElectionDeployment) error {
	existing, err := s.ElectionDeploymentStore.GetByElection(deployment.ElectionID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.ElectionDeploymentStore.Create(deployment)
		}

		return fmt.Errorf("failed to check existing election deployment: %w", err)
	}

	existing.QueryURL = deployment.QueryURL
	existing.DatabaseURI = deployment.DatabaseURI
	existing.QueueURI = deployment.QueueURI

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid election deployment data: %w", err)
	}

	return s.ElectionDeploymentStore.Update(existing)
}
