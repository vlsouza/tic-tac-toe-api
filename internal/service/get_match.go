package service

import (
	"context"
	"errors"
	"main/internal/repository"

	"github.com/google/uuid"
)

// Get state by match ID
func (svc Service) GetStateByID(ctx context.Context, matchID uuid.UUID) (GetStateResponse, error) {
	match, err := svc.repo.GetByID(ctx, matchID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return GetStateResponse{}, ErrNotFound
		}

		return GetStateResponse{}, err
	}

	return NewGetStateResponse(match)
}

// Get list of states of pending players matches given the status and result amount limit
// default status: 'RUNNING'
// limit: 5
func (svc Service) GetListByStatus(
	ctx context.Context,
	status string,
	limit int,
) ([]GetStateResponse, error) {
	match, err := svc.repo.GetListByStatus(ctx, status, limit)
	if err != nil {
		return []GetStateResponse{}, err
	}

	return NewGetStateResponseList(match)
}
