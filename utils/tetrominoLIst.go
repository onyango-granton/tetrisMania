package utils

import (
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

/*sliceIsEmpty Checks if a slice of integers is empty, 
meaning all elements are zero.*/
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

/*TetroGroupFunc Reads a text file containing Tetromino shapes, processes them, 
and returns a list of valid Tetrominoes along with the grid size required to accommodate them.*/
func TetroGroupFunc(textFile string) ([]Tetromino, int) {
	tetrominoesGroup := []Tetromino{}
	// opens text file
	sampleFile, err := os.ReadFile(textFile)
	if runtime.GOOS == "windows" {
		for i, ch := range sampleFile {
			if i+1 < len(sampleFile) && ch == byte(rune(13)) {
				// fmt.Print("here")
				sampleFile = append(sampleFile[:i], sampleFile[i+1:]...)
			} else if i+1 < len(sampleFile) && ch == byte(rune(13)) {
				sampleFile = sampleFile[:i]
			}
		}
	}
	// fmt.Println(sampleFile,234)
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
