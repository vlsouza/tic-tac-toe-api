package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DB model
type Match struct {
	ID                string `dynamodbav:"match_id"`
	Status            string `dynamodbav:"status"`
	Board             string `dynamodbav:"board"`
	CurrentPlayerTurn string `dynamodbav:"current_player_turn"`
	NextPlayerTurn    string `dynamodbav:"next_player_turn"`
	LastMoveXY        string `dynamodbav:"last_move_xy"`
}

const TableName string = "TicTacToeMatch"

func (Match) awsTableName() *string {
	return aws.String(TableName)
}

func (m Match) getDynamoRequest() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"match_id":            &types.AttributeValueMemberS{Value: m.ID},
		"status":              &types.AttributeValueMemberS{Value: m.Status},
		"board":               &types.AttributeValueMemberS{Value: m.Board},
		"current_player_turn": &types.AttributeValueMemberS{Value: m.CurrentPlayerTurn},
		"next_player_turn":    &types.AttributeValueMemberS{Value: m.NextPlayerTurn},
		"last_move_xy":        &types.AttributeValueMemberS{Value: m.LastMoveXY},
	}
}
