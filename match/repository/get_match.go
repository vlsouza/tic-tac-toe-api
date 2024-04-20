package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func (r Repository) GetByID(ctx context.Context, matchID uuid.UUID) (Match, error) {
	var response Match = Match{}

	result, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: Match{}.awsTableName(),
		Key: map[string]types.AttributeValue{
			"match_id": &types.AttributeValueMemberS{Value: matchID.String()},
		},
	})
	if err != nil {
		return Match{}, fmt.Errorf("failed to get item from DynamoDB, %v", err)
	}

	if result.Item == nil {
		return Match{}, fmt.Errorf("no item found with MatchID: %s", matchID.String())
	}

	err = attributevalue.UnmarshalMap(result.Item, &response)
	if err != nil {
		log.Fatalf("failed to unmarshal DynamoDB item to struct, %v", err)
	}

	return response, nil
}
