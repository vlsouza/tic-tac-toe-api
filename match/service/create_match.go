package service

import (
	"context"

	"main/match/repository"

	"github.com/google/uuid"
)

// Create ...
func (svc Service) Create(ctx context.Context) (GetStateResponse, error) {
	return startNewMatch(ctx, svc.repo)
}

func startNewMatch(ctx context.Context, repo repository.RepositoryI) (GetStateResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return GetStateResponse{}, err
	}

	//new match instance
	match := CreateMatchRequest{
		ID:                id,
		Status:            "PENDINGPLAYER",
		Board:             "000,000,000",
		CurrentPlayerTurn: "PLAYER1",
		NextPlayerTurn:    "PLAYER2",
		LastMoveXY:        "",
	}

	//start match on dynamoDB
	err = match.start(ctx, repo)
	if err != nil {
		return GetStateResponse{}, err
	}

	return GetStateResponse{
		MatchID:           match.ID,
		Status:            match.Status,
		Board:             match.Board,
		CurrentPlayerTurn: match.CurrentPlayerTurn,
		NextPlayerTurn:    match.NextPlayerTurn,
		LastMoveXY:        match.LastMoveXY,
	}, nil
}
