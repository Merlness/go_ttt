package main

import (
	"fmt"
)

type GameMap struct {
	Player1 string
	Player2 string
	Token1  string
	Token2  string
}

var gameMap = GameMap{
	Player1: "human",
	Player2: "human",
	Token1:  "X",
	Token2:  "O",
}



func main() {
	board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

    fmt.Println(Greeting())
    for {
		pos := GetValidMove(board)

		currentToken := CurrentPlayerToken(board)

		board = UpdateBoard(board, pos, currentToken)
		fmt.Println(Display(board))

	}
}
