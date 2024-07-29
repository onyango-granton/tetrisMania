package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
)

type Tetromino struct {
	shape [][]int
	name  string
}

func byteToInt(b byte) (int, error) {
	if b == '.' {
		return 0, nil
	} else if b == '#' {
		return 1, nil
	} else {
		fmt.Println(os.Getwd())
		return 0, errors.New("error in sample file")
	}
}

// func isValidTetro(tetro [][]int) bool {
// 	for row := range tetro{
// 		for col := range tetro[row]{
// 			if tetro[row][col] == 1 && tetro
// 		}
// 	}
// }

func stringToIntSlice(s string) ([]int, error) {
	//file,_ := os.ReadFile("sample.txt")
	//fmt.Println(file)
	res := []int{}
	if len(s) != 4 {
		fmt.Println(s)
		return nil, errors.New("invalid length entry in file")
	}
	for _, b := range s {
		num, err := byteToInt(byte(b))
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}

func allOne(num1, num2 int) bool {
	if num1 == 1 {
		return num1 == num2
	}
	return false
}

func sliceIsEmpty(num []int) bool {
	var count int
	for i := range num {
		if num[i] == 0 {
			count++
		}
	}
	if count == 4 {
		return true
	} else {
		return false
	}
}

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

func fullyConnected(tetro [][]int) bool {
	//var tetro = [][]int{{0, 0, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}}
	//var tetro = [][]int{{1, 1, 1, 1}}
	connection := 0
	for row := range tetro {
		for col := range tetro[row] {
			if tetro[row][col] == 1 {
				if col+1 <= len(tetro[row])-1 && tetro[row][col+1] == 1 {
					/*if col == 0 {
						fmt.Print("first")
					}
					if col == 1 && row == 2 {
						fmt.Print("third")
					}
					fmt.Println("right connection")*/
					connection++
				}
				if col-1 >= 0 && tetro[row][col-1] == 1 {
					//if col == 1 && row == 1 {
					//	fmt.Print("second")
					//}
					//fmt.Println("left connection")
					connection++
				}
				if row+1 <= len(tetro)-1 && tetro[row+1][col] == 1 {
					//if col == 1 && row == 1 {
					//	fmt.Print("second")
					//}
					//fmt.Println("down connection")
					connection++
				}
				if row-1 >= 0 && tetro[row-1][col] == 1 {
					//if col == 1 && row == 2 {
					//	fmt.Print("third")
					//}
					//fmt.Println("up connection")
					connection++
				}
			}
		}
	}
	//fmt.Println(connection)
	if connection == 6 || connection == 8 {
		return true
	} else {
		return false
	}
}


func isValidTetro(tetro [][]int) (bool, error) {
	// fmt.Println(tetro)
	var bordercount int
	var linecount int

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {
			if tetro[row][col] == 1 {
				linecount++
			}
			if tetro[row][col] == 1 && isSurroundedByOnes(tetro, row, col) {
				bordercount++
				// fmt.Printf("Element at (%d, %d) is surrounded by ones\n", row, col)
			}
		}
	}

	for row := 0; row < len(tetro); row++ {
		for col := 0; col < len(tetro[row]); col++ {

		}
	}

	if bordercount > 4 || linecount > 4 {
		return false, errors.New("invalid Tetromino")
	} else {
		if fullyConnected(tetro){
			return true, nil
		} else {
			return false, errors.New("tetromino is invalid")
		}
		
	}
}

func tetroGroupFunc(textFile string) ([]Tetromino, int) {
	tetrominoesGroup := []Tetromino{}
	// opens text file
	sampleFile, err := os.ReadFile(textFile)
	if runtime.GOOS == "windows"{
		
		for i, ch := range sampleFile{
			if i+1 < len(sampleFile) && ch == byte(rune(13)){
				//fmt.Print("here")
				sampleFile = append(sampleFile[:i],sampleFile[i+1:]...)
			} else if i+1 < len(sampleFile) && ch == byte(rune(13)){
				sampleFile = sampleFile[:i]
			}
		}
	}
	//fmt.Println(sampleFile,234)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0
	}
	var nums [][]int
	for i, ch := range strings.Split(string(sampleFile), "\n") {
		if ch == "" {
			continue
		}
		chArr, err := stringToIntSlice(ch)
		if err != nil {
			fmt.Println(err.Error(), "at line", i+1)
			return nil, 0
		} else {
			nums = append(nums, chArr)
		}
	}

	startAscii := 'A'
	tetrominoes := make(map[rune][][]int)

	for i := 0; i < len(nums); {
		characterMino := [][]int{}
		for j := i; j < i+4; j++ {
			if sliceIsEmpty(nums[j]) {
				continue
			}
			characterMino = append(characterMino, nums[j])
		}
		tetrominoes[startAscii] = characterMino
		startAscii++
		i += 4
	}

	for k := range tetrominoes {
		res, err := isValidTetro(tetrominoes[k])
		if err != nil {
			fmt.Println(err.Error())
			return nil, 0
		} else if res {
			newTetro := Tetromino{shape: tetrominoes[k], name: string(k)}
			tetrominoesGroup = append(tetrominoesGroup, newTetro)
		}
	}

	gridSize := math.Sqrt(float64(len(tetrominoesGroup) * 4))

	return trimTetrominoListFunc(tetrominoesGroup), int(math.Ceil(gridSize))

	// for k,_ := range tetrominoes{
	// 	newTetro := Tetromino{shape: tetrominoes[k], name: string(k)}
	// 	tetrominoesGroup = append(tetrominoesGroup, newTetro)
	// }
}

var (
	tetroGroup, gridSize = tetroGroupFunc("tetris.txt")
	grid                 = make([][]string, gridSize)
)

// initializes a square grid 2D
func initGrid() {
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]string, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = "*"
		}
	}
}

// checks whether we can place the term at the specific row and col of grid
func canPlace(term Tetromino, grid [][]string, row, col int) bool {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				if row+r >= len(grid) || col+c >= len(grid[row]) || grid[row+r][col+c] != "*" {
					return false
				}
			}
		}
	}
	return true
}

// Places the term in row and col of the grid
func place(term Tetromino, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = term.name
			}
		}
	}
}

func remove(term Tetromino, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = "*"
			}
		}
	}
}

// fuction that uses recursive backtracking to place tetro in grig
func completeGrid(tetro_group []Tetromino, grid [][]string, index int) bool {
	if index == len(tetro_group) {
		return true
	}
	for row := range grid {
		for col := range grid[row] {
			if canPlace(tetro_group[index], grid, row, col) {
				place(tetro_group[index], grid, row, col)
				if completeGrid(tetro_group, grid, index+1) {
					return true
				}
				remove(tetro_group[index], grid, row, col)
			}
		}
	}
	return false
}

func printGrid() {
	for row := range grid {
		for col := range grid[row] {
			fmt.Print(grid[row][col] + " ")
		}
		fmt.Println()
	}
}

func main() {
	//if runtime.GOOS == "windows"{
	//	fmt.Println("Here we have windows")
	//}
	// fmt.Println(tetroGroupFunc("tetris.txt"))
	// fmt.Println(trimTetrominoListFunc())
	// initializes a 6 * 6 grid
	initGrid()
	// if grid is fitted display grid
	if completeGrid(tetroGroup, grid, 0) {
		printGrid()
	} else {
		fmt.Println("No solutions found")
	}
	//fullyConnected()
}

func trimTetrominoListFunc(tetrominoList []Tetromino) []Tetromino {
	// var result []Tetromino
	// tetrominoList, _ := tetroGroupFunc("tetris.txt")

	// fmt.Println(tetrominoList)

	for tetromino := range tetrominoList {
		tetrominoList[tetromino].shape = trimTetromino(tetrominoList[tetromino].shape)
	}
	return tetrominoList
}

func trimTetromino(tetro [][]int) [][]int {
	// var res = tetro

	// for row, _ := range tetro{
	// 	for col, _ := range tetro[row]{
	// 		fmt.Println(tetro)
	// 		if len(tetro) == 3 && tetro[row][col] == 0 && tetro[1][col] == 0 && tetro[2][col] == 0{
	// 			if col+1 < len(tetro[row]) - 1{
	// 				// tetro[row] = tetro[row][col+1:]
	// 				res[row] = tetro[row][col+1:]

	// 			} else {
	// 				// tetro[row] = tetro[row][:col]
	// 				res[row] = tetro[row][:col]
	// 			}
	// 		} else if len(tetro) == 2 && tetro[row][col] == 0 && tetro[1][col] == 0 {
	// 			// fmt.Println("Here")
	// 			// fmt.Println(tetro[row])
	// 			if col+1 < len(tetro[row]) - 1{
	// 				//tetro[row] = tetro[row][col+1:]
	// 				res[row] =  tetro[row][col+1:]
	// 			} else {
	// 				//tetro[row] = tetro[row][:col]
	// 				res[row] = tetro[row][:col]
	// 			}
	// 			// fmt.Println("He2re")
	// 		}
	// 	}
	// }

	// for row,_ := range tetro{
	// 	for col,_ := range tetro[row]{
	// 		fmt.Println(tetro)
	// 		if len(tetro) == 3 && tetro[row][col] == 0 && tetro[1][col] == 0 && tetro[2][col] == 0{
	// 			if col+1 < len(tetro[row]) - 1{
	// 				// tetro[row] = tetro[row][col+1:]
	// 				res[row] = tetro[row][col+1:]

	// 			} else {
	// 				// tetro[row] = tetro[row][:col]
	// 				res[row] = tetro[row][:col]
	// 			}
	// 		}
	// 	}
	// }

	for i := 0; i < 4; i++ {

		if len(tetro) == 3 {
			if tetro[0][0] == 0 && tetro[1][0] == 0 && tetro[2][0] == 0 {
				tetro[0] = tetro[0][1:]
				tetro[1] = tetro[1][1:]
				tetro[2] = tetro[2][1:]
			} else if tetro[0][len(tetro[0])-1] == 0 && tetro[1][len(tetro[1])-1] == 0 && tetro[2][len(tetro[2])-1] == 0 {
				tetro[0] = tetro[0][:len(tetro[0])-1]
				tetro[1] = tetro[1][:len(tetro[1])-1]
				tetro[2] = tetro[2][:len(tetro[2])-1]
			}
		}

		if len(tetro) == 2 {
			if tetro[0][0] == 0 && tetro[1][0] == 0 {
				tetro[0] = tetro[0][1:]
				tetro[1] = tetro[1][1:]
			} else if tetro[0][len(tetro[0])-1] == 0 && tetro[1][len(tetro[1])-1] == 0 {
				tetro[0] = tetro[0][:len(tetro[0])-1]
				tetro[1] = tetro[1][:len(tetro[1])-1]
			}
		}
	}

	return tetro

	// if len(tetro) == 3 && tetro[row][col] == 0 && tetro[1][col] == 0 && tetro[2][col] == 0{
	// 	if col+1 > len(tetro[row]) - 1{
	// 		tetro[row] = tetro[row][col+1:]
	// 	} else {
	// 		tetro[row] = tetro[row][:col]
	// 	}
	// } else if len(tetro) == 2 && tetro[row][col] == 0 && tetro[row+1][col] == 0 {
	// 	if col+1 > len(tetro[row]) - 1{
	// 		tetro[row] = tetro[row][col+1:]
	// 	} else {
	// 		tetro[row] = tetro[row][:col]
	// 	}
	// }

}
