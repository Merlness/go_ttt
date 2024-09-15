package main

import (
    "testing"
)

func resetGameStruct() {
	gameStruct = Game{
		Player1: "human",
		Player2: "human",
		Token1:  "X",
		Token2:  "O",
	}
}

func TestGreeting(t *testing.T) {

    got := Greeting()
	want := "Welcome to Tic Tac Toe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGlobalGameStruct(t *testing.T){

	t.Run("check initial gameStruct values", func(t *testing.T) {
    	expected := Game{
    		Player1: "human",
    		Player2: "human",
    		Token1:  "X",
    		Token2:  "O",
    	}
           resetGameStruct()
           AssertEqual(t, gameStruct, expected )
    })
}

func TestInitializeGameStruct(t *testing.T) {
    t.Run("Initializes game struct correctly", func(t *testing.T) {
        gameStruct := Game{}
        player1Token := "X"
        player1Kind := "human"
        player2Kind := "ai"

        gameStruct = InitializeGameStruct(gameStruct, player1Token, player1Kind, player2Kind)

        expectedGameStruct := Game{
            Player1: "human",
            Player2: "ai",
            Token1:  "X",
            Token2:  "O",
        }
        AssertEqual(t, gameStruct, expectedGameStruct)
    })
}
