package gormstore

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type GormElectionDeploymentStore struct {
	Store
}

func (s *GormElectionDeploymentStore) Create(ed *models.ElectionDeployment) error {
	return translateError(s.DB.Create(ed).Error)
}

func (s *GormElectionDeploymentStore) GetByElection(electionID uint) (*models.ElectionDeployment, error) {
	var deployment models.ElectionDeployment
	err := s.DB.Where("election_id = ?", electionID).First(&deployment).Error
	return &deployment, translateError(err)
}

func (s *GormElectionDeploymentStore) Update(ed *models.ElectionDeployment) error {
	return translateResult(s.DB.Save(ed))
}

func (s *GormElectionDeploymentStore) Delete(id uint) error {
	return translateResult(s.DB.Delete(&models.ElectionDeployment{}, id))
}
