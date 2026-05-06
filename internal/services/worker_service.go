package services

import (
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/utils"
)

type WorkerConfig struct {
	Worker   *models.ElectionWorker
	CertPEM  string
	KeyPEM   string
	Endpoint string
}

type WorkerService struct {
	AuthService *AuthService
	WorkerStore store.ElectionWorkerStore
}

func (s *WorkerService) GetByID(id uint) (*models.ElectionWorker, error) {
	return s.WorkerStore.GetByID(id)
}

func (s *WorkerService) GetByElection(electionID uint) ([]models.ElectionWorker, error) {
	return s.WorkerStore.ListByElection(electionID)
}

func (s *WorkerService) CreateWorker(name string, electionID uint) (*WorkerConfig, error) {
	worker := &models.ElectionWorker{
		Name:       name,
		ElectionID: electionID,
	}

	cert, key, err := s.AuthService.IssueCertificate(worker.Name)
	if err != nil {
		return nil, err
	}

	if err := s.WorkerStore.Create(worker); err != nil {
		return nil, err
	}

	config := WorkerConfig{
		Worker:   worker,
		CertPEM:  string(cert),
		KeyPEM:   string(key),
		Endpoint: utils.GetEnv("SELF_ENDPOINT", "localhost:8080"),
	}

	return &config, nil
}

func (s *WorkerService) IsWorkerActive(id uint) (bool, error) {
	worker, err := s.WorkerStore.GetByID(id)
	if err != nil {
		return false, err
	}

	return worker != nil && !worker.Revoked, nil
}

func (s *WorkerService) setWorkerRevoked(id uint, revoked bool) error {
	worker, err := s.WorkerStore.GetByID(id)
	if err != nil {
		return err
	}
	if worker == nil {
		return fmt.Errorf("worker with id %d not found", id)
	}

	worker.Revoked = revoked
	return s.WorkerStore.Update(worker)
}

func (s *WorkerService) DisableWorker(id uint) error {
	return s.setWorkerRevoked(id, true)
}

func (s *WorkerService) EnableWorker(id uint) error {
	return s.setWorkerRevoked(id, false)
}

func (s *WorkerService) UpdateWorker(worker *models.ElectionWorker) (*models.ElectionWorker, error) {
	if err := s.WorkerStore.Update(worker); err != nil {
		return nil, err
	}

	return worker, nil
}

func (s *WorkerService) DeleteWorker(id uint) error {
	return s.WorkerStore.Delete(id)
}
