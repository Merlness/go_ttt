package main

import (
        "testing"
        "reflect"
        )

func resetGameMap() {
	gameMap = GameMap{
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

func TestGlobalGameMap(t *testing.T){

		t.Run("check initial gameMap values", func(t *testing.T) {
    		expected := GameMap{
    			Player1: "human",
    			Player2: "human",
    			Token1:  "X",
    			Token2:  "O",
    		}

    		if !reflect.DeepEqual(gameMap, expected) {
    			t.Errorf("expected %v, got %v", expected, gameMap)
    		}
    	})
}
