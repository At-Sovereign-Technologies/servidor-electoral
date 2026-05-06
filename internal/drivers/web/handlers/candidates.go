package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web/templates"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type CandidateHandler struct {
	CandidateService *services.CandidateService
	templates        *template.Template
}

func NewCandidateHandler(candidateService *services.CandidateService) (*CandidateHandler, error) {
	tmpl, err := templates.LoadTemplate("candidates.html")
	if err != nil {
		return nil, err
	}

	return &CandidateHandler{
		CandidateService: candidateService,
		templates:        tmpl,
	}, nil
}

func (h *CandidateHandler) ListCandidates(w http.ResponseWriter, r *http.Request) {
	electionIDStr := r.URL.Query().Get("election_id")
	if electionIDStr == "" {
		http.Error(w, "election_id is required", http.StatusBadRequest)
		return
	}

	electionID, err := strconv.ParseUint(electionIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid election_id", http.StatusBadRequest)
		return
	}

	candidates, err := h.CandidateService.CandidateStore.ListByElection(uint(electionID))
	if err != nil {
		http.Error(w, "Failed to fetch candidates", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Execute(w, map[string]interface{}{
		"candidates": candidates,
		"electionID": electionID,
	}); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}
