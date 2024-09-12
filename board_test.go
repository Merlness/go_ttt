package main

import "testing"


func checkDisplay(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDisplay(t *testing.T) {
	t.Run("displays the default board", func(t *testing.T) {
		grid := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Display(grid)
		want := "1 | 2 | 3\n4 | 5 | 6\n7 | 8 | 9"
		checkDisplay(t, got, want)
	})

	t.Run("displays a mix of numbers and tokens", func(t *testing.T) {
		grid := []interface{}{"X", "O", 3, "X", 5, "O", 7, 8, 9}
		got := Display(grid)
		want := "X | O | 3\nX | 5 | O\n7 | 8 | 9"
		checkDisplay(t, got, want)
	})

	t.Run("displays all tokens", func(t *testing.T) {
		grid := []interface{}{"X", "O", "X", "X", "O", "X", "O", "X", "O"}
		got := Display(grid)
		want := "X | O | X\nX | O | X\nO | X | O"
		checkDisplay(t, got, want)
	})
}

func TestCurrentPlayerToken(t *testing.T) {
	board := []interface{}{1, "X", 3, "O", 5, 6, 7, 8, 9}

	t.Run("Player 1's turn when odd number of integers left", func(t *testing.T) {
		got := CurrentPlayerToken(board)
		want := gameMap.Token1

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	board = []interface{}{1, "X", "X", "O", 5, 6, 7, 8, 9}

	t.Run("Player 2's turn when even number of integers left", func(t *testing.T) {
		got := CurrentPlayerToken(board)
		want := gameMap.Token2

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}
