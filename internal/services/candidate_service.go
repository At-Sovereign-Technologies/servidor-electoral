package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type CandidateService struct {
	CandidateStore store.CandidateStore
}

func (s *CandidateService) UpsertCandidate(candidate *models.Candidate) error {
	existing, err := s.CandidateStore.GetByID(candidate.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.createCandidate(candidate)
		}

		return fmt.Errorf("failed to check existing candidate: %w", err)
	}

	existing.Name = candidate.Name
	existing.Document = candidate.Document
	existing.Party = candidate.Party
	existing.Location = candidate.Location
	existing.PictureURL = candidate.PictureURL
	existing.Status = candidate.Status

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid candidate data: %w", err)
	}

	return s.CandidateStore.Update(existing)
}

func (s *CandidateService) createCandidate(candidate *models.Candidate) error {
	if err := candidate.Validate(); err != nil {
		return fmt.Errorf("invalid candidate data: %w", err)
	}

	return s.CandidateStore.Create(candidate)
}
