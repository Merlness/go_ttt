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

		AssertEqual(t, pos, expected)
	})
}

func TestIsValidMove(t *testing.T) {
	board := []interface{}{1, 2, 3, 4, "X", 6, 7, 8, 9}

	t.Run("valid move", func(t *testing.T) {
		pos := 3
		if !IsValidMove(board, pos) {
			t.Errorf("expected %d valid", pos)
		}
	})

	t.Run("invalid move", func(t *testing.T) {
		pos := 5
		if IsValidMove(board, pos) {
			t.Errorf("expected %d ", pos)
		}
	})

	t.Run("invalid move", func(t *testing.T) {
		pos := 10
		if IsValidMove(board, pos) {
			t.Errorf("expected %d invalid", pos)
		}
	})
}

func TestAskPlayer1Token(t *testing.T) {
	t.Run("Player 1 X", func(t *testing.T) {
		resetGameStruct()

		input := strings.NewReader("X\n")
		AskPlayer1Token(input)

		if gameStruct.Token1 != "X" || gameStruct.Token2 != "O" {
			t.Errorf("expected Player 1: X and Player 2: O, got Player 1: %s, Player 2: %s", gameStruct.Token1, gameStruct.Token2)
		}
	})

	t.Run("Player 1 O", func(t *testing.T) {
		resetGameStruct()

		input := strings.NewReader("O\n")
		player1Token := AskPlayer1Token(input)
        gameStruct2 := InitializeGameStruct(gameStruct, player1Token, "human", "human")

		if gameStruct2.Token1 != "O" || gameStruct2.Token2 != "X" {
			t.Errorf("expected Player 1: O and Player 2: X, got Player 1: %s, Player 2: %s", gameStruct2.Token1, gameStruct2.Token2)
		}
	})

	t.Run("Player 1 x", func(t *testing.T) {
    		resetGameStruct()

    		input := strings.NewReader("x\n")
    		AskPlayer1Token(input)

    		if gameStruct.Token1 != "X" || gameStruct.Token2 != "O" {
    			t.Errorf("expected Player 1: X and Player 2: O, got Player 1: %s, Player 2: %s", gameStruct.Token1, gameStruct.Token2)
    		}
    	})

    	t.Run("Player 1 o", func(t *testing.T) {
    		resetGameStruct()

    		input := strings.NewReader("o\n")
    		player1Token := AskPlayer1Token(input)
            gameStruct2 := InitializeGameStruct(gameStruct, player1Token, "human", "human")

            if gameStruct2.Token1 != "O" || gameStruct2.Token2 != "X" {
            	t.Errorf("expected Player 1: O and Player 2:  X, got Player 1: %s, Player 2: %s", gameStruct2.Token1, gameStruct2.Token2)
            }
    	})
}

func TestAskPlayer1Kind(t *testing.T) {
	resetGameStruct()

	t.Run("Player 1 human", func(t *testing.T) {
		input := strings.NewReader("human\n")
        AskPlayerKind(input, "Player 1")

        AssertEqual(t, gameStruct.Player1, "human")
	})

	t.Run("Player 1 ai", func(t *testing.T) {
		resetGameStruct()
		input := strings.NewReader("ai\n")
	    player1Kind  := AskPlayerKind(input, "Player 1")
	    gameStruct2 := InitializeGameStruct(gameStruct, "X", player1Kind, "ai")

        AssertEqual(t, gameStruct2.Player1, "ai")
	})
}

func TestAskPlayer2Kind(t *testing.T) {
	resetGameStruct()

	t.Run("Player 2 human", func(t *testing.T) {
		input := strings.NewReader("human\n")
        AskPlayerKind(input, "Player 2")

        AssertEqual(t, gameStruct.Player2, "human")
	})

	t.Run("Player 2 ai", func(t *testing.T) {
		resetGameStruct()
		input := strings.NewReader("ai\n")
        player2Kind  := AskPlayerKind(input, "Player 2")

	    gameStruct2 := InitializeGameStruct(gameStruct, "X", "human",player2Kind)

        AssertEqual(t, gameStruct2.Player2, "ai")
	})
}
