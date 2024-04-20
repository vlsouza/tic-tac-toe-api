package service

import (
	"context"
	"main/match/repository"
)

type Service struct {
	repo repository.RepositoryI
}

// New ..
func New(repo repository.RepositoryI) Service {
	return Service{repo}
}

// Create ...
func (svc Service) Create(ctx context.Context) (CreateMatchResponse, error) {
	return StartNewMatch(ctx, svc.repo)
}
