package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const tetrisBlockSize = 4

// ReadTetrisBlocks reads tetris blocks from a text file and returns them as a slice of 2D string slices.
func ReadTetrisBlocks(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var blocks [][]string
	var currentBlock []string
	var letter rune = 'A'

	for _, line := range lines {
		if line == "" {
			blocks = append(blocks, currentBlock)
			currentBlock = nil
			letter++
			continue
		}

		newLine, err := transformLine(line, letter)
		if err != nil {
			return nil, err
		}
		currentBlock = append(currentBlock, newLine)
	}

	if len(currentBlock) > 0 {
		blocks = append(blocks, currentBlock)
	}

	return blocks, nil
}

func transformLine(line string, letter rune) (string, error) {
	var builder strings.Builder
	for _, char := range line {
		switch char {
		case '#':
			builder.WriteRune(letter)
		case '.':
			builder.WriteRune(char)
		default:
			return "", fmt.Errorf("invalid character in tetris block: %c", char)
		}
	}
	return builder.String(), nil
}

// CreateBoard creates a 2D slice of strings representing the game board.
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

// PlaceTetrisBlock places a tetris block on the board at the given position.
func PlaceTetrisBlock(board [][]string, block []string, x, y int) {
	for dy, row := range block {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

// RemoveTetrisBlock removes a tetris block from the board at the given position.
func RemoveTetrisBlock(board [][]string, block []string, x, y int) {
	for dy, row := range block {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}

// CanPlaceTetrisBlock checks if a tetris block can be placed on the board at the given position.
func CanPlaceTetrisBlock(board [][]string, block []string, x, y int) bool {
	for dy, row := range block {
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

// SolveTetris attempts to place all tetris blocks on the board using a recursive approach.
func SolveTetris(board [][]string, blocks [][]string) [][]string {
	if placeTetrisBlocks(board, blocks, 0) {
		return board
	}
	return nil
}

func placeTetrisBlocks(board [][]string, blocks [][]string, index int) bool {
	if index == len(blocks) {
		return true
	}

	block := blocks[index]
	for y := range board {
		for x := range board[y] {
			if CanPlaceTetrisBlock(board, block, x, y) {
				PlaceTetrisBlock(board, block, x, y)
				if placeTetrisBlocks(board, blocks, index+1) {
					return true
				}
				RemoveTetrisBlock(board, block, x, y)
			}
		}
	}
	return false
}

// TrimTetrisBlocks trims the empty space around each tetris block.
func TrimTetrisBlocks(blocks [][]string) [][]string {
	var trimmedBlocks [][]string
	for _, block := range blocks {
		trimmedBlocks = append(trimmedBlocks, trimBlock(block))
	}
	return trimmedBlocks
}

func trimBlock(block []string) []string {
	var result []string
	columnHasLetter := make([]bool, len(block[0]))

	for col := range columnHasLetter {
		for row := range block {
			if block[row][col] != '.' {
				columnHasLetter[col] = true
				break
			}
		}
	}

	for _, row := range block {
		var newRow strings.Builder
		for col, hasLetter := range columnHasLetter {
			if hasLetter {
				newRow.WriteByte(row[col])
			}
		}
		result = append(result, newRow.String())
	}
	return result
}

// ValidateTetrisBlocks checks if the input tetris blocks are valid.
func ValidateTetrisBlocks(blocks [][]string) error {
	if len(blocks) > 26 {
		return errors.New("invalid file: too many tetris blocks")
	}
	for _, block := range blocks {
		if len(block) != tetrisBlockSize {
			return errors.New("invalid file: tetris block size is not 4x4")
		}
		if err := validateConnections(block); err != nil {
			return err
		}
	}
	return nil
}

func validateConnections(block []string) error {
	connectionCount, letterCount := 0, 0
	for i, row := range block {
		for j, char := range row {
			if char != '.' {
				letterCount++
				connectionCount += countAdjacentConnections(block, i, j, char)
			}
		}
	}

	if connectionCount < 6 || letterCount != 4 {
		return errors.New("invalid file: tetris block is not valid")
	}
	return nil
}

func countAdjacentConnections(block []string, i, j int, char rune) int {
	connections := 0
	if i > 0 && block[i-1][j] == byte(char) {
		connections++
	}
	if i < len(block)-1 && block[i+1][j] == byte(char) {
		connections++
	}
	if j > 0 && block[i][j-1] == byte(char) {
		connections++
	}
	if j < len(block[i])-1 && block[i][j+1] == byte(char) {
		connections++
	}
	return connections
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: Invalid number of arguments")
	}

	blocks, err := ReadTetrisBlocks(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if err := ValidateTetrisBlocks(blocks); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	blocks = TrimTetrisBlocks(blocks)

	size := int(math.Ceil(math.Sqrt(float64(len(blocks) * tetrisBlockSize))))
	var finalBoard [][]string
	for {
		board := CreateBoard(size)
		finalBoard = SolveTetris(board, blocks)
		if finalBoard != nil {
			break
		}
		size++
	}

	printBoard(finalBoard)
}

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}
