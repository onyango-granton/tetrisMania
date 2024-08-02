package test

import "testing"

func TestParseTetrominoes(t *testing.T) {
	lines := []string{
		"....",
		"####",
		"..#.",
		"",
		"#...",
		"#...",
		"#...",
		"#...",
	}

	tetrominoes := ParseTetrominoes(lines)

	expectedShapes := [][][]int{
		{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 1, 0},
		},
		{
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
	}

	for i, tetromino := range tetrominoes {
		for j := range tetromino.shape {
			for k := range tetromino.shape[j] {
				if tetromino.shape[j][k] != expectedShapes[i][j][k] {
					t.Errorf("ParseTetrominoes() expected %v, got %v", expectedShapes, tetromino.shape)
				}
			}
		}
	}
}
