package handlers

import (
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
	"github.com/labstack/echo/v4"
)

type WorkerHandler struct {
	workerService *services.WorkerService
}

func NewWorkerHandler(workerService *services.WorkerService) (*WorkerHandler, error) {
	return &WorkerHandler{
		workerService: workerService,
	}, nil
}

// GET
func (h *WorkerHandler) List(c echo.Context) error {
	idStr := c.Param("id")

	electionID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.String(400, "invalid election id")
	}

	workers, err := h.workerService.GetByElection(uint(electionID))
	if err != nil {
		return c.String(500, "failed to fetch workers")
	}

	data := map[string]interface{}{
		"Title":      "Puestos de Votación",
		"Workers":    workers,
		"ElectionID": electionID,
	}

	// HTMX partial: return page-level partial for HTMX navigation
	if isHTMX(c) {
		return c.Render(http.StatusOK, "partials/workers_page.html", data)
	}

	return c.Render(http.StatusOK, "pages/workers.html", data)
}

// POST (HTMX-driven)
func (h *WorkerHandler) Post(c echo.Context) error {
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
	default:
		return c.String(400, "invalid action")
	}
}

func (h *WorkerHandler) create(c echo.Context, electionID uint) error {
	config, err := h.workerService.CreateWorker(c.FormValue("name"), electionID)
	if err != nil {
		return c.String(500, "failed to create")
	}

	workers, err := h.workerService.GetByElection(uint(electionID))
	if err != nil {
		return c.String(500, "failed to fetch workers")
	}

	data := map[string]interface{}{
		"Workers":      workers,
		"ElectionID":   electionID,
		"WorkerConfig": config,
	}

	return c.Render(http.StatusOK, "partials/workers_page.html", data)
}

func (h *WorkerHandler) enable(c echo.Context, electionID uint) error {
	id, _ := strconv.ParseUint(c.FormValue("worker_id"), 10, 32)
	h.workerService.EnableWorker(uint(id))

	return h.renderTable(c, electionID)
}

func (h *WorkerHandler) disable(c echo.Context, electionID uint) error {
	id, _ := strconv.ParseUint(c.FormValue("place_id"), 10, 32)
	h.workerService.DeleteWorker(uint(id))

	return h.renderTable(c, electionID)
}

func (h *WorkerHandler) renderTable(c echo.Context, electionID uint) error {
	workers, _ := h.workerService.GetByElection(electionID)

	return c.Render(http.StatusOK, "partials/workers_table.html", map[string]interface{}{
		"Workers":    workers,
		"ElectionID": electionID,
	})
}
