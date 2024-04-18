package service

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type CreateMatchRequest struct {
	MatchID uuid.UUID
	//TODO move to enum
	Status string
	Board  string
	//TODO move to enum
	CurrentPlayerTurn string
	NextPlayerTurn    string
	LastMoveXY        string
	DynamoRequest     map[string]types.AttributeValue
}

type CreateMatchResponse struct {
	MatchID uuid.UUID
}

func NewMatchRequest() (CreateMatchRequest, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return CreateMatchRequest{}, err
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
		"MatchID":           &types.AttributeValueMemberS{Value: id.String()},
		"Status":            &types.AttributeValueMemberS{Value: "Running"},
		"Board":             &types.AttributeValueMemberS{Value: "000,000,000"},
		"CurrentPlayerTurn": &types.AttributeValueMemberS{Value: "Player1"},
		"NextPlayerTurn":    &types.AttributeValueMemberS{Value: "Player2"},
		"LastMoveXY":        &types.AttributeValueMemberS{Value: ""},
	}

	return CreateMatchRequest{
		MatchID:           matchID,
		Status:            status,
		Board:             board,
		CurrentPlayerTurn: currentPlayerTurn,
		NextPlayerTurn:    NextPlayerTurn,
		LastMoveXY:        LastMoveXY,
		DynamoRequest:     dynamoRequest,
	}, nil
}
