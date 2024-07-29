package utilities

import "fmt"

// Print prints the final square.
func Print(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
