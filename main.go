package main

import (
	"math"
	"tetris/utilities"
)

func main() {
	tetro := utilities.Reader()

	utilities.Valid(tetro)

	tetro = utilities.Trimmer(tetro)

	Size := int(math.Ceil(math.Sqrt(float64(len(tetro) * 4))))
	var finalboard [][]string
	for {
		board := utilities.CreateInitialBoard(Size)
	}

}
