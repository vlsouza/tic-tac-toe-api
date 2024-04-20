package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetString gets a request parameter as string
func GetUUID(r *http.Request, param string) (uuid.UUID, error) {
	p, ok := mux.Vars(r)[param]
	if !ok {
		return uuid.Nil, errors.New("param not found")
	}

	// Parse the string to a UUID
	id, err := uuid.Parse(p)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse UUID: %v", err)
	}

	return id, nil
}
