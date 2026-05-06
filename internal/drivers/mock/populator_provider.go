package mock

import (
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/populator"
)

var _ populator.Provider = (*MockPopulatorProvider)(nil)

type MockPopulatorProvider struct{}

func (m *MockPopulatorProvider) Elections() ([]*models.Election, error) {
	return []*models.Election{
		{
			ID:        1,
			Name:      "Mock Election 1",
			Status:    models.ElectionStatusOngoing,
			StartDate: time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 6, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			Name:      "Mock Election 2",
			Status:    models.ElectionStatusOngoing,
			StartDate: time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 7, 14, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

func (m *MockPopulatorProvider) CandidatesForElection(electionID uint) ([]*models.Candidate, error) {
	return []*models.Candidate{
		{
			ID:         ((electionID - 1) * 2) + 1,
			Name:       "Mock Candidate 1",
			Document:   "123456",
			Party:      "Some Party",
			Location:   "Bogota",
			PictureURL: "https://i.scdn.co/image/ab67616d00001e0291a5c5a6cc2d8e48ebe3ae8d",
			Status:     models.CandidateStatusApproved,
			ElectionID: electionID,
		},
		{
			ID:         ((electionID - 1) * 2) + 2,
			Name:       "Mock Candidate 2",
			Document:   "654321",
			Party:      "Another Party",
			Location:   "Medellin",
			PictureURL: "https://i.scdn.co/image/ab67616d00001e0291a5c5a6cc2d8e48ebe3ae8d",
			Status:     models.CandidateStatusApproved,
			ElectionID: electionID,
		},
	}, nil
}
