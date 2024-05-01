package enumer

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

//go:generate enumer -json -text -sql -type PlayerType -trimprefix Player -transform snake-upper

// Status represents the available players in a match
type PlayerType int

const (
	PLAYER1 PlayerType = iota
	PLAYER2
)

func (p PlayerType) ToAttributeValue() types.AttributeValue {
	return &types.AttributeValueMemberS{Value: p.String()}
}

func (p *PlayerType) FromAttributeValue(av types.AttributeValue) error {
	value, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return fmt.Errorf("expected string value for PlayerType")
	}
	var err error
	*p, err = PlayerTypeString(value.Value)
	return err
}
