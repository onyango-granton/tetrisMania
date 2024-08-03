package utils

/*InitGrid function initializes a square grid of a specified size with each cell initially set to the string "*".
This function is useful for creating a grid structure for games, simulations, or any application requiring a two-dimensional array initialized to a default value.*/
func InitGrid(gridSize int) [][]string {
	grid := make([][]string, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]string, gridSize)
		for j := 0; j < gridSize; j++ {
			grid[i][j] = "*"
		}
	}
	return grid
}
