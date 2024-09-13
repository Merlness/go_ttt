package main

import (
	"testing"
)


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

func TestGetAIMove(t *testing.T) {
	t.Run("AI places 1 as its first move on an empty board", func(t *testing.T) {
		board := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

		expectedMove := 1
		aiMove := GetAIMove(board)

		if aiMove != expectedMove {
			t.Errorf("expected AI to choose %d, but got %d", expectedMove, aiMove)
		}
	})
}
