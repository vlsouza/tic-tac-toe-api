package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type RepositoryI interface {
	Create(ctx context.Context, match Match) (*dynamodb.PutItemOutput, error)
	GetState(ctx context.Context, matchID uuid.UUID) (Match, error)
}

type Repository struct {
	db *dynamodb.Client
}

func New(db *dynamodb.Client) Repository {
	return Repository{db: db}
}
