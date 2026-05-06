package main

import (
	"log"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database/gormstore"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/mock"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web"
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
	votingPlaceStore := &gormstore.GormVotingPlaceStore{Store: store}
	workerStore := &gormstore.GormWorkerStore{Store: store}
	// voterStore := &gormstore.GormVoterStore{Store: store}
	// votingBoothStore := &gormstore.GormVotingBoothStore{Store: store}
	// electionDeploymentStore := &gormstore.GormElectionDeploymentStore{Store: store}

	// Service creation
	authService, err := services.NewAuthService("config")
	if err != nil {
		log.Fatalf("Failed to create auth service: %v", err)
	}

	electionService := &services.ElectionService{ElectionStore: electionStore, AuthService: authService}
	candidateService := &services.CandidateService{CandidateStore: candidateStore}
	votingPlaceService := &services.VotingPlaceService{VotingPlaceStore: votingPlaceStore, AuthService: authService}
	workerService := &services.WorkerService{AuthService: authService, WorkerStore: workerStore}

	// Populator creation
	populator := populator.NewElectionPopulator(populator.Options{
		Populator:        &mock.MockPopulatorProvider{},
		ElectionService:  electionService,
		CandidateService: candidateService,
		RefreshInterval:  10 * time.Second,
	})

	populator.Start()
	defer populator.Stop()

	// Router creation
	router, err := web.NewRouter(electionService, candidateService, votingPlaceService, workerService)
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}

	// Start HTTP server
	log.Println("Starting server on :8080")
	if err := router.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
