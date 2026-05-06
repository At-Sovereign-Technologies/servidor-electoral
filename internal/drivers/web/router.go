package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web/handlers"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type Router struct {
	e                  *echo.Echo
	electionHandler    *handlers.ElectionHandler
	votingPlaceHandler *handlers.VotingPlaceHandler
	workersHandler     *handlers.WorkerHandler
}

func NewRouter(
	electionService *services.ElectionService,
	candidateService *services.CandidateService,
	votingPlaceService *services.VotingPlaceService,
	workerService *services.WorkerService,
) (*Router, error) {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = HTMLHTTPErrorHandler

	renderer, err := NewTemplateRenderer()
	if err != nil {
		return nil, err
	}

	e.Renderer = renderer

	electionHandler, err := handlers.NewElectionHandler(electionService, candidateService)
	if err != nil {
		return nil, err
	}

	votingPlaceHandler, err := handlers.NewVotingPlaceHandler(votingPlaceService)
	if err != nil {
		return nil, err
	}

	workerHandler, err := handlers.NewWorkerHandler(workerService)
	if err != nil {
		return nil, err
	}

	r := &Router{
		e:                  e,
		electionHandler:    electionHandler,
		votingPlaceHandler: votingPlaceHandler,
		workersHandler:     workerHandler,
	}

	r.registerRoutes()
	return r, nil
}

func (r *Router) registerRoutes() {
	// Default redirection
	r.e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/elections")
	})

	// Elections
	r.e.GET("/elections", r.electionHandler.ListElections)
	r.e.GET("/elections/:id", r.electionHandler.GetElection)

	// Candidates
	r.e.GET("/elections/:id/candidates", r.electionHandler.ListCandidates)

	// Voting places
	r.e.GET("/elections/:id/voting-places", r.votingPlaceHandler.List)
	r.e.POST("/elections/:id/voting-places", r.votingPlaceHandler.Post)

	// Workers
	r.e.GET("/elections/:id/workers", r.workersHandler.List)
	r.e.POST("/elections/:id/workers", r.workersHandler.Post)

	// Static files
	r.e.Static("/static", "static")

	r.e.RouteNotFound("/*", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusNotFound, "Página no encontrada")
	})
}

func (r *Router) Start(addr string) error {
	return r.e.Start(addr)
}
