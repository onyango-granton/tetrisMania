package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Constants for tetromino properties
const tetrominoBlockSize = 4

// DisplayBoard outputs the board to the console.
func DisplayBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

// InitializeBoard creates an empty board of the given size.
func InitializeBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// LoadTetrominoes reads tetromino shapes from a file and labels them with unique letters.
func LoadTetrominoes(filename string) ([][]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Unable to read file: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var tetrominoes [][]string
	var currentShape []string
	var label rune = 'A'

	for _, line := range lines {
		if line == "" {
			tetrominoes = append(tetrominoes, currentShape)
			currentShape = nil
			label++
			continue
		}

		transformedLine := replaceLineCharacters(line, label)
		if transformedLine == "" {
			return nil, fmt.Errorf("ERROR: Invalid character in tetromino")
		}
		currentShape = append(currentShape, transformedLine)
	}

	if len(currentShape) > 0 {
		tetrominoes = append(tetrominoes, currentShape) // Add the last tetromino
	}

	return tetrominoes, nil
}

// replaceLineCharacters transforms the line by replacing '#' with the tetromino label and keeping '.'.
func replaceLineCharacters(line string, label rune) string {
	var builder strings.Builder
	for _, char := range line {
		if char == '#' {
			builder.WriteRune(label)
		} else if char == '.' {
			builder.WriteRune(char)
		} else {
			return "" // Invalid character found
		}
	}
	return builder.String()
}

// SolvePuzzle attempts to fit all tetrominoes on the board using a recursive approach.
func SolvePuzzle(board [][]string, tetrominoes [][]string) [][]string {
	if fitTetrominoes(board, tetrominoes, 0) {
		return board
	}
	return nil
}

// fitTetrominoes tries to place each tetromino on the board recursively.
func fitTetrominoes(board [][]string, tetrominoes [][]string, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	for y := range board {
		for x := range board[y] {
			if canFit(board, tetromino, x, y) {
				placeTetromino(board, tetromino, x, y)
				if fitTetrominoes(board, tetrominoes, index+1) {
					return true
				}
				removeTetromino(board, tetromino, x, y)
			}
		}
	}
	return false
}

// canFit checks if the tetromino can be placed at the specified position on the board.
func canFit(board [][]string, tetromino []string, x, y int) bool {
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

// placeTetromino sets the tetromino on the board at the specified position.
func placeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

// removeTetromino removes the tetromino from the board at the specified position.
func removeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}

// TrimTetrominoes removes empty rows and columns from tetromino shapes.
func TrimTetrominoes(tetrominoes [][]string) [][]string {
	var trimmedShapes [][]string
	for _, shape := range tetrominoes {
		trimmedShapes = append(trimmedShapes, trimShape(shape))
	}
	return trimmedShapes
}

// trimShape removes empty rows and columns from a single tetromino shape.
func trimShape(shape []string) []string {
	var trimmed []string
	columnFilled := make([]bool, len(shape[0]))

	for col := range columnFilled {
		for row := range shape {
			if shape[row][col] != '.' {
				columnFilled[col] = true
				break
			}
		}
	}

	for _, row := range shape {
		var newRow strings.Builder
		for col, filled := range columnFilled {
			if filled {
				newRow.WriteByte(row[col])
			}
		}
		trimmed = append(trimmed, newRow.String())
	}
	return trimmed
}

// ValidateTetrominoes checks the validity of the tetromino shapes.
func ValidateTetrominoes(tetrominoes [][]string) string {
	if len(tetrominoes) > 26 {
		return "Invalid File"
	}
	for _, shape := range tetrominoes {
		if len(shape) != tetrominoBlockSize {
			return "Invalid File"
		}
		if result := checkConnections(shape); result != "ok" {
			return result
		}
	}
	return "ok"
}

// checkConnections verifies the tetromino connections.
func checkConnections(shape []string) string {
	connectionCount, blockCount := 0, 0
	for i, row := range shape {
		for j, char := range row {
			if char != '.' {
				blockCount++
				connectionCount += countConnections(shape, i, j, char)
			}
		}
	}

	if connectionCount < 6 || blockCount != 4 {
		return "Invalid File"
	}
	return "ok"
}

// countConnections counts the adjacent connections for a block.
func countConnections(shape []string, i, j int, char rune) int {
	connections := 0
	if i > 0 && shape[i-1][j] == byte(char) {
		connections++
	}
	if i < len(shape)-1 && shape[i+1][j] == byte(char) {
		connections++
	}
	if j > 0 && shape[i][j-1] == byte(char) {
		connections++
	}
	if j < len(shape[i])-1 && shape[i][j+1] == byte(char) {
		connections++
	}
	return connections
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: Invalid number of arguments")
	}

	tetrominoes, err := LoadTetrominoes(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if validationMessage := ValidateTetrominoes(tetrominoes); validationMessage != "ok" {
		fmt.Println("ERROR:", validationMessage)
		return
	}

	tetrominoes = TrimTetrominoes(tetrominoes)

	boardSize := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))
	var finalBoard [][]string
	for {
		board := InitializeBoard(boardSize)
		finalBoard = SolvePuzzle(board, tetrominoes)
		if finalBoard != nil {
			break
		}
		boardSize++
	}

	DisplayBoard(finalBoard)
}
