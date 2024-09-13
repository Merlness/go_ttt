package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Greeting() string {
 return "Welcome to Tic Tac Toe"
}

func assignTokens(choice string) {
	if choice == "X" {
		gameMap.Token1 = "X"
		gameMap.Token2 = "O"
	} else {
		gameMap.Token1 = "O"
		gameMap.Token2 = "X"
	}
}

func assignPlayer1Type(choice string) {
	if choice == "HUMAN" {
		gameMap.Player1 = "human"
	} else {
		gameMap.Player1 = "ai"
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
		return isBetweenOneAndNine(pos) && isAvailable(board, pos)
}

func isAvailable(board []interface{}, pos int) bool {
    valueAtPosition := board[pos-1]
    _, isInt := valueAtPosition.(int)

    return isInt
}

func isBetweenOneAndNine(pos int) bool {
	return pos >= 1 && pos <= 9
}

func AskPlayer1Token(input io.Reader) {
	askUserInput(input, "Would you like Player 1 to be X or O? ", map[string]func(string){
		"X": assignTokens,
		"O": assignTokens,
	})
}

func AskPlayer1Type(input io.Reader) {
	askUserInput(input, "Would you like Player 1 to be human or ai? ", map[string]func(string){
		"HUMAN": assignPlayer1Type,
		"AI":    assignPlayer1Type,
	})
}

func askUserInput(input io.Reader, prompt string, validInputs map[string]func(string)) {
	reader := bufio.NewReader(input)

	for {
		fmt.Print(prompt)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.ToUpper(userInput))

        assignFunc, ok := validInputs[userInput]
		if  ok {
			assignFunc(userInput)
			break
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
