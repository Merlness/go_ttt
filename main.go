package main

import (
	"fmt"
	"os"
)

type Game struct {
	Player1 string
	Player2 string
	Token1  string
	Token2  string
}

var gameStruct = Game{
	Player1: "",
	Player2: "",
	Token1:  "",
	Token2:  "",
}

func main() {
	board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	gameStruct = startGame(board)
    playGame(board, gameStruct)
}

func startGame(board []interface{}) Game{
    fmt.Println(Greeting())

    player1Token := AskPlayer1Token(os.Stdin)
    player1Kind  := AskPlayerKind(os.Stdin, "Player 1")
    player2Kind  := AskPlayerKind(os.Stdin, "Player 2")

    fmt.Println(Display(board))

    return InitializeGameStruct(gameStruct, player1Token, player1Kind, player2Kind)
}

func playGame(board []interface{}, gameStruct Game) {
	for {
		if isGameOver(board, gameStruct) {
			break
		}
		board = playTurn(board, gameStruct)
	}
}

func playTurn(board []interface{}, gameStruct Game) []interface{} {
	currentToken := CurrentPlayerToken(board, gameStruct)
	currentPlayer := CurrentPlayer(board, gameStruct)

	pos := GetMove(board, currentPlayer, gameStruct)
	board = UpdateBoard(board, pos, currentToken)

	fmt.Println(Display(board))
	return board
}

func isGameOver(board []interface{}, gameStruct Game) bool {
	over, message := GameOver(board, gameStruct.Token1, gameStruct.Token2)
	if over {
		fmt.Println(message)
	}
	return over
}
