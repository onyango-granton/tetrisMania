package utilities

import (
	"fmt"
	"os"
)

func Valid(tetro [][]string) {
	for _, tet := range tetro {
		if len(tet) != 4 {
			fmt.Println("invalid file")
			os.Exit(0)
		}
		Connection(tet)
		for _, str := range tet {
			if len(str) != 4 {
				fmt.Println("invalid file")
				os.Exit(0)
			}
		}
	}
}

func Connection(tet []string) {
	countConnections := 0
	countchar := 0
	for i, str := range tet {
		for j, char := range str {
			if char != '.' {
				countchar++
				if i > 0 && tet[i-1][j] == byte(char) {
					countConnections++
				}
				if i < len(tet)-1 && tet[i+1][j] == byte(char) {
					countConnections++
				}
				if j > 0 && tet[i][j-1] == byte(char) {
					countConnections++
				}
				if j < len(str)-1 && tet[i][j+1] == byte(char) {
					countConnections++
				}
			}

		}
	}

	if countConnections < 6 || countchar != 4 {
		fmt.Println("invalid file")
		os.Exit(0)
	}
}