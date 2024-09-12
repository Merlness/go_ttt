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
		pos := 5 // Position already taken by "X"
		if IsValidMove(board, pos) {
			t.Errorf("expected move at position %d to be invalid", pos)
		}
	})

	t.Run("invalid move (out of range)", func(t *testing.T) {
		pos := 10 // Out of range
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
