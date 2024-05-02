package service

import (
	"context"
	"errors"
	"main/internal/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

//This is an example of a unit test for a handler. In the ideal scenario, it is necessary to have all handler functions tested following the same pattern.

func TestService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryI(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)

	svc := Service{repo: mockRepo}

	ctx := context.Background()
	_, err := svc.Create(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryI(ctrl)
	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Internal error"))

	svc := Service{repo: mockRepo}

	ctx := context.Background()
	_, err := svc.Create(ctx)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
