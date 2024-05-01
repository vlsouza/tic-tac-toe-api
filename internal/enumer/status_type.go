package enumer

import (
	"fmt"

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

func (s StatusType) ToAttributeValue() types.AttributeValue {
	return &types.AttributeValueMemberS{Value: s.String()}
}

func (s *StatusType) FromAttributeValue(av types.AttributeValue) error {
	value, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return fmt.Errorf("expected string value for StatusType")
	}
	var err error
	*s, err = StatusTypeString(value.Value)
	return err
}
