package enumer

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

//go:generate enumer -json -text -sql -type PlayerType -trimprefix Player -transform snake-upper

// Status represents the available players in a match
type PlayerType int

const (
	PLAYER1 PlayerType = iota
	PLAYER2
)

func (p PlayerType) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	s := p.String() // Ensure that String method returns the correct string representation
	return &types.AttributeValueMemberS{Value: s}, nil
}

func (p *PlayerType) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	// Type assert to *types.AttributeValueMemberS to handle string values
	if s, ok := av.(*types.AttributeValueMemberS); ok {
		var err error
		*p, err = PlayerTypeString(s.Value)
		return err
	}
	return errors.New("expected string value for PlayerType")
}
