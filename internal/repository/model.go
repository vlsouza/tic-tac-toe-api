package repository

import (
	"main/internal/enumer"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DB model
type Match struct {
	ID                string            `dynamodbav:"match_id"`
	Status            enumer.StatusType `dynamodbav:"status"`
	Board             string            `dynamodbav:"board"`
	CurrentPlayerTurn enumer.PlayerType `dynamodbav:"current_player_turn"`
	NextPlayerTurn    enumer.PlayerType `dynamodbav:"next_player_turn"`
	LastMoveXY        string            `dynamodbav:"last_move_xy"`
}

const TableName string = "tic_tac_toe_match"

func (Match) awsTableName() *string {
	return aws.String(TableName)
}

func (m Match) getDynamoRequest() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"match_id":            &types.AttributeValueMemberS{Value: m.ID},
		"status":              &types.AttributeValueMemberS{Value: m.Status.String()},
		"board":               &types.AttributeValueMemberS{Value: m.Board},
		"current_player_turn": &types.AttributeValueMemberS{Value: m.CurrentPlayerTurn.String()},
		"next_player_turn":    &types.AttributeValueMemberS{Value: m.NextPlayerTurn.String()},
		"last_move_xy":        &types.AttributeValueMemberS{Value: m.LastMoveXY},
	}
}
