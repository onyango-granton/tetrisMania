package test

import "testing"

func TestSolve(t *testing.T) {
	grid := [][]string{
		{"*", "*", "*"},
		{"*", "*", "*"},
		{"*", "*", "*"},
	}

	tetrominoes := []Tetromino{
		{
			shape: [][]int{
				{1, 1},
				{1, 1},
			},
			name: "T",
		},
	}

	if !Solve(tetrominoes, grid, 0) {
		t.Errorf("Solve() expected true, got false")
	}

	expectedGrid := [][]string{
		{"T", "T", "*"},
		{"T", "T", "*"},
		{"*", "*", "*"},
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != expectedGrid[i][j] {
				t.Errorf("Solve() expected %v, got %v", expectedGrid, grid)
			}
		}
	}
}
