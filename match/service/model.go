package service

import (
	"context"
	"main/match/repository"

	"github.com/google/uuid"
)

// TODO move status and players to enum
type CreateMatchRequest struct {
	ID                uuid.UUID
	Status            string
	Board             string
	CurrentPlayerTurn int8
	NextPlayerTurn    int8
	LastMoveXY        string
}

type CreateMatchResponse struct {
	MatchID uuid.UUID `json:"match_id"`
}

type GetStateResponse struct {
	MatchID           uuid.UUID `json:"match_id"`
	Status            string    `json:"status"`
	Board             string    `json:"board"`
	CurrentPlayerTurn int8      `json:"current_player_turn"`
	NextPlayerTurn    int8      `json:"next_player_turn"`
	LastMoveXY        string    `json:"last_move_xy"`
}

type MoveRequest struct {
	MatchID uuid.UUID `json:"match_id"`
	Player  int8      `json:"player"`
	Row     int8      `json:"row"`
	Col     int8      `json:"col"`
}

func (match CreateMatchRequest) start(ctx context.Context, repo repository.RepositoryI) error {
	_, err := repo.Create(ctx,
		repository.Match{
			ID:                match.ID.String(),
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
