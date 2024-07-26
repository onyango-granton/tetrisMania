package main

import (
	"fmt"
	"os"
)

func fullyConnected() bool {
	//var tetro = [][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}}
	var tetro = [][]int{{1, 1, 1, 1}}
	connection := 0
	for row := range tetro {
		for col := range tetro[row] {
			if tetro[row][col] == 1 {
				if col+1 <= len(tetro[row])-1 && tetro[row][col+1] == 1 {
					if col == 0 {
						fmt.Print("first")
					}
					if col == 1 && row == 2 {
						fmt.Print("third")
					}
					fmt.Println("right connection")
					connection++
				}
				if col-1 >= 0 && tetro[row][col-1] == 1 {
					if col == 1 && row == 1 {
						fmt.Print("second")
					}
					fmt.Println("left connection")
					connection++
				}
				if row+1 <= len(tetro)-1 && tetro[row+1][col] == 1 {
					if col == 1 && row == 1 {
						fmt.Print("second")
					}
					fmt.Println("down connection")
					connection++
				}
				if row-1 >= 0 && tetro[row-1][col] == 1 {
					if col == 1 && row == 2 {
						fmt.Print("third")
					}
					fmt.Println("up connection")
					connection++
				}
			}
		}
	}
	fmt.Println(connection)
	if connection == 6 || connection == 8 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(os.Getwd())
	fmt.Println(fullyConnected())
}
