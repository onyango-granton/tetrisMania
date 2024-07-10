package main

import (
	"fmt"
	"math"
	"tetris/utilities"
)

func main() {
	tetro := utilities.Reader()

	err := utilities.Valid(tetro)
	if err == "Invalid file" {
		fmt.Println("ERROR")
		return
	}

	tetro = utilities.Trimmer(tetro)

	Size := int(math.Ceil(math.Sqrt(float64(len(tetro) * 4))))
	var finalboard [][]string
	for {
		board := utilities.CreateBoard(Size)
		finalboard = utilities.Solve(board, tetro)
		if finalboard != nil {
			break
		}
		Size++
	}

	utilities.Print(finalboard)
}
