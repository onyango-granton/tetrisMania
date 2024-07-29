package main

import (
	"fmt"
	"os"

	"tetris/finalProject/utils"
)

func main() {
	// 	if len(os.Args) != 2 {
	// 		fmt.Println(`invalid number of arguments
	// Usage:
	// go run . <filename.txt>`)
	// 	}

	// 	//multiple txt extensions
	// 	filenameList := strings.Split(os.Args[1],".")
	// 	if len(filenameList) != 2{
	// 		fmt.Println(`invalid filename and/or extension
	// Usage:
	// go run . <filename.txt>`)
	// 	}

	err := utils.ErrorHandling()
	if err != nil {
		fmt.Println(err)
		return
	}

	tetrogroup, gridSize := utils.TetroGroupFunc(os.Args[1])
	grid := utils.InitGrid(gridSize)
	if utils.CompleteGrid(tetrogroup, grid,0){
		utils.PrintGrid(grid)
	} else {
		fmt.Println("No solutions found")
	}
}
