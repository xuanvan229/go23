package main

import (
	"testing"
)

func TestCountRectangles(t *testing.T) {
	// Test case 1: A simple rectangle
	board1 := [][]int{
		{1, 1},
		{1, 1},
	}
	expected1 := 1
	if result1 := countRectangles(board1); result1 != expected1 {
		t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
	}

	// Test case 2: Two separate rectangles
	board2 := [][]int{
		{1, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}
	expected2 := 2
	if result2 := countRectangles(board2); result2 != expected2 {
		t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
	}

	// Test case 3: No rectangles, all cells are empty
	board3 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	expected3 := 0
	if result3 := countRectangles(board3); result3 != expected3 {
		t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
	}

	// Test case 4: Complex board with multiple rectangles
	board4 := [][]int{
		{1, 0, 1, 1, 0},
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1},
	}
	expected4 := 5
	if result4 := countRectangles(board4); result4 != expected4 {
		t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
	}
}
