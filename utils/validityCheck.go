package utils

import "errors"

/*allOne func Checks if both input integers are equal to 1.*/
func allOne(num1, num2 int) bool {
	if num1 == 1 {
		return num1 == num2
	}
	return false
}

/*
isSurroundedByOnes Determines if a specific element in a 2D array is surrounded by elements that are all 1s.
This function checks horizontally and vertically adjacent elements.
*/
func isSurroundedByOnes(arr [][]int, row, col int) bool {
	// Check horizontally
	if col-1 >= 0 && allOne(arr[row][col-1], arr[row][col]) || col+1 < len(arr[row]) && allOne(arr[row][col+1], arr[row][col]) {
		return true
	}
	// Check vertically
	if row-1 >= 0 && allOne(arr[row][col], arr[row-1][col]) || row+1 < len(arr) && allOne(arr[row][col], arr[row+1][col]) {
		return true
	}
	return false
}

/*
fullyConnected func Checks if a tetromino shape (represented as a 2D slice) is fully connected. A tetromino is considered fully
connected if each '1' in the shape is directly connected to at least one other '1' horizontally or vertically.
*/
func fullyConnected(tetro [][]int) bool {
	connection := 0
	for row := range tetro {
		for col := range tetro[row] {
			if tetro[row][col] == 1 {
				if col+1 <= len(tetro[row])-1 && tetro[row][col+1] == 1 {
					connection++
				}
				if col-1 >= 0 && tetro[row][col-1] == 1 {
					connection++
				}
				if row+1 <= len(tetro)-1 && tetro[row+1][col] == 1 {
					connection++
				}
				if row-1 >= 0 && tetro[row-1][col] == 1 {
					connection++
				}
			}
		}
	}
	// fmt.Println(connection)
	if connection == 6 || connection == 8 {
		return true
	} else {
		return false
	}
}

/*
isValidTetro Validates if a given tetromino shape is valid. A valid tetromino has exactly 4 '1's,
is fully connected, and does not have more than 4 borders surrounded by '1's.
*/
func isValidTetro(tetro [][]int) (bool, error) {
	var bordercount int
	var linecount int

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {
			if tetro[row][col] == 1 {
				linecount++
			}
			if tetro[row][col] == 1 && isSurroundedByOnes(tetro, row, col) {
				bordercount++
			}
		}
	}

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {
		}
	}

	if bordercount > 4 || linecount > 4 {
		return false, errors.New("ERROR")
	} else {
		if fullyConnected(tetro) {
			return true, nil
		} else {
			return false, errors.New("ERROR")
		}
	}
}

// isConnected checks for non connected tetrominoes
func isConnected(arr []string) bool {
	count := 1
	for _, ch := range arr {
		if count%5 == 0 && ch != "" {
			return true
		}
		count++
	}
	return false
}
