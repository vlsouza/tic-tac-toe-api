package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DB model
type Match struct {
	ID                string `dynamodbav:"match_id"`
	Status            string `dynamodbav:"status"`
	Board             string `dynamodbav:"board"`
	CurrentPlayerTurn int8   `dynamodbav:"current_player_turn"`
	NextPlayerTurn    int8   `dynamodbav:"next_player_turn"`
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
		"current_player_turn": &types.AttributeValueMemberS{Value: fmt.Sprint(m.CurrentPlayerTurn)},
		"next_player_turn":    &types.AttributeValueMemberS{Value: fmt.Sprint(m.NextPlayerTurn)},
		"last_move_xy":        &types.AttributeValueMemberS{Value: m.LastMoveXY},
	}
}

func (m Match) UpdateBoard(
	ctx context.Context,
	player, row, col int,
) string {
	rows := strings.Split(m.Board, ",")
	for i, rowContent := range rows {
		cells := strings.Split(rowContent, "")
		if i == row {
			cells[col] = fmt.Sprintf("%d", player)
			rows[i] = strings.Join(cells, "")
		}
	}

	return strings.Join(rows, ",")
}
