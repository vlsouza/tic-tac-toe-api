package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// SendJSON is a helper function to send a JSON as an HTTP response.
// It sets the header Content-Type as application/json.
// If it fails to write the JSON an internal server error is generated.
func SendJSON(w http.ResponseWriter, payload interface{}) error {
	return writeJSON(context.Background(), w, payload)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return errors.Wrap(err, "rest: writeJSON")
	}

	return nil
}
