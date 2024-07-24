package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"tetris/utilities"
)

// Constants for tetromino properties
const tetrominoSize = 4

// Print prints the final square.
func Print(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

// CreateBoard creates a board for placing tetrominos.
func CreateBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// Reader reads tetrominos from a text file and appends letters to the tetrominos.
func Reader(filename string) ([][]string, error) {
	output, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Unable to read file: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var tetromino [][]string
	var currentTetromino []string
	var letter rune = 'A'

	for _, line := range lines {
		if line == "" {
			tetromino = append(tetromino, currentTetromino)
			currentTetromino = nil
			letter++
			continue
		}

		newLine := transformLine(line, letter)
		if newLine == "" {
			return nil, fmt.Errorf("ERROR: Invalid character in tetromino")
		}
		currentTetromino = append(currentTetromino, newLine)
	}

	if len(currentTetromino) > 0 {
		tetromino = append(tetromino, currentTetromino) // Add the last tetromino
	}

	return tetromino, nil
}

func transformLine(line string, letter rune) string {
	var builder strings.Builder
	for _, char := range line {
		if char == '#' {
			builder.WriteRune(letter)
		} else if char == '.' {
			builder.WriteRune(char)
		} else {
			return "" // Invalid character found
		}
	}
	return builder.String()
}

// Solve attempts to place tetrominos in the smallest square possible using a recursive method.
func Solve(board [][]string, tetrominoes [][]string) [][]string {
	if placeTetrominos(board, tetrominoes, 0) {
		return board
	}
	return nil
}

func placeTetrominos(board [][]string, tetrominoes [][]string, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetromino, x, y) {
				placeTetromino(board, tetromino, x, y)
				if placeTetrominos(board, tetrominoes, index+1) {
					return true
				}
				removeTetromino(board, tetromino, x, y)
			}
		}
	}
	return false
}

func canPlace(board [][]string, tetromino []string, x, y int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

func placeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

func removeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}

func Trimmer(tetro [][]string) [][]string {
	var newTetromino [][]string
	for _, tet := range tetro {
		newTetromino = append(newTetromino, trimTetromino(tet))
	}
	return newTetromino
}

func trimTetromino(tet []string) []string {
	var result []string
	columnHasLetters := make([]bool, len(tet[0]))

	for col := range columnHasLetters {
		for row := range tet {
			if tet[row][col] != '.' {
				columnHasLetters[col] = true
				break
			}
		}
	}

	for _, row := range tet {
		var newRow strings.Builder
		for col, hasLetter := range columnHasLetters {
			if hasLetter {
				newRow.WriteByte(row[col])
			}
		}
		result = append(result, newRow.String())
	}
	return result
}

func Valid(tetro [][]string) string {
	if len(tetro) > 26 {
		return "Invalid File"
	}
	for _, tet := range tetro {
		if len(tet) != tetrominoSize {
			return "Invalid File"
		}
		if ans := Connection(tet); ans != "ok" {
			return ans
		}
	}
	return "ok"
}

func Connection(tet []string) string {
	countConnections, countChar := 0, 0
	for i, str := range tet {
		for j, char := range str {
			if char != '.' {
				countChar++
				countConnections += countAdjacentConnections(tet, i, j, char)
			}
		}
	}

	if countConnections < 6 || countChar != 4 {
		return "Invalid File"
	}
	return "ok"
}

func countAdjacentConnections(tet []string, i, j int, char rune) int {
	connections := 0
	if i > 0 && tet[i-1][j] == byte(char) {
		connections++
	}
	if i < len(tet)-1 && tet[i+1][j] == byte(char) {
		connections++
	}
	if j > 0 && tet[i][j-1] == byte(char) {
		connections++
	}
	if j < len(tet[i])-1 && tet[i][j+1] == byte(char) {
		connections++
	}
	return connections
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: Invalid number of arguments")
	}

	tetro, err := Reader(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if errMsg := utilities.Valid(tetro); errMsg != "ok" {
		fmt.Println("ERROR:", errMsg)
		return
	}

	tetro = utilities.Trimmer(tetro)

	size := int(math.Ceil(math.Sqrt(float64(len(tetro) * 4))))
	var finalBoard [][]string
	for {
		board := utilities.CreateBoard(size)
		finalBoard = utilities.Solve(board, tetro)
		if finalBoard != nil {
			break
		}
		size++
	}

	utilities.Print(finalBoard)
}
