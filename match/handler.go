package match

import (
	"net/http"
)

type matchService interface {
	Create() error
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

}
