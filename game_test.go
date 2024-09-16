package main

import (
	"testing"
)


func TestUpdateBoard(t *testing.T) {
	board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("update board with 'X'", func(t *testing.T) {
		pos := 3
		updatedBoard := UpdateBoard(board, pos, "X")
		expectedBoard := []interface{}{1, 2, "X", 4, 5, 6, 7, 8, 9}

		for i, v := range updatedBoard {
		    AssertEqual(t, v, expectedBoard[i])
		}
	})
}

func TestAIMoves(t *testing.T) {

	t.Run("AI makes first move on empty board", func(t *testing.T) {
    	gameStruct.Player1 = "ai"
    	gameStruct.Player2 = "human"
    	gameStruct.Token1 = "X"
    	gameStruct.Token2 = "O"

    	board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
    	expectedMove := 1

    	pos := GetMove(board, "ai", gameStruct)

		AssertEqual(t, pos, expectedMove)
    })

	t.Run("AI makes move at position 5", func(t *testing.T) {
		gameStruct.Player1 = "human"
		gameStruct.Player2 = "ai"
		gameStruct.Token1 = "X"
		gameStruct.Token2 = "O"

		board := []interface{}{"X", 2, 3, 4, 5, 6, 7, 8, 9}
		expectedMove := 5

		pos := GetMove(board, "ai", gameStruct)

		AssertEqual(t, pos, expectedMove)
	})

	t.Run("AI makes move at position 4", func(t *testing.T) {
		gameStruct.Player1 = "ai"
		gameStruct.Player2 = "human"
		gameStruct.Token1 = "X"
		gameStruct.Token2 = "O"

		board := []interface{}{"X", "O", 3, 4, 5, 6, 7, 8, 9}
		expectedMove := 4
		pos := GetMove(board, "ai", gameStruct)

		AssertEqual(t, pos, expectedMove)
	})

	t.Run("AI blocks winning move", func(t *testing.T) {
		gameStruct.Player1 = "human"
		gameStruct.Player2 = "ai"
		gameStruct.Token1 = "X"
		gameStruct.Token2 = "O"

		board := []interface{}{"X", "X", 3, 4, 5, 6, 7, 8, 9}
		expectedMove := 3
		pos := GetMove(board, "ai", gameStruct)

		AssertEqual(t, pos, expectedMove)
	})

	t.Run("AI takes winning move", func(t *testing.T) {
		gameStruct.Player1 = "ai"
		gameStruct.Player2 = "human"
		gameStruct.Token1 = "X"
		gameStruct.Token2 = "O"

		board := []interface{}{"X", "X", 3, "O", 5, 6, "O", 8, 9}
		expectedMove := 3
		pos := GetMove(board, "ai", gameStruct)

		AssertEqual(t, pos, expectedMove)
	})
}

func TestIsPlayerOnesTurn(t *testing.T) {
    t.Run("Returns true when it's Player 1's turn", func(t *testing.T) {
        board := []interface{}{"X", "O", "X", 4, "O", 6, 7, 8, 9}

		AssertTrue(t, isPlayerOnesTurn(board), "Expected Player 1's turn")
    })

    t.Run("Returns false when it's Player 2's turn", func(t *testing.T) {
        board := []interface{}{"X", "O", "X", "O", "X", 6, 7, 8, 9}
		AssertTrue(t, !isPlayerOnesTurn(board), "Expected Player 2's turn")
    })
}

func TestAreAllInts(t *testing.T) {
    t.Run("Returns true when all board positions are integers", func(t *testing.T) {
        board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
		AssertTrue(t, areAllInts(board), "Expected areAllInts to return true")
    })

    t.Run("Returns false when board positions include tokens", func(t *testing.T) {
        board := []interface{}{"X", 2, 3, 4, "O", 6, 7, 8, 9}
		AssertTrue(t, !areAllInts(board), "Expected areAllInts to return false")
    })
}

func TestCurrentPlayerFunctions(t *testing.T) {
    gameStruct := Game{Player1: "human", Player2: "ai", Token1: "X", Token2: "O"}

    t.Run("CurrentPlayer returns correct player", func(t *testing.T) {
        board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
        expectedPlayer := "human"
        currentPlayer := CurrentPlayer(board, gameStruct)

		AssertEqual(t, currentPlayer, expectedPlayer)
    })

    t.Run("OpponentPlayerToken returns correct token", func(t *testing.T) {
        board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
        expectedToken := "O"
        opponentToken := OpponentPlayerToken(board, gameStruct)

		AssertEqual(t, expectedToken, opponentToken)
    })
}
