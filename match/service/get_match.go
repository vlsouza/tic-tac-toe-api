package service

import (
	"context"

	"github.com/google/uuid"
)

// Get state by match ID
func (svc Service) GetStateByID(ctx context.Context, matchID uuid.UUID) (GetStateResponse, error) {
	matchState, err := svc.repo.GetState(ctx, matchID)
	if err != nil {
		return GetStateResponse{}, err
	}

	return GetStateResponse{
		CreateMatchResponse{
			MatchID:           matchState.ID,
			Status:            matchState.Status,
			Board:             matchState.Board,
			CurrentPlayerTurn: matchState.CurrentPlayerTurn,
			NextPlayerTurn:    matchState.NextPlayerTurn,
			LastMoveXY:        matchState.LastMoveXY,
		},
	}, nil
}
