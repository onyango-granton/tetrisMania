package main

import (
	"tetris/utilities"
)

func main() {
	tetro := utilities.Reader()

	utilities.Valid(tetro)

	tetro = utilities.Trimmer(tetro)
}
