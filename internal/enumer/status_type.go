package enumer

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

//go:generate enumer -json -text -sql -type StatusType -trimprefix Status -transform snake-upper

// Status represents the current status of a game
type StatusType int

const (
	PENDINGPLAYER StatusType = iota
	RUNNING
	PLAYER1WON
	PLAYER2WON
	DRAW
)

func (p StatusType) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	s := p.String() // Ensure that String method returns the correct string representation
	return &types.AttributeValueMemberS{Value: s}, nil
}

func (p *StatusType) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	// Type assert to *types.AttributeValueMemberS to handle string values
	if s, ok := av.(*types.AttributeValueMemberS); ok {
		var err error
		*p, err = StatusTypeString(s.Value)
		return err
	}
	return errors.New("expected string value for StatusType")
}
