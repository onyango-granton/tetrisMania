package utilities

import "strings"

func Trimmer(tetro [][]string) [][]string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var newtetro [][]string
	for _, tet := range tetro {
		var newstet []string
		for _, str := range tet {
			for _, letter := range letters {
				if strings.Contains(str, string(letter)) {
					newstet = append(newstet, str)
				}
			}
		}
		newtetro = append(newtetro, newstet)
	}

	var newtetro2 [][]string

	for _, tet := range newtetro {
		width := len(tet[0])

		// Check each column for the presence of letters ('A')
		columnHasLetters := make([]bool, width)
		for col := 0; col < width; col++ {
			for row := 0; row < len(tet); row++ {
				for _, letter := range letters {
					if tet[row][col] == byte(letter) {
						columnHasLetters[col] = true
						break
					}
				}
			}
		}

		// Remove columns that do not contain any letters ('A')
		var result []string
		for _, row := range tet {
			var newRow strings.Builder
			for col := 0; col < width; col++ {
				if columnHasLetters[col] {
					newRow.WriteByte(row[col])
				}
			}
			result = append(result, newRow.String())
		}
		newtetro2 = append(newtetro2, result)
	}
	return newtetro2
}
