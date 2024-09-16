package main

import (
    "testing"
    "reflect"
)

func AssertEqual(t testing.TB, got, want interface{}) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}

func AssertTrue(t testing.TB, condition bool, msg string) {
    t.Helper()
    if !condition {
        t.Errorf("Assertion failed: %s", msg)
    }
}


func TestDisplay(t *testing.T) {
	t.Run("displays the default board", func(t *testing.T) {
		grid := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Display(grid)
		want := "1 | 2 | 3\n4 | 5 | 6\n7 | 8 | 9"
		AssertEqual(t, got, want)
	})

	t.Run("displays a mix of numbers and tokens", func(t *testing.T) {
		grid := []interface{}{"X", "O", 3, "X", 5, "O", 7, 8, 9}
		got := Display(grid)
		want := "X | O | 3\nX | 5 | O\n7 | 8 | 9"
		AssertEqual(t, got, want)
	})

	t.Run("displays all tokens", func(t *testing.T) {
		grid := []interface{}{"X", "O", "X", "X", "O", "X", "O", "X", "O"}
		got := Display(grid)
		want := "X | O | X\nX | O | X\nO | X | O"
		AssertEqual(t, got, want)
	})
}

func TestCurrentPlayerToken(t *testing.T) {
	board := []interface{}{1, "X", 3, "O", 5, 6, 7, 8, 9}

	t.Run("Player 1's turn when odd number of integers left", func(t *testing.T) {
        resetGameStruct()
		got := CurrentPlayerToken(board, gameStruct)
		want := gameStruct.Token1

		AssertEqual(t, got, want)
	})

	board1 := []interface{}{1, "X", "X", "O", 5, 6, 7, 8, 9}

	t.Run("Player 2's turn", func(t *testing.T) {
	    resetGameStruct()
		got := CurrentPlayerToken(board1, gameStruct)
		want := gameStruct.Token2

		AssertEqual(t, got, want)
	})
}

func TestCheckWinner(t *testing.T) {
	t.Run("Player X wins with the first row", func(t *testing.T) {
		board := []interface{}{"X", "X", "X", 4, 5, 6, 7, 8, 9}
		AssertTrue(t, CheckWinner(board, "X"), "expected X to win")
	})

	t.Run("Player X wins with the middle row", func(t *testing.T) {
		board := []interface{}{1, 2, 3, "X", "X", "X", 7, 8, 9}
		AssertTrue(t, CheckWinner(board, "X"), "expected X to win")
	})

	t.Run("Player O wins with the last row", func(t *testing.T) {
		board := []interface{}{1, 2, 3, 4, 5, 6,"O", "O", "O"}
		AssertTrue(t, CheckWinner(board, "O"), "expected O to win")
	})

	t.Run("Player O wins with the first column", func(t *testing.T) {
		board := []interface{}{"O", 2, 3, "O", 5, 6,"O", "X", "X"}
		AssertTrue(t, CheckWinner(board, "O"), "expected O to win")
	})

	t.Run("Player O wins with the second column", func(t *testing.T) {
		board := []interface{}{"X", "O", 3, "O", "O", 6, 7, "O", "X"}
		AssertTrue(t, CheckWinner(board, "O"), "expected O to win")
	})

	t.Run("Player X wins with the last column", func(t *testing.T) {
		board := []interface{}{1, "O", "X", "O", 5, "X", 7, "O", "X"}
		AssertTrue(t, CheckWinner(board, "X"), "expected X to win")
	})

	t.Run("Player O wins with  diagonal", func(t *testing.T) {
		board := []interface{}{"O", 2, 3, 4, "O", 6, 7, 8, "O"}
		AssertTrue(t, CheckWinner(board, "O"), "expected O to win")
	})

	t.Run("Player X wins with  diagonal", func(t *testing.T) {
		board := []interface{}{"O", 2, "X", 4, "X", 6, "X", 8, "O"}
		AssertTrue(t, CheckWinner(board, "X"), "expected X to win")
	})

	t.Run("No winner", func(t *testing.T) {
		board := []interface{}{"X", "O", "X", "O", "X", "O", "O", "X", "O"}
		if CheckWinner(board, "X") || CheckWinner(board, "O") {
			t.Errorf("expected no winner")
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

		AssertTrue(t, !over, "expected game not to be over, but got game over")
	})
}

func TestCountAvailableMoves(t *testing.T) {
    t.Run("Counts available moves correctly", func(t *testing.T) {
        board := []interface{}{"X", "O", 3, "X", 5, "O", 7, 8, 9}
        expectedCount := 5

        count := countAvailableMoves(board)
		AssertEqual(t, count, expectedCount)
    })
}

func TestAvailableMoves(t *testing.T) {
    t.Run("Returns correct list of available moves", func(t *testing.T) {
        board := []interface{}{"X", "O", 3, "X", 5, "O", 7, 8, 9}
        expectedMoves := []int{3, 5, 7, 8, 9}

        moves := AvailableMoves(board)
		AssertEqual(t, moves, expectedMoves)
    })
}


func TestCopyBoard(t *testing.T) {
    t.Run("CopyBoard creates an exact copy of the board", func(t *testing.T) {
        board := []interface{}{"X", "O", 3, "X", 5, "O", 7, 8, 9}
        copiedBoard := CopyBoard(board)

		AssertEqual(t, board, copiedBoard)
    })
}
