package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/web/templates"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type VotingPlaceHandler struct {
	VotingPlaceService *services.VotingPlaceService
	templates          *template.Template
}

func NewVotingPlaceHandler(votingPlaceService *services.VotingPlaceService) (*VotingPlaceHandler, error) {
	tmpl, err := templates.LoadTemplate("voting_places.html")
	if err != nil {
		return nil, err
	}

	return &VotingPlaceHandler{
		VotingPlaceService: votingPlaceService,
		templates:          tmpl,
	}, nil
}

func (h *VotingPlaceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleList(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *VotingPlaceHandler) handleList(w http.ResponseWriter, r *http.Request) {
	electionID, err := parseElectionID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	places, err := h.VotingPlaceService.GetByElection(electionID)
	if err != nil {
		http.Error(w, "Failed to fetch voting places", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"VotingPlaces": places,
		"ElectionID":   electionID,
	}

	if secret := r.URL.Query().Get("secret"); secret != "" {
		data["Secret"] = secret
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func (h *VotingPlaceHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	electionID, err := parseElectionID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	action := r.FormValue("action")
	if action == "create" {
		h.handleCreate(w, r, electionID)
		return
	}

	if action == "disable" {
		h.handleDisable(w, r)
		return
	}

	if action == "enable" {
		h.handleEnable(w, r)
		return
	}

	if action == "regenerate" {
		h.handleRegenerate(w, r)
		return
	}

	http.Error(w, "Invalid action", http.StatusBadRequest)
}

func (h *VotingPlaceHandler) handleCreate(w http.ResponseWriter, r *http.Request, electionID uint) {
	name := r.FormValue("name")
	address := r.FormValue("address")

	place := &models.VotingPlace{
		Name:       name,
		Address:    address,
		ElectionID: electionID,
	}

	secret, err := h.VotingPlaceService.CreateVotingPlace(place)
	if err != nil {
		http.Error(w, "Failed to create voting place", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/voting-places?election_id="+strconv.FormatUint(uint64(electionID), 10)+"&secret="+secret, http.StatusSeeOther)
}

func (h *VotingPlaceHandler) handleDisable(w http.ResponseWriter, r *http.Request) {
	placeIDStr := r.FormValue("place_id")
	placeID, err := strconv.ParseUint(placeIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid place_id", http.StatusBadRequest)
		return
	}

	if err := h.VotingPlaceService.DisableVotingPlace(uint(placeID)); err != nil {
		http.Error(w, "Failed to disable voting place", http.StatusInternalServerError)
		return
	}

	redirectURL := r.Referer()
	if redirectURL == "" {
		redirectURL = "/elections"
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (h *VotingPlaceHandler) handleEnable(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("place_id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.VotingPlaceService.EnableVotingPlace(uint(id)); err != nil {
		http.Error(w, "Failed to enable", 500)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func (h *VotingPlaceHandler) handleRegenerate(w http.ResponseWriter, r *http.Request) {
	placeIDStr := r.FormValue("place_id")
	placeID, err := strconv.ParseUint(placeIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid place_id", http.StatusBadRequest)
		return
	}

	secret, err := h.VotingPlaceService.RegenerateVotingPlaceSecret(uint(placeID))
	if err != nil {
		http.Error(w, "Failed to regenerate voting place secret", http.StatusInternalServerError)
		return
	}

	place, err := h.VotingPlaceService.GetByID(uint(placeID))
	if err != nil {
		http.Error(w, "Failed to refresh voting place", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/voting-places?election_id="+strconv.FormatUint(uint64(place.ElectionID), 10)+"&secret="+secret, http.StatusSeeOther)
}

func parseElectionID(r *http.Request) (uint, error) {
	electionIDStr := r.URL.Query().Get("election_id")
	if electionIDStr == "" {
		electionIDStr = r.FormValue("election_id")
	}

	if electionIDStr == "" {
		return 0, fmt.Errorf("election_id is required")
	}

	electionID, err := strconv.ParseUint(electionIDStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid election_id")
	}

	return uint(electionID), nil
}
