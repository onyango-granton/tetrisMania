package utilities

import (
	"fmt"
	"os"
	"strings"
)

// Reader reads tetrominos from a text file and append hashes(#) with letters.
func Reader() [][]string {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		os.Exit(0)
	}
	output, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	output2 := strings.Split(string(output), "\n")

	var tetro []string
	var tetromino [][]string
	var letter rune = 'A'

	for i, str := range output2 {
		if str != "" {
			var new strings.Builder
			for _, char := range str {
				if char == '#' {
					new.WriteRune(letter)
				} else if char == '.' {
					new.WriteRune(char)
				} else {
					fmt.Println("ERROR")
					os.Exit(0)
				}
			}
			tetro = append(tetro, new.String())
		}
		if str == "" || i == len(output2)-1 {
			tetromino = append(tetromino, tetro)
			tetro = []string{}
			letter++
		}
	}

	return tetromino
}
