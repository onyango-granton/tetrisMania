package test

import "testing"

func TestCheckForSurroundedOne(t *testing.T) {
	arr := [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}

	if !CheckForSurroundedOne(arr) {
		t.Errorf("CheckForSurroundedOne() expected true, got false")
	}

	arr = [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	if CheckForSurroundedOne(arr) {
		t.Errorf("CheckForSurroundedOne() expected false, got true")
	}
}
