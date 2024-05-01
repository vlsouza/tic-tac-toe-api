package match

import (
	"encoding/json"
	"fmt"
	"net/http"

	"main/internal/service"
	"main/rest"
)

// Handler is used to aggregate all endpoints related
type Handler struct {
	service service.ServiceI
}

// NewHandler instance a new API handler
func NewHandler(service service.ServiceI) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	match, err := h.service.Create(r.Context())
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, match)
}

func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	matchState, err := h.service.GetStateByID(r.Context(), matchID)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, matchState)
}

func (h *Handler) GetListByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := rest.GetString(r, "status")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	matchState, err := h.service.GetListByStatus(r.Context(), status, 10 /* default, more values to be implemented */)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, matchState)
}

// improve name
func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	var moveRequest service.MoveRequest
	err = json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		//TODO handle error properly
		fmt.Fprint(w, err.Error())
	}
	moveRequest.MatchID = matchID

	if err := moveRequest.Validate(); err != nil {
		fmt.Fprint(w, err.Error())
	}

	newMatchState, err := h.service.PlaceMove(r.Context(), moveRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, newMatchState)
}

func (h *Handler) Start(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	err = h.service.Start(r.Context(), matchID)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
}