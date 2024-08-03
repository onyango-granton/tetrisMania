package main

import (
	"fmt"
	"os"

	"tetris/utils"
)

func main() {
	err := utils.ErrorHandling()
	if err != nil {
		fmt.Println(err)
		return
	}

	tetrogroup, gridSize := utils.TetroGroupFunc(os.Args[1])
	for {
		grid := utils.InitGrid(gridSize)
		if utils.CompleteGrid(tetrogroup, grid,0) {
			utils.PrintGrid(grid)
			break
		}
		gridSize++
	}

}
