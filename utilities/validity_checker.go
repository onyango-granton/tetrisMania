package utilities

func Valid(tetro [][]string) string {
	for _, tet := range tetro {
		// Checks if length of a tetromino is 4
		if len(tet) != 4 {
			return "Invalid File"
		}
		ans := Connection(tet)
		if ans == "Invalid File" {
			return "Invalid File"
		}
		//checks the length of the string if it is 4
		for _, str := range tet {
			if len(str) != 4 {
				return "Invalid File"
			}
		}
	}
	return "ok"
}

// checks if the connection from one charachter to the other is 6 and above and if there are only 4 charachters
func Connection(tet []string) string {
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
		return "Invalid File"
	}
	return "ok"
}
