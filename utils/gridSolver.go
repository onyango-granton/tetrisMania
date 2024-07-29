package utils

/*
canPlace Checks if a Tetromino can be placed at a specified position (row, col) on the grid without
overlapping other Tetrominoes or going out of the grid boundaries.
*/
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

/*place function Places a Tetromino at a specified position (row, col) on the grid.
The Tetromino's name will replace the * in the grid to mark its placement.*/
func place(term Tetromino, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = term.name
			}
		}
	}
}

/*remove func Removes a Tetromino from a specified position (row, col) on the grid, 
replacing its name with * to indicate an empty cell.*/
func remove(term Tetromino, grid [][]string, row, col int) {
	for r := range term.shape {
		for c := range term.shape[r] {
			if term.shape[r][c] == 1 {
				grid[row+r][col+c] = "*"
			}
		}
	}
}

/*completeGrid function Attempts to place all Tetrominoes in tetro_group onto the grid using recursive backtracking. 
Starts from the indexth Tetromino and tries to place each one in every possible position on the grid.*/
func CompleteGrid(tetro_group []Tetromino, grid [][]string, index int) bool {
	if index == len(tetro_group) {
		return true
	}
	for row := range grid {
		for col := range grid[row] {
			if canPlace(tetro_group[index], grid, row, col) {
				place(tetro_group[index], grid, row, col)
				if CompleteGrid(tetro_group, grid, index+1) {
					return true
				}
				remove(tetro_group[index], grid, row, col)
			}
		}
	}
	return false
}
