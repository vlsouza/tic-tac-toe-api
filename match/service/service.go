package service

import (
	"context"
	"main/match/repository"

	"github.com/google/uuid"
)

type ServiceI interface {
	Create(context.Context) (GetStateResponse, error)
	GetStateByID(context.Context, uuid.UUID) (GetStateResponse, error)
	PlaceMove(context.Context, MoveRequest) (GetStateResponse, error)
	GetListByStatus(context.Context, string, int) ([]GetStateResponse, error)
}

type Service struct {
	repo repository.RepositoryI
}

// New ..
func New(repo repository.RepositoryI) Service {
	return Service{repo}
}
