package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
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

func (r Repository) GetListByStatus(
	ctx context.Context,
	status string,
	limit int,
) ([]Match, error) {
	var response []Match

	results, err := getListByStatus(r.db, status, limit)
	if err != nil {
		return []Match{}, fmt.Errorf("failed to get item from DynamoDB, %v", err)
	}

	if results == nil {
		return []Match{}, errors.New("no available matches")
	}

	err = attributevalue.UnmarshalListOfMaps(results.Items, &response)
	if err != nil {
		log.Fatalf("failed to unmarshal DynamoDB item to struct, %v", err)
	}

	return response, nil
}

func getListByStatus(
	svc *dynamodb.Client,
	status string,
	limit int,
) (*dynamodb.QueryOutput, error) {
	// Construa a condição de chave e a expressão de projeção
	keyCond := expression.Key("status").Equal(expression.Value(status))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	// Executa a consulta
	return svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 Match{}.awsTableName(),
		IndexName:                 aws.String("match_id"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		Limit:                     aws.Int32(int32(limit)), // Limita a 5 resultados
		ScanIndexForward:          aws.Bool(false),
	})
}
