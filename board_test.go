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

func TestCheckWinner(t *testing.T) {
	t.Run("Player X wins with the first row", func(t *testing.T) {
		board := []interface{}{"X", "X", "X", 4, 5, 6, 7, 8, 9}
		if !CheckWinner(board, "X") {
			t.Errorf("expected X to win, but got no winner")
		}
	})

	t.Run("Player X wins with the middle row", func(t *testing.T) {
		board := []interface{}{1, 2, 3, "X", "X", "X", 7, 8, 9}
		if !CheckWinner(board, "X") {
			t.Errorf("expected X to win, but got no winner")
		}
	})

	t.Run("Player O wins with the last row", func(t *testing.T) {
		board := []interface{}{1, 2, 3, 4, 5, 6,"O", "O", "O"}
		if !CheckWinner(board, "O") {
			t.Errorf("expected O to win, but got no winner")
		}
	})

	t.Run("Player O wins with the first column", func(t *testing.T) {
		board := []interface{}{"O", 2, 3, "O", 5, 6,"O", "X", "X"}
		if !CheckWinner(board, "O") {
			t.Errorf("expected O to win, but got no winner")
		}
	})

	t.Run("Player O wins with the second column", func(t *testing.T) {
		board := []interface{}{"X", "O", 3, "O", "O", 6, 7, "O", "X"}
		if !CheckWinner(board, "O") {
			t.Errorf("expected O to win, but got no winner")
		}
	})

	t.Run("Player X wins with the last column", func(t *testing.T) {
		board := []interface{}{1, "O", "X", "O", 5, "X", 7, "O", "X"}
		if !CheckWinner(board, "X") {
			t.Errorf("expected X to win, but got no winner")
		}
	})

	t.Run("Player O wins with  diagonal", func(t *testing.T) {
		board := []interface{}{"O", 2, 3, 4, "O", 6, 7, 8, "O"}
		if !CheckWinner(board, "O") {
			t.Errorf("expected O to win, but got no winner")
		}
	})

	t.Run("Player X wins with  diagonal", func(t *testing.T) {
		board := []interface{}{"O", 2, "X", 4, "X", 6, "X", 8, "O"}
		if !CheckWinner(board, "X") {
			t.Errorf("expected X to win, but got no winner")
		}
	})

	t.Run("No winner", func(t *testing.T) {
		board := []interface{}{"X", "O", "X", "O", "X", "O", "O", "X", "O"}
		if CheckWinner(board, "X") || CheckWinner(board, "O") {
			t.Errorf("expected no winner, but got a winner")
		}
	})
}

func TestGameOver(t *testing.T) {
	t.Run("Player X wins", func(t *testing.T) {
		board := []interface{}{"X", "X", "X", 4, 5, 6, 7, 8, 9}
		over, message := GameOver(board, "X", "O")
		if !over || message != "X wins!" {
			t.Errorf("expected game over with message 'X wins!', got %v, %q", over, message)
		}
	})

	t.Run("Player O wins", func(t *testing.T) {
		board := []interface{}{1, 2, "O", 4, "O", 6, "O", 8, 9}
		over, message := GameOver(board, "O", "X")
		if !over || message != "O wins!" {
			t.Errorf("expected game over with message 'O wins!', got %v, %q", over, message)
		}
	})

	t.Run("Game is a tie", func(t *testing.T) {
		board := []interface{}{"X", "O", "X", "O", "X", "O", "O", "X", "O"}
		over, message := GameOver(board, "X", "O")
		if !over || message != "Tie!" {
			t.Errorf("expected Tie!, got %v, %q", over, message)
		}
	})

	t.Run("Game is not over", func(t *testing.T) {
		board := []interface{}{1, 2, "X", "O", "X", "O", 7, 8, 9}
		over, _ := GameOver(board, "X", "O")
		if over {
			t.Errorf("expected game not to be over, but got game over")
		}
	})
}
