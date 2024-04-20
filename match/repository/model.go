package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

// DB model
type Match struct {
	ID uuid.UUID
	//TODO move to enum
	Status string
	Board  string
	//TODO move to enum
	CurrentPlayerTurn string
	NextPlayerTurn    string
	LastMoveXY        string
}

const TableName string = "TicTacToeMatch"

func (Match) awsTableName() *string {
	return aws.String(TableName)
}

func (m Match) getDynamoRequest() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"match_id":            &types.AttributeValueMemberS{Value: m.ID.String()},
		"status":              &types.AttributeValueMemberS{Value: m.Status},
		"board":               &types.AttributeValueMemberS{Value: m.Board},
		"current_player_turn": &types.AttributeValueMemberS{Value: m.CurrentPlayerTurn},
		"next_player_turn":    &types.AttributeValueMemberS{Value: m.NextPlayerTurn},
		"last_move_xy":        &types.AttributeValueMemberS{Value: m.LastMoveXY},
	}
}
