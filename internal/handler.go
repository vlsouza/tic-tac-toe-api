package internal

import (
	"encoding/json"
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
		rest.InternalError(w, err)
	}

	rest.SendJSON(w, match)
}

func (h *Handler) GetState(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		rest.InvalidParameter(w, err)
	}

	matchState, err := h.service.GetStateByID(r.Context(), matchID)
	if err != nil {
		rest.HandleError(w, err)
	}

	rest.SendJSON(w, matchState)
}

func (h *Handler) GetListByStatus(w http.ResponseWriter, r *http.Request) {
	status, err := rest.GetString(r, "status")
	if err != nil {
		rest.InvalidParameter(w, err)
	}

	matchState, err := h.service.GetListByStatus(r.Context(), status, 10 /* default */)
	if err != nil {
		rest.InternalError(w, err)
	}

	rest.SendJSON(w, matchState)
}

func (h *Handler) Move(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		rest.InvalidParameter(w, err)
	}

	var moveRequest service.MoveRequest
	err = json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		rest.InternalError(w, err)
	}
	moveRequest.MatchID = matchID

	if err := moveRequest.Validate(); err != nil {
		rest.ValidationFailed(w, err)
	}

	newMatchState, err := h.service.PlaceMove(r.Context(), moveRequest)
	if err != nil {
		rest.HandleError(w, err)
	}

	rest.SendJSON(w, newMatchState)
}

func (h *Handler) Start(w http.ResponseWriter, r *http.Request) {
	matchID, err := rest.GetUUID(r, "id")
	if err != nil {
		rest.InvalidParameter(w, err)
	}

	err = h.service.Start(r.Context(), matchID)
	if err != nil {
		rest.HandleError(w, err)
	}
}
