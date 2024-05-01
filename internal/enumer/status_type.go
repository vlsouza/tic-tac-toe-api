package enumer

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
