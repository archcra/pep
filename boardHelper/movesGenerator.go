package boardHelper

import (
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
	var move string
	var movesInfo MoveInfo
	//  mobility = 0,
	//  attack = 0;
	for i := 1; i <= BOARD_ROWS; i++ {
		for j := 1; j <= BOARD_COLUMNS; j++ {
			if board[i][j] != 0 &&
				roundColor*(COLOR-board[i][j]) > 0 {
				// This this is the peice of the color which in turn to move
				//console.log("should get moves of position: x,y,isRedFlag,  COLOR - board[i][j]: ", i, j, isRedFlag, COLOR - board[i][j]);

				movesInfo = getPossibleMovesInfo(board, i, j)
				//mobility = mobility + movesInfo.weight*len(movesInfo.moves)
				//console.log("In generate, position : board, i,j, moves",board,  i,j ,movesInfo);
				if len(movesInfo.moves) > 0 {
					//console.log("movesInfo is: ", movesInfo);
					// If no possible moves, then do not push
					for k := 0; k < len(movesInfo.moves); k++ {
						changedBoard = board                                                       // Is this deep copy? for there is pointer, this should be
						changedBoard[movesInfo.moves[k].row][movesInfo.moves[k].col] = board[i][j] //move the piece
						changedBoard[i][j] = 0                                                     // clear the original piece
						move = strconv.Itoa(i) + ":" + strconv.Itoa(j) + "-" + strconv.Itoa(movesInfo.moves[k].row) + ":" + strconv.Itoa(movesInfo.moves[k].col)
						//attack = attack + changedBoard

						moves = append(moves, MoveResult{move, changedBoard})
					}
				}
			}
		}
	}
	return moves
}
