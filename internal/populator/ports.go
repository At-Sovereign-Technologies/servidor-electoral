package populator

import (
	"errors"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

var (
	ErrElectionNotFound = errors.New("election not found")
)

type Provider interface {
	Elections() ([]*models.Election, error)
	CandidatesForElection(electionID uint) ([]*models.Candidate, error)
}
