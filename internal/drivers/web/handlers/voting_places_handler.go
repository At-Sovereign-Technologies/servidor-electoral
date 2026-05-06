package handlers

import (
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
	"github.com/labstack/echo/v4"
)

type VotingPlaceHandler struct {
	VotingPlaceService *services.VotingPlaceService
}

func NewVotingPlaceHandler(votingPlaceService *services.VotingPlaceService) (*VotingPlaceHandler, error) {
	return &VotingPlaceHandler{
		VotingPlaceService: votingPlaceService,
	}, nil
}

// GET
func (h *VotingPlaceHandler) List(c echo.Context) error {
	idStr := c.Param("id")

	electionID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.String(400, "invalid election id")
	}

	places, err := h.VotingPlaceService.GetByElection(uint(electionID))
	if err != nil {
		return c.String(500, "failed to fetch voting places")
	}

	data := map[string]interface{}{
		"Title":        "Puestos de Votación",
		"VotingPlaces": places,
		"ElectionID":   electionID,
	}

	// HTMX partial: return page-level partial for HTMX navigation
	if isHTMX(c) {
		return c.Render(http.StatusOK, "partials/voting_places_page.html", data)
	}

	return c.Render(http.StatusOK, "pages/voting_places.html", data)
}

// POST (HTMX-driven)
func (h *VotingPlaceHandler) Post(c echo.Context) error {
	idStr := c.Param("id")

	electionID, _ := strconv.ParseUint(idStr, 10, 32)

	action := c.FormValue("action")

	switch action {
	case "create":
		return h.create(c, uint(electionID))
	case "enable":
		return h.enable(c, uint(electionID))
	case "disable":
		return h.disable(c, uint(electionID))
	case "regenerate":
		return h.regenerate(c, uint(electionID))
	default:
		return c.String(400, "invalid action")
	}
}

func (h *VotingPlaceHandler) create(c echo.Context, electionID uint) error {
	place := &models.VotingPlace{
		Name:       c.FormValue("name"),
		Address:    c.FormValue("address"),
		ElectionID: electionID,
	}

	secret, err := h.VotingPlaceService.CreateVotingPlace(place)
	if err != nil {
		return c.String(500, "failed to create")
	}

	places, _ := h.VotingPlaceService.GetByElection(electionID)

	data := map[string]interface{}{
		"VotingPlaces": places,
		"ElectionID":   electionID,
		"Secret":       secret,
	}

	return c.Render(http.StatusOK, "partials/voting_places_table.html", data)
}

func (h *VotingPlaceHandler) enable(c echo.Context, electionID uint) error {
	id, _ := strconv.ParseUint(c.FormValue("place_id"), 10, 32)
	h.VotingPlaceService.EnableVotingPlace(uint(id))

	return h.renderTable(c, electionID)
}

func (h *VotingPlaceHandler) disable(c echo.Context, electionID uint) error {
	id, _ := strconv.ParseUint(c.FormValue("place_id"), 10, 32)
	h.VotingPlaceService.DisableVotingPlace(uint(id))

	return h.renderTable(c, electionID)
}

func (h *VotingPlaceHandler) regenerate(c echo.Context, electionID uint) error {
	id, _ := strconv.ParseUint(c.FormValue("place_id"), 10, 32)

	secret, _ := h.VotingPlaceService.RegenerateVotingPlaceSecret(uint(id))

	places, _ := h.VotingPlaceService.GetByElection(electionID)

	return c.Render(http.StatusOK, "partials/voting_places_table.html", map[string]interface{}{
		"VotingPlaces": places,
		"ElectionID":   electionID,
		"Secret":       secret,
	})
}

func (h *VotingPlaceHandler) renderTable(c echo.Context, electionID uint) error {
	places, _ := h.VotingPlaceService.GetByElection(electionID)

	return c.Render(http.StatusOK, "partials/voting_places_table.html", map[string]interface{}{
		"VotingPlaces": places,
		"ElectionID":   electionID,
	})
}
