package match

import (
	"net/http"
)

type service interface {
	Create() error
}

// Handler is used to aggregate all endpoints related
type Handler struct {
	svc service
}

// NewHandler instance a new API handler
func NewHandler(svc service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}
