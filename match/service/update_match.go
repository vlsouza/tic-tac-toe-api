package service

import (
	"context"
	"fmt"
	"main/match/repository"
	"strings"
)

func (svc Service) PlaceMove(ctx context.Context, request MoveRequest) error {
	newMatchState, err := svc.GetNextState(ctx, request)
	_, err = svc.repo.Update(ctx, newMatchState)
	if err != nil {
		return err
	}
	return nil
}

func (svc Service) GetNextState(
	ctx context.Context,
	request MoveRequest,
) (repository.Match, error) {
	currentMatchState, err := svc.repo.GetByID(ctx, request.MatchID)
	if err != nil {
		return repository.Match{}, err
	}

	newBoard := getBoard(currentMatchState.Board, request.Player, request.Row, request.Col)
	newStatus := getGameStatus(currentMatchState.Board, request.Player)
	newCurrentPlayer, newNextPlayerTurn := getPlayers(currentMatchState.CurrentPlayerTurn)

	return repository.Match{
		ID:                currentMatchState.ID,
		Status:            newStatus,
		Board:             newBoard,
		CurrentPlayerTurn: newCurrentPlayer,
		NextPlayerTurn:    newNextPlayerTurn,
		LastMoveXY:        fmt.Sprintf("%v,%v", request.Row, request.Col),
	}, nil
}

// TODO optimize updateBoard and inside functions
// Updates the tic-tac-toe board and checks game status
func getBoard(board string, player, row, col int8) string {
	rows := strings.Split(board, ",")
	for i, rowContent := range rows {
		cells := strings.Split(rowContent, "")
		if int8(i) == row {
			cells[col] = fmt.Sprintf("%d", player)
			rows[i] = strings.Join(cells, "")
		}
	}
	return strings.Join(rows, ",")
}

// Get the status of the game
func getGameStatus(board string, player int8) string {
	rows := strings.Split(board, ",")
	winningPlayer := fmt.Sprintf("%d%d%d", player, player, player)

	// Check rows and columns for win
	for i := 0; i < 3; i++ {
		if strings.Join(strings.Split(rows[i], ""), "") == winningPlayer {
			return winStatus(player)
		}
		if strings.Join([]string{string(rows[0][i]), string(rows[1][i]), string(rows[2][i])}, "") == winningPlayer {
			return winStatus(player)
		}
	}

	// Check diagonals for win
	if strings.Join([]string{string(rows[0][0]), string(rows[1][1]), string(rows[2][2])}, "") == winningPlayer ||
		strings.Join([]string{string(rows[0][2]), string(rows[1][1]), string(rows[2][0])}, "") == winningPlayer {
		return winStatus(player)
	}

	// Check for tie
	if !strings.Contains(strings.Join(rows, ""), "0") {
		return "TIE"
	}

	return "RUNNING"
}

func getPlayers(currentPlayer int8) (int8, int8) {
	if currentPlayer == 1 {
		return 2, 1
	} else {
		return 1, 2
	}
}

// Returns the win status based on the player
func winStatus(player int8) string {
	if player == 1 {
		return "PLAYER1WON"
	}
	return "PLAYER2WON"
}
