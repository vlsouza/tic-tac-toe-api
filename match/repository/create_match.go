package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (r Repository) Create(ctx context.Context, match Match) (*dynamodb.PutItemOutput, error) {
	return r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: Match{}.awsTableName(),
		Item:      match.getDynamoRequest(),
	})
}
