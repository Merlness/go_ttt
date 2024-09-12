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

func AskPlayer1Token(input io.Reader) {
	reader := bufio.NewReader(input)

	for {
		fmt.Print("Would you like Player 1 to be X or O? ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.ToUpper(userInput))

		if userInput == "X" {
        	gameMap.Token1 = "X"
        	gameMap.Token2 = "O"
        	break
        } else if userInput == "O" {
			gameMap.Token1 = "O"
			gameMap.Token2 = "X"
			break
		} else {
			fmt.Println("Invalid input. Please enter 'X' or 'O'.")
		}
	}
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
