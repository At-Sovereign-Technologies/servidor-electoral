package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web/templates"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type ElectionHandler struct {
	ElectionService *services.ElectionService
	templates       *template.Template
}

func NewElectionHandler(electionService *services.ElectionService) (*ElectionHandler, error) {
	tmpl, err := templates.LoadTemplate("elections.html")
	if err != nil {
		return nil, err
	}

	return &ElectionHandler{
		ElectionService: electionService,
		templates:       tmpl,
	}, nil
}

func (h *ElectionHandler) ListElections(w http.ResponseWriter, r *http.Request) {
	elections, err := h.ElectionService.ElectionStore.List()
	if err != nil {
		http.Error(w, "Failed to fetch elections", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Execute(w, map[string]interface{}{
		"elections": elections,
	}); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *ElectionHandler) GetElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	election, err := h.ElectionService.ElectionStore.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Election not found", http.StatusNotFound)
		return
	}

	tmpl, err := templates.LoadTemplate("election_detail.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Election": election,
		"Start":    election.StartDate.Format("02/01/2006 15:04"),
		"End":      election.EndDate.Format("02/01/2006 15:04"),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}
