package boardHelper

import (
	"fmt"
	"math"
)

var matablePiece = map[string]int{
	"r": COLOR + ROOK,
	"h": COLOR + HORSE,
	"c": COLOR + CANNON,
	"s": COLOR + SOLDIER,
	"R": ROOK,
	"H": HORSE,
	"C": CANNON,
	"S": SOLDIER,
}

func countMateScore(board [13][12]int) int {
	// Calculate both side, BLACK as negative
	redGeneralPos := GetGeneralPos(board, RED)
	blackGeneralPos := GetGeneralPos(board, BLACK)

	/* TODO
	if blackGeneralPos == nil || redGeneralPos == nil {
		return 0 // Missed a general.
	}
	*/

	var piece, score, distance, cannonBonusFactor int
	for i := 1; i <= BOARD_ROWS; i++ {
		for j := 1; j <= BOARD_COLUMNS; j++ {
			piece = board[i][j]
			// fmt.Printf("Now piece is %d and i;j is:%d;%d", piece, i, j)
			thisPieceStr := pieceStr[piece]
			if matablePiece[thisPieceStr] == 0 {
				continue // Skip this piece
			}
			if piece < COLOR {
				distance = int(math.Abs(float64(i-blackGeneralPos.row)) +
					math.Abs(float64(j-blackGeneralPos.col)))
			} else {
				distance = int(math.Abs(float64(i-redGeneralPos.row)) +
					math.Abs(float64(j-redGeneralPos.col)))
			}

			switch true {
			case thisPieceStr == "C":
				cannonBonusFactor = getCannonBous(board, Pos{i, j}, blackGeneralPos)
				distance = 1 // distance is irrelevant

			case thisPieceStr == "c":
				cannonBonusFactor = getCannonBous(board, Pos{i, j}, redGeneralPos)
				distance = 1 // TODO

			default:
				cannonBonusFactor = 1
			}

			pieceValue := piece
			if pieceValue > COLOR {
				pieceValue = -(pieceValue - COLOR)
			}

			if cannonBonusFactor == 0 {
				fmt.Println("cannonBonusFactor is ZERO???") // TODO
			}
			score = score + pieceValue/distance*cannonBonusFactor
		}
	}

	return score // Distance near is good.
}

func getCannonBous(board [13][12]int, fromPos Pos, toPos Pos) int {
	bonusFactor := obstaclesCount(board, fromPos, toPos)

	if fromPos.row == toPos.row || fromPos.col == toPos.col {
		bonusFactor = bonusFactor + 2
	}
	return bonusFactor
}

func obstaclesCount(board [13][12]int, posFrom Pos, posTo Pos) int {
	// TODO
	return 1
}
