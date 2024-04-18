package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Service struct {
	DB *dynamodb.Client
}

// New ..
func New(db *dynamodb.Client) Service {
	return Service{db}
}

// Create ...
func (svc Service) Create(ctx context.Context) (CreateMatchResponse, error) {
	newMatchRequest, err := NewMatchRequest()
	if err != nil {
		//TODO set status code
		return CreateMatchResponse{}, err
	}

	// Usar PutItem para criar o registro
	_, err = svc.DB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("TicTacToeMatch"),
		Item:      newMatchRequest.DynamoRequest,
	})
	if err != nil {
		return CreateMatchResponse{}, err
	}

	fmt.Println("Match Created!")
	return CreateMatchResponse{
		MatchID: newMatchRequest.MatchID,
	}, nil
}
