package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"os"
)

func Greeting() string {
 return "Welcome to Tic Tac Toe"
}

func GetValidMove(board []interface{}) int {
	for {
		pos := GetUserInput(os.Stdin)

		if IsValidMove(board, pos) {
			return pos
		}
		fmt.Println("Invalid move, try again.")
	}
}

func GetUserInput(input io.Reader) int {
	reader := bufio.NewReader(input)
	fmt.Print("Place the next move: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	pos, _ := strconv.Atoi(userInput)

	return pos
}


func IsValidMove(board []interface{}, pos int) bool {
		return isBetweenOneAndNine(pos) && !isTakenByToken(board, pos)
}

func UpdateBoard(board []interface{}, pos int, mark string) []interface{} {
	board[pos-1] = mark
	return board
}

func isTakenByToken(board []interface{}, pos int) bool {
    valueAtPosition := board[pos-1]
    _, isInt := valueAtPosition.(int)

    return !isInt
}

func isBetweenOneAndNine(pos int) bool {
	return pos >= 1 && pos <= 9
}
