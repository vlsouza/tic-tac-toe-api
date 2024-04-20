package service

import (
	"context"
	"main/match/repository"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type MatchI interface {
	Start() error
}

type match struct {
	ID uuid.UUID
	//TODO move to enum
	Status string
	Board  string
	//TODO move to enum
	CurrentPlayerTurn string
	NextPlayerTurn    string
	LastMoveXY        string
	DynamoRequest     repository.CreateDynamoRequest
}

type CreateMatchResponse struct {
	MatchID uuid.UUID
}

func StartNewMatch(ctx context.Context, repo repository.RepositoryI) (CreateMatchResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return CreateMatchResponse{}, err
	}

	//initial match values
	var (
		matchID           = id
		status            = "Running"
		board             = "000,000,000"
		currentPlayerTurn = "Player1"
		NextPlayerTurn    = "Player2"
		LastMoveXY        = ""
	)

	// Definir o item que será criado
	dynamoRequest := map[string]types.AttributeValue{
		"match_id":            &types.AttributeValueMemberS{Value: id.String()},
		"status":              &types.AttributeValueMemberS{Value: "Running"},
		"board":               &types.AttributeValueMemberS{Value: "000,000,000"},
		"current_player_turn": &types.AttributeValueMemberS{Value: "Player1"},
		"next_player_turn":    &types.AttributeValueMemberS{Value: "Player2"},
		"last_move_xy":        &types.AttributeValueMemberS{Value: ""},
	}

	match := match{
		ID:                matchID,
		Status:            status,
		Board:             board,
		CurrentPlayerTurn: currentPlayerTurn,
		NextPlayerTurn:    NextPlayerTurn,
		LastMoveXY:        LastMoveXY,
		DynamoRequest:     dynamoRequest,
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

func (match match) start(ctx context.Context, repo repository.RepositoryI) error {
	_, err := repo.Create(ctx, match.DynamoRequest)
	if err != nil {
		return err
	}
	return nil
}
