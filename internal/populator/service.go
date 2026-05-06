package populator

import (
	"fmt"
	"log"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type Options struct {
	Populator        Provider
	ElectionService  *services.ElectionService
	CandidateService *services.CandidateService
	RefreshInterval  time.Duration
}

type ElectionPopulator struct {
	provider        Provider
	refreshInterval time.Duration

	electionService  *services.ElectionService
	candidateService *services.CandidateService

	timer *time.Timer
}

func NewElectionPopulator(options Options) *ElectionPopulator {
	return &ElectionPopulator{
		provider:         options.Populator,
		electionService:  options.ElectionService,
		candidateService: options.CandidateService,
		refreshInterval:  options.RefreshInterval,
	}
}

func (p *ElectionPopulator) Refresh() error {
	elections, err := p.provider.Elections()
	if err != nil {
		return fmt.Errorf("failed to fetch elections: %w", err)
	}

	for _, election := range elections {
		if err := p.electionService.UpsertElection(election); err != nil {
			log.Printf("failed to upsert election with ID: %d: %v", election.ID, err)
		}

		candidates, err := p.provider.CandidatesForElection(election.ID)
		if err != nil {
			log.Printf("failed to fetch candidates for election %d: %v", election.ID, err)
			continue
		}

		for _, candidate := range candidates {
			if err := p.candidateService.UpsertCandidate(candidate); err != nil {
				log.Printf("failed to upsert candidate with ID: %d: %v", candidate.ID, err)
			}
		}

	}

	return nil
}

func (p *ElectionPopulator) Start() error {
	if p.timer != nil {
		p.timer.Stop()
	}

	p.timer = time.AfterFunc(p.refreshInterval, func() {
		if err := p.Refresh(); err != nil {
			log.Printf("Error refreshing elections: %v", err)
		}

		p.Start()
	})

	return nil
}

func (p *ElectionPopulator) Stop() {
	if p.timer != nil {
		p.timer.Stop()
	}
}
