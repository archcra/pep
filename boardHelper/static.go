package boardHelper

func GetGeneralPos(board [13][12]int, color int) Pos {
	// color 1: RED, -1: BLACK

	var generalPos Pos

	for i := 1; i <= 3; i++ {
		for j := 4; j <= 6; j++ {
			row := i + 7*(1+color)/2
			if board[row][j] == GENERAL+LIMBO*(1-color)/2 {
				generalPos = Pos{row, j}
			}
		}
	}
	return generalPos
}

func GeneralBeTaken(board [13][12]int, roundColor int) bool {
	return GetGeneralPos(board, -roundColor).row == 0
}
