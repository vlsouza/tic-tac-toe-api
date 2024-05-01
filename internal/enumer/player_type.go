package enumer

//go:generate enumer -json -text -sql -type PlayerType -trimprefix Player -transform snake-upper

// Status represents the available players in a match
type PlayerType int

const (
	PLAYER1 PlayerType = iota
	PLAYER2
)
