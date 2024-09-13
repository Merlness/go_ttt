package main

import (
)

var valueMax = -100000
var valueMin = 100000

func Maximize(board []interface{}, maxToken string, minToken string, depth int) (int, int) {
    var value int
    bestValue := valueMax
    var bestAction int

    depth += 1
    gameOver, _ := GameOver(board, maxToken, minToken)

    if gameOver {
    	value = findValue(board, depth, maxToken, minToken)
    	return value, bestAction
    }

    for _, action := range AvailableMoves(board) {

        var newBoard []interface{}
    	for _, value := range board {
            newBoard = append(newBoard, value)
    	}

        UpdateBoard(newBoard, action, maxToken)
        value, _ = minimize(newBoard, maxToken, minToken, depth)

          if value > bestValue {
                bestValue = value
                bestAction = action
            }
    }
    return bestValue, bestAction
}


func minimize(board []interface{}, maxToken string, minToken string, depth int) (int, int) {
    var value int
    bestValue := valueMin
    var bestAction int

    depth += 1
    gameOver, _ := GameOver(board, maxToken, minToken)

    if gameOver {
    	value = findValue(board, depth, maxToken, minToken)
    	return value, bestAction
    }

    for _, action := range AvailableMoves(board) {

        var newBoard []interface{}
    	for _, value := range board {
            newBoard = append(newBoard, value)
    	}
        UpdateBoard(newBoard, action, minToken)
        value, _ = Maximize(newBoard, maxToken, minToken, depth)

          if value < bestValue {
                bestValue = value
                bestAction = action
            }
    }
    return bestValue, bestAction
}

func computerWins(board []interface{}, token string) bool {
    return CheckWinner(board, token)
}

func opponentWins(board []interface{}, token string) bool {
    return CheckWinner(board, token)
}

func findValue(board []interface{}, depth int, maxToken, minToken string) int {
    if computerWins(board, maxToken) {
        return 12 / depth
    } else if opponentWins(board, minToken) {
        return -12 / depth
    }
    return 0
}
