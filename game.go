package main

import (
	"fmt"
	"os"
	"strings"

)

func InitializeGameStruct(gameStruct Game, player1Token string, player1Kind string, player2Kind string) Game {
	gameStruct = assignTokens(gameStruct, player1Token)
	gameStruct = assignPlayerKinds(gameStruct, player1Kind, player2Kind)

	return gameStruct
}

func assignTokens(gameStruct Game, choice string) Game {
	if choice == "X" {
		gameStruct.Token1 = "X"
		gameStruct.Token2 = "O"
	} else {
		gameStruct.Token1 = "O"
		gameStruct.Token2 = "X"
	}
	return gameStruct
}

func assignPlayerKinds(gameStruct Game, player1 string, player2 string) Game {
	gameStruct.Player1 = strings.ToLower(player1)
	gameStruct.Player2 = strings.ToLower(player2)
	return gameStruct
}

func UpdateBoard(board []interface{}, pos int, mark string) []interface{} {
	board[pos-1] = mark
	return board
}

func GetMove(board []interface{}, currentPlayer string, gameStruct Game) int {

	if currentPlayer == "human" {
		return GetHumanMove(board)
	}

	return GetAIMove(board, gameStruct)
}

func GetHumanMove(board []interface{}) int {
	for {
		pos := GetUserInput(os.Stdin)

		if IsValidMove(board, pos) {
			return pos
		}
		fmt.Println("Invalid move, try again.")
	}
}

func GetAIMove(board []interface{}, gameStruct Game) int {
    maxToken := CurrentPlayerToken(board, gameStruct)
    minToken := OpponentPlayerToken(board, gameStruct)

	if areAllInts(board) {
		fmt.Println("AI's move")
    	return 1
    	}

    _, bestAction := Maximize(board, maxToken, minToken, 1)
    fmt.Println("AI's move")

    return bestAction
}

func areAllInts(board []interface{}) bool {
	for _, value := range board {
		_, isInt := value.(int)
        if !isInt {
        	return false
        }
	}
	return true
}

func isPlayerOnesTurn(board []interface{}) bool {
	return countAvailableMoves(board) % 2 != 0
}

func CurrentPlayerToken(board []interface{}, gameStruct Game) string {
	return getCurrentValue(board, gameStruct.Token1, gameStruct.Token2)
}

func OpponentPlayerToken(board []interface{}, gameStruct Game) string {
	return getCurrentValue(board, gameStruct.Token2, gameStruct.Token1)
}

func CurrentPlayer(board []interface{}, gameStruct Game) string {
	return getCurrentValue(board, gameStruct.Player1, gameStruct.Player2)
}

func getCurrentValue(board []interface{}, player1Val, player2Val string) string {
	if isPlayerOnesTurn(board) {
		return player1Val
	}
	return player2Val
}
