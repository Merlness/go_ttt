package main

import (
        "testing"
        "reflect"
        "strings"
        )

func resetGameMap() {
	gameMap = GameMap{
		Player1: "human",
		Player2: "human",
		Token1:  "X",
		Token2:  "O",
		Moves:   []int{},
	}
}

func TestGreeting(t *testing.T) {

    got := Greeting()
	want := "Welcome to Tic Tac Toe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGlobalGameMap(t *testing.T){

		t.Run("check initial gameMap values", func(t *testing.T) {
    		expected := GameMap{
    			Player1: "human",
    			Player2: "human",
    			Token1:  "X",
    			Token2:  "O",
    			Moves:   []int{},
    		}

    		if !reflect.DeepEqual(gameMap, expected) {
    			t.Errorf("expected %v, got %v", expected, gameMap)
    		}
    	})
}

func TestValidMoveUpdatesGameMap(t *testing.T) {

	resetGameMap()
	input := strings.NewReader("3\n")
	pos := GetUserInput(input)
	UpdateMoves(pos)

	expectedMoves := []int{3}
	if !reflect.DeepEqual(gameMap.Moves, expectedMoves) {
		t.Errorf("expected moves %v, got %v", expectedMoves, gameMap.Moves)
	}
}
