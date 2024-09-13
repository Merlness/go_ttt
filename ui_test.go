package main

import (
	"testing"
	"strings"
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
			t.Errorf("expected %d to be valid", pos)
		}
	})

	t.Run("invalid move (already taken)", func(t *testing.T) {
		pos := 5
		if IsValidMove(board, pos) {
			t.Errorf("expected %d to be invalid", pos)
		}
	})

	t.Run("invalid move (out of range)", func(t *testing.T) {
		pos := 10
		if IsValidMove(board, pos) {
			t.Errorf("expected %d to be invalid", pos)
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

func TestAskPlayer1Type(t *testing.T) {
	resetGameMap()

	t.Run("Player 1 chooses human", func(t *testing.T) {
		input := strings.NewReader("human\n")
		AskPlayer1Type(input)

		if gameMap.Player1 != "human" {
			t.Errorf("expected Player 1 to be human, got %s", gameMap.Player1)
		}
	})

	t.Run("Player 1 chooses ai", func(t *testing.T) {
		resetGameMap()
		input := strings.NewReader("ai\n")
		AskPlayer1Type(input)

		if gameMap.Player1 != "ai" {
			t.Errorf("expected Player 1 to be ai, got %s", gameMap.Player1)
		}
	})
}
