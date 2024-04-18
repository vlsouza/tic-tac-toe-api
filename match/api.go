package match

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Config is used to setup the API.
type Config struct {
	Router *mux.Router
}

// New is used to initialize the API.
func New(c Config) {
	handler := NewHandler(nil)
	SetRoutes(handler, c.Router)
}

// SetRoutes is used to declare all endpoints managed by this API.
func SetRoutes(handler *Handler, router *mux.Router) {
	router.HandleFunc("/matches", handler.Create).Methods(http.MethodPost)
}
