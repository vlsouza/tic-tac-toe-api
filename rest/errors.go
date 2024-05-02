package rest

import (
	"errors"
	"fmt"
	"main/internal/service"
	"net/http"
)

func InvalidParameter(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("Invalid parameter: %s", err.Error()), http.StatusBadRequest)
}

func InternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func ValidationFailed(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("Request validation failed: %s", err.Error()), http.StatusBadRequest)
}

func HandleError(w http.ResponseWriter, err error) {
	if errors.Is(err, service.ErrNotFound) {
		http.Error(w, fmt.Sprintf("Resource not found: %s", err.Error()), http.StatusNotFound)
	} else if errors.Is(err, service.ErrMatchNotRunnig) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if errors.Is(err, service.ErrMatchAlreadyRunnig) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
