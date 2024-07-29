Tetromino Placement in a Grid

This Go program is designed to read a text file containing tetromino shapes, validate them, and then attempt to place them into a square grid using a recursive backtracking algorithm. The program ensures that each tetromino is placed without overlapping and fits within the bounds of the grid.
Table of Contents

    Structures
    Functions
        byteToInt
        stringToIntSlice
        allOne
        sliceIsEmpty
        isSurroundedByOnes
        isValidTetro
        tetroGroupFunc
        initGrid
        canPlace
        place
        remove
        completeGrid
        printGrid
        trimTetrominoListFunc
        trimTetromino
    Main Function

Structures
Tetromino

type Tetromino struct {
	shape [][]int
	name  string
}

    shape: A 2D slice representing the shape of the tetromino, where 1 indicates a block and 0 indicates an empty space.
    name: A string representing the name or identifier of the tetromino.

Functions
byteToInt

func byteToInt(b byte) (int, error)

Converts a byte to an integer.

    Parameters:
        b: A byte representing either a block (#) or an empty space (.).
    Returns:
        int: 1 if the byte is #, 0 if the byte is ..
        error: An error if the byte is neither # nor ..

stringToIntSlice

func stringToIntSlice(s string) ([]int, error)

Converts a string of length 4 to a slice of integers.

    Parameters:
        s: A string of length 4 containing characters # and ..
    Returns:
        []int: A slice of integers representing the converted string.
        error: An error if the string length is not 4 or contains invalid characters.

allOne

func allOne(num1, num2 int) bool

Checks if both integers are 1.

    Parameters:
        num1: An integer.
        num2: An integer.
    Returns:
        bool: true if both num1 and num2 are 1, otherwise false.

sliceIsEmpty

func sliceIsEmpty(num []int) bool

Checks if a slice of integers is empty (all zeros).

    Parameters:
        num: A slice of integers.
    Returns:
        bool: true if all elements in the slice are 0, otherwise false.

isSurroundedByOnes

func isSurroundedByOnes(arr [][]int, row, col int) bool

Checks if a block at a specific position in a 2D array is surrounded by other blocks.

    Parameters:
        arr: A 2D slice of integers.
        row: The row index.
        col: The column index.
    Returns:
        bool: true if the block is surrounded by other blocks, otherwise false.

isValidTetro

func isValidTetro(tetro [][]int) (bool, error)

Validates if a tetromino shape is valid.

    Parameters:
        tetro: A 2D slice of integers representing the tetromino shape.
    Returns:
        bool: true if the tetromino is valid, otherwise false.
        error: An error if the tetromino is invalid.

tetroGroupFunc

func tetroGroupFunc(textFile string) ([]Tetromino, int)

Reads a text file, processes the tetromino shapes, and returns a list of valid tetrominos and the grid size.

    Parameters:
        textFile: The path to the text file containing tetromino shapes.
    Returns:
        []Tetromino: A list of valid tetrominos.
        int: The size of the grid.

initGrid

func initGrid()

Initializes a square grid with a size determined by the number of tetrominos.
canPlace

func canPlace(term Tetromino, grid [][]string, row, col int) bool

Checks if a tetromino can be placed at a specific position in the grid.

    Parameters:
        term: The tetromino to be placed.
        grid: The grid where the tetromino is to be placed.
        row: The starting row index.
        col: The starting column index.
    Returns:
        bool: true if the tetromino can be placed, otherwise false.

place

func place(term Tetromino, grid [][]string, row, col int)

Places a tetromino at a specific position in the grid.

    Parameters:
        term: The tetromino to be placed.
        grid: The grid where the tetromino is to be placed.
        row: The starting row index.
        col: The starting column index.

remove

func remove(term Tetromino, grid [][]string, row, col int)

Removes a tetromino from a specific position in the grid.

    Parameters:
        term: The tetromino to be removed.
        grid: The grid where the tetromino is to be removed.
        row: The starting row index.
        col: The starting column index.

completeGrid

func completeGrid(tetro_group []Tetromino, grid [][]string, index int) bool

Uses recursive backtracking to place all tetrominos in the grid.

    Parameters:
        tetro_group: The list of tetrominos to be placed.
        grid: The grid where the tetrominos are to be placed.
        index: The current index in the tetromino list.
    Returns:
        bool: true if all tetrominos are successfully placed, otherwise false.

printGrid

func printGrid()

Prints the current state of the grid.
trimTetrominoListFunc

func trimTetrominoListFunc(tetrominoList []Tetromino) []Tetromino

Trims the tetromino shapes to remove unnecessary empty rows and columns.

    Parameters:
        tetrominoList: The list of tetrominos to be trimmed.
    Returns:
        []Tetromino: The trimmed list of tetrominos.

trimTetromino

func trimTetromino(tetro [][]int) [][]int

Trims a single tetromino shape to remove unnecessary empty rows and columns.

    Parameters:
        tetro: The tetromino shape to be trimmed.
    Returns:
        [][]int: The trimmed tetromino shape.

Main Function

func main()

The main function initializes the grid, attempts to place all tetrominos, and prints the result. If no solution is found, it prints "No solutions found".This Go program provides a comprehensive solution for placing tetromino shapes in a square grid. It includes functions to read and process tetromino shapes, validate them, and use a backtracking algorithm to place them in the grid without overlapping. The program ensures that each tetromino fits within the grid boundaries and is surrounded by other blocks.