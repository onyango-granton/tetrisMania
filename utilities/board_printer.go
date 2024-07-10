package utilities

import "fmt"

// PrintBoard prints the final square.
func PrintBoard(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
