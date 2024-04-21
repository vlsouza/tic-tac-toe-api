package service

import (
	"context"
	"main/match/repository"

	"github.com/google/uuid"
)

type ServiceI interface {
	Create(ctx context.Context) (CreateMatchResponse, error)
	GetStateByID(ctx context.Context, matchID uuid.UUID) (GetStateResponse, error)
	PlaceMove(ctx context.Context, request MoveRequest) (GetStateResponse, error)
}

type Service struct {
	repo repository.RepositoryI
}

// New ..
func New(repo repository.RepositoryI) Service {
	return Service{repo}
}
