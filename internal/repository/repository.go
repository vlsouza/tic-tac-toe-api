package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type RepositoryI interface {
	Create(context.Context, Match) (*dynamodb.PutItemOutput, error)
	GetByID(context.Context, uuid.UUID) (Match, error)
	Update(context.Context, Match) (*dynamodb.UpdateItemOutput, error)
	GetListByStatus(context.Context, string, int) ([]Match, error)
}

type Repository struct {
	db *dynamodb.Client
}

func New(db *dynamodb.Client) Repository {
	return Repository{db: db}
}
