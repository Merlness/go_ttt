package main

import (
	"fmt"
	"os"
)


func UpdateBoard(board []interface{}, pos int, mark string) []interface{} {
	board[pos-1] = mark
	return board
}

func GetMove(board []interface{}) int {
	currentPlayer := CurrentPlayer(board)

	if currentPlayer == "human" {
		return GetHumanMove(board)
	}

	return GetAIMove(board)
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

func GetAIMove(board []interface{}) int {
    maxToken := CurrentPlayerToken(board)
    minToken := OpponentPlayerToken(board)

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
