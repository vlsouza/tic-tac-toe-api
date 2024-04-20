package match

import (
	"fmt"
	"net/http"

	"main/match/service"
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

//func (h *Handler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
//}
