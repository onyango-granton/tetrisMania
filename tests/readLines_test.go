package test

import (
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("unable to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write some lines to the file
	lines := []string{"....", "####", "..#."}
	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			t.Fatalf("unable to write to temp file: %v", err)
		}
	}
	file.Close()

	readLines, err := ReadLines(file.Name())
	if err != nil {
		t.Fatalf("ReadLines() unexpected error: %v", err)
	}

	if len(readLines) != len(lines) {
		t.Errorf("ReadLines() expected %d lines, got %d lines", len(lines), len(readLines))
	}

	for i, line := range readLines {
		if line != lines[i] {
			t.Errorf("ReadLines() expected line %d to be %q, got %q", i, lines[i], line)
		}
	}
}
