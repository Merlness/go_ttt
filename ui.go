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

func AskPlayer1Token(input io.Reader) string {
	return askPlayerInput(input, "Would you like Player 1 to be 'X' or 'O'? ", []string{"x", "o"})
}

func AskPlayerKind(input io.Reader, player string) string {
	question := fmt.Sprintf("Would you like %s to be 'human' or 'ai'? ", player)
	return askPlayerInput(input, question, []string{"human", "ai"})
}


func askPlayerInput(input io.Reader, prompt string, validChoices []string) string {
	reader := bufio.NewReader(input)

	for {
		fmt.Print(prompt)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(strings.ToLower(userInput))

		for _, choice := range validChoices {
			if userInput == choice {
				return userInput
			}
		}
		fmt.Println("Invalid input. Please try again.")
	}
}
