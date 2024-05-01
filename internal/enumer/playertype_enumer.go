// Code generated by "enumer -json -text -sql -type PlayerType -trimprefix Player -transform snake-upper"; DO NOT EDIT.

package enumer

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _PlayerTypeName = "PLAYER1PLAYER2"

var _PlayerTypeIndex = [...]uint8{0, 7, 14}

const _PlayerTypeLowerName = "player1player2"

func (i PlayerType) String() string {
	if i < 0 || i >= PlayerType(len(_PlayerTypeIndex)-1) {
		return fmt.Sprintf("PlayerType(%d)", i)
	}
	return _PlayerTypeName[_PlayerTypeIndex[i]:_PlayerTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PlayerTypeNoOp() {
	var x [1]struct{}
	_ = x[PLAYER1-(0)]
	_ = x[PLAYER2-(1)]
}

var _PlayerTypeValues = []PlayerType{PLAYER1, PLAYER2}

var _PlayerTypeNameToValueMap = map[string]PlayerType{
	_PlayerTypeName[0:7]:       PLAYER1,
	_PlayerTypeLowerName[0:7]:  PLAYER1,
	_PlayerTypeName[7:14]:      PLAYER2,
	_PlayerTypeLowerName[7:14]: PLAYER2,
}

var _PlayerTypeNames = []string{
	_PlayerTypeName[0:7],
	_PlayerTypeName[7:14],
}

// PlayerTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PlayerTypeString(s string) (PlayerType, error) {
	if val, ok := _PlayerTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PlayerTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PlayerType values", s)
}

// PlayerTypeValues returns all values of the enum
func PlayerTypeValues() []PlayerType {
	return _PlayerTypeValues
}

// PlayerTypeStrings returns a slice of all String values of the enum
func PlayerTypeStrings() []string {
	strs := make([]string, len(_PlayerTypeNames))
	copy(strs, _PlayerTypeNames)
	return strs
}

// IsAPlayerType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PlayerType) IsAPlayerType() bool {
	for _, v := range _PlayerTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for PlayerType
func (i PlayerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PlayerType
func (i *PlayerType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PlayerType should be a string, got %s", data)
	}

	var err error
	*i, err = PlayerTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for PlayerType
func (i PlayerType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for PlayerType
func (i *PlayerType) UnmarshalText(text []byte) error {
	var err error
	*i, err = PlayerTypeString(string(text))
	return err
}

func (i PlayerType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *PlayerType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of PlayerType: %[1]T(%[1]v)", value)
	}

	val, err := PlayerTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
