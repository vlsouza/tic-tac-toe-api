package service

import "errors"

var (
	ErrNotFound       = errors.New("match not found")
	ErrMatchNotRunnig = errors.New("the match is current not Running. Cannot update the board")
)
