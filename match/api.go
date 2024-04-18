package match

import (
	"net/http"

	"main/match/service"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
)

// Config is used to setup the API.
type Config struct {
	DB     *dynamodb.Client
	Router *mux.Router
}

// New is used to initialize the API.
func NewAPI(c Config) {
	handler := NewHandler(service.New(c.DB))
	SetRoutes(handler, c.Router)
}

// SetRoutes is used to declare all endpoints managed by this API.
func SetRoutes(handler *Handler, router *mux.Router) {
	router.HandleFunc("/matches", handler.Create).Methods(http.MethodPost)
}
