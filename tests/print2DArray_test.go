package test

import "testing"

func TestPrint2DArray(t *testing.T) {
	arr := [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}

	Print2DArray(arr)
}
