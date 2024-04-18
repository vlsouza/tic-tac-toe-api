package match

import (
	"context"
	"fmt"
	"net/http"

	"main/match/service"
	"main/rest"
)

type matchService interface {
	Create(ctx context.Context) (service.CreateMatchResponse, error)
}

// Handler is used to aggregate all endpoints related
type Handler struct {
	service matchService
}

// NewHandler instance a new API handler
func NewHandler(service matchService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	match, err := h.service.Create(r.Context())
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, match)
}
