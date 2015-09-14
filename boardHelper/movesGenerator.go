package boardHelper

import (
	"bytes"
	"strconv"
)

type MoveResult struct {
	Move  string
	Board [13][12]int
}

func Generate(board [13][12]int, roundColor int) []MoveResult {
	// roundColor = 1 if red, roundColor = -1 if black
	//   // move should be in the form of "a8-87"(ma er jin san)
	var moves []MoveResult
	var changedBoard [13][12]int
	var movesInfo MoveInfo
	//  mobility = 0,
	//  attack = 0;
	for i := 1; i <= BOARD_ROWS; i++ {
		for j := 1; j <= BOARD_COLUMNS; j++ {
			if board[i][j] != 0 &&
				roundColor*(COLOR-board[i][j]) > 0 {
				// This this is the peice of the color which in turn to move

				movesInfo = getPossibleMovesInfo(board, i, j)
				if len(movesInfo.moves) > 0 {
					// If no possible moves, then do not push
					for k := 0; k < len(movesInfo.moves); k++ {
						changedBoard = board                                                       // Is this deep copy? for there is pointer, this should be
						changedBoard[movesInfo.moves[k].row][movesInfo.moves[k].col] = board[i][j] //move the piece
						changedBoard[i][j] = 0
						// clear the original piece
						buffer := make([]byte, 10) // Max: 10:9-10:8
						bl := 0
						bl += copy(buffer[bl:], strconv.Itoa(i))
						bl += copy(buffer[bl:], ":")
						bl += copy(buffer[bl:], strconv.Itoa(j))
						bl += copy(buffer[bl:], "-")
						bl += copy(buffer[bl:], strconv.Itoa(movesInfo.moves[k].row))
						bl += copy(buffer[bl:], ":")
						bl += copy(buffer[bl:], strconv.Itoa(movesInfo.moves[k].col))
						buffer = bytes.Trim(buffer, "\x00") // 否则字符串长度不对
						moves = append(moves, MoveResult{string(buffer), changedBoard})
					}
				}
			}
		}
	}
	return moves
}
