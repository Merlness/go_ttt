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
		_, isInt := v.(int)
		if  isInt {
			intCount++
		}
	}
	return intCount
}

func GameOver(board []interface{}, token1 string, token2 string) (bool, string) {
	if CheckWinner(board, token1) {
		return true, fmt.Sprintf("%s wins!", token1)
	} else if CheckWinner(board, token2){
		return true, fmt.Sprintf("%s wins!", token2)
	} else if !HasEmptySpaces(board) {
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

func HasEmptySpaces(board []interface{}) bool {
	for _, value := range board {
		_, isInt := value.(int)
		if  isInt {
			return true
		}
	}
	return false
}

func AvailableMoves(board []interface{}) []int {
	var available []int
	for _, value := range board {
		move, isInt := value.(int)
		if  isInt {
			available = append(available, move)
		}
	}
	return available
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

func CopyBoard(board []interface{}) []interface{} {
    newBoard := make([]interface{}, len(board))
    copy(newBoard, board)
    return newBoard
}
