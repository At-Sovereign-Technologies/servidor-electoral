package web

import (
	"github.com/labstack/echo/v4"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web/handlers"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type Router struct {
	e                  *echo.Echo
	electionHandler    *handlers.ElectionHandler
	candidateHandler   *handlers.CandidateHandler
	votingPlaceHandler *handlers.VotingPlaceHandler
}

func NewRouter(
	electionService *services.ElectionService,
	candidateService *services.CandidateService,
	votingPlaceService *services.VotingPlaceService,
) (*Router, error) {
	e := echo.New()

	electionHandler, err := handlers.NewElectionHandler(electionService)
	if err != nil {
		return nil, err
	}

	candidateHandler, err := handlers.NewCandidateHandler(candidateService)
	if err != nil {
		return nil, err
	}

	votingPlaceHandler, err := handlers.NewVotingPlaceHandler(votingPlaceService)
	if err != nil {
		return nil, err
	}

	r := &Router{
		e:                  e,
		electionHandler:    electionHandler,
		candidateHandler:   candidateHandler,
		votingPlaceHandler: votingPlaceHandler,
	}

	r.registerRoutes()
	return r, nil
}

func (r *Router) registerRoutes() {
	// Elections
	r.e.GET("/elections", r.electionHandler.ListElections)
	r.e.GET("/elections/:id", r.electionHandler.GetElection)

	// Candidates
	r.e.GET("/candidates", r.candidateHandler.ListCandidates)

	// Voting places (method-based)
	r.e.GET("/voting-places", r.votingPlaceHandler.List)
	r.e.POST("/voting-places", r.votingPlaceHandler.Post)

	// Static files
	r.e.Static("/", "static") // or use embed if needed
}

func (r *Router) Start(addr string) error {
	return r.e.Start(addr)
}
