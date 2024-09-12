package main

import (
	"testing"
	"strings"
// 	"fmt"
)

func TestGetUserInput(t *testing.T) {
	t.Run("gets valid input", func(t *testing.T) {
		input := strings.NewReader("3\n")
		pos := GetUserInput(input)
		expected := 3

		if pos != expected {
			t.Errorf("expected %d, got %d", expected, pos)
		}
	})
}

func TestIsValidMove(t *testing.T) {
	board := []interface{}{1, 2, 3, 4, "X", 6, 7, 8, 9}

	t.Run("valid move", func(t *testing.T) {
		pos := 3
		if !IsValidMove(board, pos) {
			t.Errorf("expected move at position %d to be valid", pos)
		}
	})

	t.Run("invalid move (already taken)", func(t *testing.T) {
		pos := 5
		if IsValidMove(board, pos) {
			t.Errorf("expected move at position %d to be invalid", pos)
		}
	})

	t.Run("invalid move (out of range)", func(t *testing.T) {
		pos := 10
		if IsValidMove(board, pos) {
			t.Errorf("expected move at position %d to be invalid", pos)
		}
	})
}

func TestUpdateBoard(t *testing.T) {
	board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("update board with 'X'", func(t *testing.T) {
		pos := 3
		updatedBoard := UpdateBoard(board, pos, "X")
		expectedBoard := []interface{}{1, 2, "X", 4, 5, 6, 7, 8, 9}

		for i, v := range updatedBoard {
			if v != expectedBoard[i] {
				t.Errorf("expected %v, got %v", expectedBoard, updatedBoard)
			}
		}
	})
}

func TestAskPlayer1Token(t *testing.T) {
	t.Run("Player 1 chooses X", func(t *testing.T) {
		resetGameMap()

		input := strings.NewReader("X\n")
		AskPlayer1Token(input)

		if gameMap.Token1 != "X" || gameMap.Token2 != "O" {
			t.Errorf("expected Player 1 to be X and Player 2 to be O, got Player 1: %s, Player 2: %s", gameMap.Token1, gameMap.Token2)
		}
	})

	t.Run("Player 1 chooses O", func(t *testing.T) {
		resetGameMap()

		input := strings.NewReader("O\n")
		AskPlayer1Token(input)

		if gameMap.Token1 != "O" || gameMap.Token2 != "X" {
			t.Errorf("expected Player 1 to be O and Player 2 to be X, got Player 1: %s, Player 2: %s", gameMap.Token1, gameMap.Token2)
		}
	})

	t.Run("Player 1 chooses x", func(t *testing.T) {
    		resetGameMap()

    		input := strings.NewReader("x\n")
    		AskPlayer1Token(input)

    		if gameMap.Token1 != "X" || gameMap.Token2 != "O" {
    			t.Errorf("expected Player 1 to be X and Player 2 to be O, got Player 1: %s, Player 2: %s", gameMap.Token1, gameMap.Token2)
    		}
    	})

    	t.Run("Player 1 chooses o", func(t *testing.T) {
    		resetGameMap()

    		input := strings.NewReader("o\n")
    		AskPlayer1Token(input)

    		if gameMap.Token1 != "O" || gameMap.Token2 != "X" {
    			t.Errorf("expected Player 1 to be O and Player 2 to be X, got Player 1: %s, Player 2: %s", gameMap.Token1, gameMap.Token2)
    		}
    	})
}
