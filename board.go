package main

import "fmt"

func Display(grid []interface{}) string {
	var result string
	for i, val := range grid {
		if i > 0 && i%3 == 0 {
			result += "\n"
		} else if i%3 != 0 {
			result += " | "
		}
		result += fmt.Sprintf("%v", val)
	}
	return result
}

func countAvailableMoves(board []interface{}) int {
	intCount := 0
	for _, v := range board {
		if _, isInt := v.(int); isInt {
			intCount++
		}
	}
	return intCount
}

func isPlayerOnesTurn(board []interface{}) bool {
	return countAvailableMoves(board) % 2 != 0
}

func CurrentPlayerToken(board []interface{}) string {

	if isPlayerOnesTurn(board) {
		return gameMap.Token1
	}
	return gameMap.Token2
}
