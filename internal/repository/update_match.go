package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r Repository) Update(ctx context.Context, match Match) (*dynamodb.UpdateItemOutput, error) {
	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(
			expression.Name("status"), expression.Value(match.Status),
		).Set(
			expression.Name("board"), expression.Value(match.Board),
		).Set(
			expression.Name("current_player_turn"), expression.Value(match.CurrentPlayerTurn.ToAttributeValue()),
		).Set(
			expression.Name("next_player_turn"), expression.Value(match.NextPlayerTurn.ToAttributeValue()),
		).Set(
			expression.Name("last_move_xy"), expression.Value(match.LastMoveXY),
		),
	).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: Match{}.awsTableName(),
		Key: map[string]types.AttributeValue{
			"match_id": &types.AttributeValueMemberS{Value: match.ID},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	return r.db.UpdateItem(ctx, input)
}
