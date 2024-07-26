package main


import (
	"fmt"
)

var directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func isValid(grid [][]int, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}

func countAdjacentOnes(grid [][]int, visited [][]bool, x, y int) int {
	stack := [][]int{{x, y}}
	count := 0

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, dir := range directions {
			newX, newY := point[0]+dir[0], point[1]+dir[1]
			if isValid(grid, newX, newY) && grid[newX][newY] == 1 && !visited[newX][newY] {
				count++
				visited[newX][newY] = true
				stack = append(stack, []int{newX, newY})
			}
		}
	}
	return count
}

func main() {
	grid := [][]int{
		{1, 0, 0, 1},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	totalCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				visited[i][j] = true
				totalCount += countAdjacentOnes(grid, visited, i, j)
			}
		}
	}

	fmt.Println("Number of adjacent ones:", totalCount)
}
