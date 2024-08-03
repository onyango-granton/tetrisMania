package utils

import "fmt"

/*
PrintGrid function is designed to print a two-dimensional slice of strings in a grid format.
Each element in the grid is printed with a space separating it from the next, and each row is printed on a new line.
*/
func PrintGrid(grid [][]string) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Print(grid[row][col] + " ")
		}
		fmt.Println()
	}
}
