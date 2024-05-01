package service

import (
	"context"
	"errors"
	"fmt"

	"main/internal/enumer"
	"main/internal/repository"
	"strings"

	"github.com/google/uuid"
)

func (svc Service) Start(ctx context.Context, matchID uuid.UUID) error {
	matchState, err := svc.repo.GetByID(ctx, matchID)
	if err != nil {
		return err
	}

	matchState.Status = enumer.RUNNING
	_, err = svc.repo.Update(ctx, matchState)
	if err != nil {
		return err
	}

	return nil
}

func (svc Service) PlaceMove(ctx context.Context, request MoveRequest) (GetStateResponse, error) {
	newMatchState, err := svc.GetNextState(ctx, request)
	if err != nil {
		return GetStateResponse{}, err
	}
	_, err = svc.repo.Update(ctx, newMatchState)
	if err != nil {
		return GetStateResponse{}, err
	}

	return NewGetStateResponse(newMatchState)
}

func (svc Service) GetNextState(
	ctx context.Context,
	request MoveRequest,
) (repository.Match, error) {
	currentMatchState, err := svc.repo.GetByID(ctx, request.MatchID)
	if err != nil {
		return repository.Match{}, err
	}

	if currentMatchState.Status != enumer.RUNNING {
		return repository.Match{}, errors.New("the match is current not Running. Cannot update the board")
	}

	newBoard := getBoard(currentMatchState.Board, currentMatchState.CurrentPlayerTurn, request.Row, request.Col)
	newStatus := getGameStatus(newBoard, currentMatchState.CurrentPlayerTurn)
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
func getBoard(
	board string,
	player enumer.PlayerType,
	row, col int8,
) string {
	//fix bug generating new board
	rows := strings.Split(board, ",")
	for i := 0; i < 3; i++ {
		cells := strings.Split(rows[i], "")
		if int8(i) == row-1 {
			cells[col-1] = strings.Replace(player.String(), "PLAYER", "", -1)
			rows[i] = strings.Join(cells, "")
		}
	}

	return strings.Join(rows, ",")
}

// Get the status of the game
func getGameStatus(board string, player enumer.PlayerType) enumer.StatusType {
	rows := strings.Split(board, ",")
	//move to function or improve 'player' type
	playerNumber := strings.Replace(player.String(), "PLAYER", "", -1)
	winningPlayer := fmt.Sprintf("%s%s%s", playerNumber, playerNumber, playerNumber)

	// Check rows and columns for win
	for i := 0; i < 3; i++ {
		if rows[i] == winningPlayer {
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
		return enumer.DRAW
	}

	return enumer.RUNNING
}

func getPlayers(currentPlayer enumer.PlayerType) (enumer.PlayerType, enumer.PlayerType) {
	if currentPlayer == enumer.PLAYER1 {
		return enumer.PLAYER2, enumer.PLAYER1
	}
	return enumer.PLAYER1, enumer.PLAYER2
}

// Returns the win status based on the player
func winStatus(player enumer.PlayerType) enumer.StatusType {
	if player == enumer.PLAYER1 {
		return enumer.PLAYER1WON
	}
	return enumer.PLAYER2WON
}
