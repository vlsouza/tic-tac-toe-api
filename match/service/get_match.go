package service

import (
	"context"

	"github.com/google/uuid"
)

// Get state by match ID
func (svc Service) GetStateByID(ctx context.Context, matchID uuid.UUID) (GetStateResponse, error) {
	match, err := svc.repo.GetByID(ctx, matchID)
	if err != nil {
		return GetStateResponse{}, err
	}

	return NewGetStateResponse(match), nil
}
