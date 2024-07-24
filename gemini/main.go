package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Constants for piece properties
const pieceDimension = 4

// Display outputs the final grid.
func Display(grid [][]string) {
	for _, line := range grid {
		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println()
}

// InitializeGrid creates a grid for positioning pieces.
func InitializeGrid(size int) [][]string {
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return grid
}

// LoadPieces reads pieces from a file and assigns characters to the pieces.
func LoadPieces(filepath string) ([][]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Unable to read file: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var pieces [][]string
	var currentPiece []string
	var identifier rune = 'A'

	for _, line := range lines {
		if line == "" {
			pieces = append(pieces, currentPiece)
			currentPiece = nil
			identifier++
			continue
		}

		processedLine := convertLine(line, identifier)
		if processedLine == "" {
			return nil, fmt.Errorf("ERROR: Invalid character in piece")
		}
		currentPiece = append(currentPiece, processedLine)
	}

	if len(currentPiece) > 0 {
		pieces = append(pieces, currentPiece) // Add the last piece
	}

	return pieces, nil
}

func convertLine(line string, identifier rune) string {
	var sb strings.Builder
	for _, char := range line {
		if char == '#' {
			sb.WriteRune(identifier)
		} else if char == '.' {
			sb.WriteRune(char)
		} else {
			return "" // Invalid character found
		}
	}
	return sb.String()
}

// Arrange attempts to fit pieces in the smallest grid possible using recursion.
func Arrange(grid [][]string, pieces [][]string) [][]string {
	if positionPieces(grid, pieces, 0) {
		return grid
	}
	return nil
}

func positionPieces(grid [][]string, pieces [][]string, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]
	for y := range grid {
		for x := range grid[y] {
			if canFit(grid, piece, x, y) {
				placePiece(grid, piece, x, y)
				if positionPieces(grid, pieces, index+1) {
					return true
				}
				removePiece(grid, piece, x, y)
			}
		}
	}
	return false
}

func canFit(grid [][]string, piece []string, x, y int) bool {
	for dy, row := range piece {
		for dx, char := range row {
			if char != '.' {
				if y+dy >= len(grid) || x+dx >= len(grid[0]) || grid[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

func placePiece(grid [][]string, piece []string, x, y int) {
	for dy, row := range piece {
		for dx, char := range row {
			if char != '.' {
				grid[y+dy][x+dx] = string(char)
			}
		}
	}
}

func removePiece(grid [][]string, piece []string, x, y int) {
	for dy, row := range piece {
		for dx, char := range row {
			if char != '.' {
				grid[y+dy][x+dx] = "."
			}
		}
	}
}

func TrimPieces(pieces [][]string) [][]string {
	var trimmedPieces [][]string
	for _, piece := range pieces {
		trimmedPieces = append(trimmedPieces, trimPiece(piece))
	}
	return trimmedPieces
}

func trimPiece(piece []string) []string {
	var trimmed []string
	columnFlags := make([]bool, len(piece[0]))

	for col := range columnFlags {
		for row := range piece {
			if piece[row][col] != '.' {
				columnFlags[col] = true
				break
			}
		}
	}

	for _, row := range piece {
		var newRow strings.Builder
		for col, hasChar := range columnFlags {
			if hasChar {
				newRow.WriteByte(row[col])
			}
		}
		trimmed = append(trimmed, newRow.String())
	}
	return trimmed
}

func ValidatePieces(pieces [][]string) string {
	if len(pieces) > 26 {
		return "Invalid File"
	}
	for _, piece := range pieces {
		if len(piece) != pieceDimension {
			return "Invalid File"
		}
		if result := CheckConnection(piece); result != "ok" {
			return result
		}
	}
	return "ok"
}

func CheckConnection(piece []string) string {
	connectionCount, charCount := 0, 0
	for i, str := range piece {
		for j, char := range str {
			if char != '.' {
				charCount++
				connectionCount += countAdjacentConnections(piece, i, j, char)
			}
		}
	}

	if connectionCount < 6 || charCount != 4 {
		return "Invalid File"
	}
	return "ok"
}

func countAdjacentConnections(piece []string, i, j int, char rune) int {
	connections := 0
	if i > 0 && piece[i-1][j] == byte(char) {
		connections++
	}
	if i < len(piece)-1 && piece[i+1][j] == byte(char) {
		connections++
	}
	if j > 0 && piece[i][j-1] == byte(char) {
		connections++
	}
	if j < len(piece[i])-1 && piece[i][j+1] == byte(char) {
		connections++
	}
	return connections
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: Invalid number of arguments")
	}

	pieces, err := LoadPieces(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if validationMessage := ValidatePieces(pieces); validationMessage != "ok" {
		fmt.Println("ERROR:", validationMessage)
		return
	}

	pieces = TrimPieces(pieces)

	gridSize := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))
	var completedGrid [][]string
	for {
		grid := InitializeGrid(gridSize)
		completedGrid = Arrange(grid, pieces)
		if completedGrid != nil {
			break
		}
		gridSize++
	}

	Display(completedGrid)
}
