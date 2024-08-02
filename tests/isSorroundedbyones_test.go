package test

import "testing"

func TestIsSurroundedByOnes(t *testing.T) {
	arr := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	if !isSurroundedByOnes(arr, 1, 1) {
		t.Errorf("isSurroundedByOnes() expected true, got false")
	}

	if isSurroundedByOnes(arr, 0, 0) {
		t.Errorf("isSurroundedByOnes() expected false, got true")
	}
}
