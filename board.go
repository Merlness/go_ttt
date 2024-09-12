package main

import (
    "fmt"
	"math"
)

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

func GameOver(board []interface{}, token string) (bool, string) {
	if CheckWinner(board, token) {
		return true, fmt.Sprintf("%s wins!", token)
	}

	if !hasEmptySpaces(board) {
		return true, "Tie!"
	}

	return false, ""
}

func CheckWinner(board []interface{}, token string) bool {
	size := boardSize(board)

	return checkLines(board, token, rows(size)) ||
		checkLines(board, token, columns(size)) ||
		checkLines(board, token, diagonals(size))
}

func hasEmptySpaces(board []interface{}) bool {
	for _, v := range board {
		if _, isInt := v.(int); isInt {
			return true
		}
	}
	return false
}

func checkLines(board []interface{}, token string, lines [][]int) bool {
	for _, line := range lines {
		if checkLine(board, token, line) {
			return true
		}
	}
	return false
}

func checkLine(board []interface{}, token string, indices []int) bool {
	for _, index := range indices {
		if board[index] != token {
			return false
		}
	}
	return true
}

func rows(size int) [][]int {
	rowIndices := make([][]int, size)
	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			row[j] = i*size + j
		}
		rowIndices[i] = row
	}
	return rowIndices
}

func columns(size int) [][]int {
	colIndices := make([][]int, size)
	for i := 0; i < size; i++ {
		column := make([]int, size)
		for j := 0; j < size; j++ {
			column[j] = i + j*size
		}
		colIndices[i] = column
	}
	return colIndices
}

func diagonals(size int) [][]int {
	return [][]int{
		backDiagonal(size),
		frontDiagonal(size),
	}
}

func backDiagonal(size int) []int {
	indices := make([]int, size)
	for i := 0; i < size; i++ {
		indices[i] = i * (size + 1)
	}
	return indices
}

func frontDiagonal(size int) []int {
	indices := make([]int, size)
	for i := 0; i < size; i++ {
		indices[i] = (i + 1) * (size - 1)
	}
	return indices
}

func boardSize(board []interface{}) int {
	return int(math.Sqrt(float64(len(board))))
}
