package handlers

import (
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
	"github.com/labstack/echo/v4"
)

type ElectionHandler struct {
	ElectionService  *services.ElectionService
	CandidateService *services.CandidateService
}

func NewElectionHandler(electionService *services.ElectionService, candidateService *services.CandidateService) (*ElectionHandler, error) {
	return &ElectionHandler{
		ElectionService:  electionService,
		CandidateService: candidateService,
	}, nil
}

// GET /elections
func (h *ElectionHandler) ListElections(c echo.Context) error {
	elections, err := h.ElectionService.ElectionStore.List()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch elections")
	}

	data := map[string]interface{}{
		"Title":     "Elecciones",
		"Elections": elections,
	}

	// HTMX support (optional partial rendering)
	if isHTMX(c) {
		return c.Render(http.StatusOK, "partials/elections_list.html", data)
	}

	return c.Render(http.StatusOK, "pages/elections.html", data)
}

// GET /elections/:id
func (h *ElectionHandler) GetElection(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	election, err := h.ElectionService.ElectionStore.GetByID(uint(id))
	if err != nil {
		return c.String(http.StatusNotFound, "Election not found")
	}

	data := map[string]interface{}{
		"Title":      election.Name,
		"Election":   election,
		"Candidates": election.Candidates,
		"Start":      election.StartDate.Format("02/01/2006 15:04"),
		"End":        election.EndDate.Format("02/01/2006 15:04"),
	}

	if isHTMX(c) {
		return c.Render(http.StatusOK, "partials/election_detail.html", data)
	}

	return c.Render(http.StatusOK, "pages/election_detail.html", data)
}

// GET /elections/:id/candidates
func (h *ElectionHandler) ListCandidates(c echo.Context) error {
	electionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	candidates, err := h.CandidateService.CandidateStore.ListByElection(uint(electionID))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to fetch candidates")
	}

	data := map[string]interface{}{
		"Title":      "Candidatos",
		"Candidates": candidates,
		"ElectionID": electionID,
	}

	// HTMX partial render: return a page-level partial for HTMX navigation
	if isHTMX(c) {
		return c.Render(http.StatusOK, "partials/candidates_page.html", data)
	}

	return c.Render(http.StatusOK, "pages/candidates.html", data)
}
