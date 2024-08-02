package test

import (
	"bufio"
	"fmt"
	"os"
)

// Tetromino structure
type Tetromino struct {
	shape [][]int
	name  string
}

// Check if an element is surrounded by ones
func isSurroundedByOnes(arr [][]int, row, col int) bool {
	if row > 0 && arr[row-1][col] == 1 &&
		row < len(arr)-1 && arr[row+1][col] == 1 {
		return true
	}
	if col > 0 && arr[row][col-1] == 1 &&
		col < len(arr[0])-1 && arr[row][col+1] == 1 {
		return true
	}
	return false
}

// Check for any surrounded one in the array
func CheckForSurroundedOne(arr [][]int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 1 && isSurroundedByOnes(arr, i, j) {
				return true
			}
		}
	}
	return false
}

// Print a 2D array
func Print2DArray(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

// Read lines from a file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Parse tetrominoes from lines
func ParseTetrominoes(lines []string) []Tetromino {
	var tetrominoes []Tetromino
	var shape [][]int

	for _, line := range lines {
		if line == "" {
			tetrominoes = append(tetrominoes, Tetromino{shape: shape})
			shape = nil
			continue
		}
		var row []int
		for _, char := range line {
			if char == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		shape = append(shape, row)
	}
	if len(shape) > 0 {
		tetrominoes = append(tetrominoes, Tetromino{shape: shape})
	}

	return tetrominoes
}

// Solve the Tetromino placement problem
func Solve(tetrominoes []Tetromino, grid [][]string, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]

	for i := 0; i < len(grid)-len(tetromino.shape)+1; i++ {
		for j := 0; j < len(grid[0])-len(tetromino.shape[0])+1; j++ {
			if canPlace(tetromino, grid, i, j) {
				place(tetromino, grid, i, j)
				if Solve(tetrominoes, grid, index+1) {
					return true
				}
				remove(tetromino, grid, i, j)
			}
		}
	}

	return false
}

// Check if a Tetromino can be placed on the grid
func canPlace(tetromino Tetromino, grid [][]string, row, col int) bool {
	for i := 0; i < len(tetromino.shape); i++ {
		for j := 0; j < len(tetromino.shape[0]); j++ {
			if tetromino.shape[i][j] == 1 && grid[row+i][col+j] != "*" {
				return false
			}
		}
	}
	return true
}

// Place a Tetromino on the grid
func place(tetromino Tetromino, grid [][]string, row, col int) {
	for i := 0; i < len(tetromino.shape); i++ {
		for j := 0; j < len(tetromino.shape[0]); j++ {
			if tetromino.shape[i][j] == 1 {
				grid[row+i][col+j] = tetromino.name
			}
		}
	}
}

// Remove a Tetromino from the grid
func remove(tetromino Tetromino, grid [][]string, row, col int) {
	for i := 0; i < len(tetromino.shape); i++ {
		for j := 0; j < len(tetromino.shape[0]); j++ {
			if tetromino.shape[i][j] == 1 {
				grid[row+i][col+j] = "*"
			}
		}
	}
}
