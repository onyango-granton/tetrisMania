package utils

/*trimTetrominoListFunc function takes a list of tetrominos and trims each tetromino in the list by removing empty rows and columns.*/
func trimTetrominoListFunc(tetrominoList []Tetromino) []Tetromino {
	for tetromino := range tetrominoList {
		tetrominoList[tetromino].shape = trimTetromino(tetrominoList[tetromino].shape)
	}
	return tetrominoList
}

/*trimTetromino trims a single tetromino's shape by removing empty rows and columns. It specifically handles tetrominos of dimensions 3xN and 2xN.*/
func trimTetromino(tetro [][]int) [][]int {
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

}
