package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RepositoryI interface {
	Create(ctx context.Context, dbRequest CreateDynamoRequest) (*dynamodb.PutItemOutput, error)
}

type CreateDynamoRequest map[string]types.AttributeValue

type Repository struct {
	DB *dynamodb.Client
}

func (r Repository) Create(ctx context.Context, dbRequest CreateDynamoRequest) (*dynamodb.PutItemOutput, error) {
	return r.DB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("TicTacToeMatch"),
		Item:      dbRequest,
	})
}
