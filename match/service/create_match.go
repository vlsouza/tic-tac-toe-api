package service

import (
	"context"

	"main/match/repository"

	"github.com/google/uuid"
)

// Create ...
func (svc Service) Create(ctx context.Context) (CreateMatchResponse, error) {
	return startNewMatch(ctx, svc.repo)
}

func startNewMatch(ctx context.Context, repo repository.RepositoryI) (CreateMatchResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return CreateMatchResponse{}, err
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
		return CreateMatchResponse{}, err
	}

	return CreateMatchResponse{
		MatchID: match.ID,
	}, nil
}
