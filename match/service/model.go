package service

import (
	"context"
	"main/match/repository"

	"github.com/google/uuid"
)

type match struct {
	ID uuid.UUID
	//TODO move to enum
	Status string
	Board  string
	//TODO move to enum
	CurrentPlayerTurn string
	NextPlayerTurn    string
	LastMoveXY        string
}

type CreateMatchResponse struct {
	MatchID uuid.UUID `json:"match_id"`
	//TODO move to enum
	Status string `json:"status,omitempty"`
	Board  string `json:"board,omitempty"`
	//TODO move to enum
	CurrentPlayerTurn string `json:"current_player_turn,omitempty"`
	NextPlayerTurn    string `json:"next_player_turn,omitempty"`
	LastMoveXY        string `json:"last_move_xy,omitempty"`
}

type GetStateResponse struct {
	CreateMatchResponse
}

func (match match) start(ctx context.Context, repo repository.RepositoryI) error {
	_, err := repo.Create(ctx,
		repository.Match{
			ID:                match.ID,
			Status:            match.Status,
			Board:             match.Board,
			CurrentPlayerTurn: match.CurrentPlayerTurn,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
