package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Tetromino struct {
	shape []string
	name  string
}

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

// Reader reads tetrominos from a text file and returns a slice of Tetromino structs.
func Reader() []Tetromino {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid number of arguments")
		os.Exit(1)
	}
	output, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Unable to read file")
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	var tetrominoes []Tetromino
	var currentShape []string
	var letter rune = 'A'

	for _, line := range lines {
		if line == "" {
			tetrominoes = append(tetrominoes, Tetromino{shape: currentShape, name: string(letter)})
			currentShape = nil
			letter++
			continue
		}

		newLine := transformLine(line, letter)
		if newLine == "" {
			fmt.Fprintln(os.Stderr, "ERROR: Invalid character in tetromino")
			os.Exit(1)
		}
		currentShape = append(currentShape, newLine)
	}

	if len(currentShape) > 0 {
		tetrominoes = append(tetrominoes, Tetromino{shape: currentShape, name: string(letter)}) // Add the last tetromino
	}

	return tetrominoes
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
func Solve(board [][]string, tetrominoes []Tetromino) [][]string {
	if placeTetrominos(board, tetrominoes, 0) {
		return board
	}
	return nil
}

func placeTetrominos(board [][]string, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetromino.shape, x, y) {
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

func canPlace(board [][]string, shape []string, x, y int) bool {
	for dy, row := range shape {
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

func placeTetromino(board [][]string, tetromino Tetromino, x, y int) {
	for dy, row := range tetromino.shape {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

func removeTetromino(board [][]string, tetromino Tetromino, x, y int) {
	for dy, row := range tetromino.shape {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}

func Trimmer(tetro []Tetromino) []Tetromino {
	var newTetrominoes []Tetromino
	for _, tet := range tetro {
		newTetrominoes = append(newTetrominoes, Tetromino{shape: trimTetromino(tet.shape), name: tet.name})
	}
	return newTetrominoes
}

func trimTetromino(shape []string) []string {
	var result []string
	columnHasLetters := make([]bool, len(shape[0]))

	for col := range columnHasLetters {
		for row := range shape {
			if shape[row][col] != '.' {
				columnHasLetters[col] = true
				break
			}
		}
	}

	for _, row := range shape {
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

func Valid(tetro []Tetromino) string {
	if len(tetro) > 26 {
		return "Invalid File"
	}
	for _, tet := range tetro {
		if len(tet.shape) != 4 {
			return "Invalid File"
		}
		if ans := Connection(tet.shape); ans != "ok" {
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
	tetro := Reader()

	if err := Valid(tetro); err != "ok" {
		fmt.Println("ERROR")
		return
	}

	tetro = Trimmer(tetro)

	size := int(math.Ceil(math.Sqrt(float64(len(tetro) * 4))))
	fmt.Println(size)
	var finalBoard [][]string
	for {
		board := CreateBoard(size)
		finalBoard = Solve(board, tetro)
		if finalBoard != nil {
			break
		}
		size++
	}

	Print(finalBoard)
}
