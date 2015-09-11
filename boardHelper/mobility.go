package boardHelper

func countMobility(board [13][12]int) int {
	// From RED's viewpoint
	var mobility int

	for i := 1; i <= BOARD_ROWS; i++ {
		for j := 1; j <= BOARD_COLUMNS; j++ {
			piece := board[i][j]
			if piece <= 0 {
				continue // No piece here, continue
			}
			movesInfo := getPossibleMovesInfo(board, i, j)

			if piece < COLOR {
				mobility = mobility + movesInfo.weight*len(movesInfo.moves)
			} else {
				mobility = mobility - movesInfo.weight*len(movesInfo.moves)
			}
		}
	}
	return mobility
}
