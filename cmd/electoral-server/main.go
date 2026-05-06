package main

import (
	"log"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database/gormstore"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/mock"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/populator"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Store creation
	store := gormstore.Store{DB: database.GetDB()}
	electionStore := &gormstore.GormElectionStore{Store: store}
	candidateStore := &gormstore.GormCandidateStore{Store: store}
	// voterStore := &gormstore.GormVoterStore{Store: store}
	// votingBoothStore := &gormstore.GormVotingBoothStore{Store: store}
	// votingPlaceStore := &gormstore.GormVotingPlaceStore{Store: store}
	// electionDeploymentStore := &gormstore.GormElectionDeploymentStore{Store: store}

	// Service creation
	authService := &services.AuthService{}
	electionService := &services.ElectionService{ElectionStore: electionStore, AuthService: authService}
	candidateService := &services.CandidateService{CandidateStore: candidateStore}

	// Populator creation
	populator := populator.NewElectionPopulator(populator.Options{
		Populator:        &mock.MockPopulatorProvider{},
		ElectionService:  electionService,
		CandidateService: candidateService,
		RefreshInterval:  10 * time.Second,
	})

	populator.Start()
	defer populator.Stop()

	select {}
}
