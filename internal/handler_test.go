package internal

import (
	"errors"
	"main/internal/service"
	"main/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

//This is an example of a unit test for a handler. In the ideal scenario, it is necessary to have all handler functions tested following the same pattern.

func TestHandler_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceI(ctrl)
	mockService.EXPECT().Create(gomock.Any()).Return(service.GetStateResponse{}, nil)

	handler := &Handler{service: mockService}

	req, err := http.NewRequest("POST", "/matches", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandler_Create_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceI(ctrl)
	mockService.EXPECT().Create(gomock.Any()).Return(service.GetStateResponse{}, errors.New("Internal error"))

	handler := &Handler{service: mockService}

	req, err := http.NewRequest("POST", "/matches", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.Create(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
