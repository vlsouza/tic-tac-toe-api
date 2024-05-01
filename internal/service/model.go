package service

import (
	"context"

	"main/internal/repository"

	"github.com/google/uuid"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// TODO move status and players to enum
type CreateMatchRequest struct {
	ID                uuid.UUID
	Status            string
	Board             string
	CurrentPlayerTurn string
	NextPlayerTurn    string
	LastMoveXY        string
}

type CreateMatchResponse struct {
	MatchID uuid.UUID `json:"match_id"`
}

type GetStateResponse struct {
	MatchID           uuid.UUID `json:"match_id"`
	Status            string    `json:"status"`
	Board             string    `json:"board"`
	CurrentPlayerTurn string    `json:"current_player_turn"`
	NextPlayerTurn    string    `json:"next_player_turn"`
	LastMoveXY        string    `json:"last_move_xy"`
}

type GetStateResponseList []GetStateResponse

type MoveRequest struct {
	MatchID uuid.UUID
	Row     int8 `json:"row"`
	Col     int8 `json:"col"`
}

func (match CreateMatchRequest) start(ctx context.Context, repo repository.RepositoryI) error {
	_, err := repo.Create(ctx,
		repository.Match{
			ID:                match.ID.String(),
			Status:            match.Status,
			Board:             match.Board,
			CurrentPlayerTurn: match.CurrentPlayerTurn,
			NextPlayerTurn:    match.NextPlayerTurn,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewGetStateResponse(match repository.Match) (GetStateResponse, error) {
	matchStateMatchID, err := uuid.Parse(match.ID)
	if err != nil {
		return GetStateResponse{}, err
	}
	return GetStateResponse{
		MatchID:           matchStateMatchID,
		Status:            match.Status,
		Board:             match.Board,
		CurrentPlayerTurn: match.CurrentPlayerTurn,
		NextPlayerTurn:    match.NextPlayerTurn,
		LastMoveXY:        match.LastMoveXY,
	}, nil
}

func NewGetStateResponseList(matches []repository.Match) ([]GetStateResponse, error) {
	response := make([]GetStateResponse, len(matches))
	for i := range matches {
		match := matches[i]
		matchStateMatchID, err := uuid.Parse(match.ID)
		if err != nil {
			return []GetStateResponse{}, err
		}

		response[i] = GetStateResponse{
			MatchID:           matchStateMatchID,
			Status:            match.Status,
			Board:             match.Board,
			CurrentPlayerTurn: match.CurrentPlayerTurn,
			NextPlayerTurn:    match.NextPlayerTurn,
			LastMoveXY:        match.LastMoveXY,
		}
	}

	return response, nil
}

func (a MoveRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Row, validation.Required, validation.Min(1), validation.Max(3)),
		validation.Field(&a.Col, validation.Required, validation.Min(1), validation.Max(3)),
	)
}
