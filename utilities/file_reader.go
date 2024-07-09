package utilities

import (
	"fmt"
	"os"
	"strings"
)

func Reader() [][]string {
	output, _ := os.ReadFile(os.Args[1])

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
					fmt.Println("invalid file")
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
