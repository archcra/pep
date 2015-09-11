package boardHelper

import (
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
	var redGeneralPos, blackGeneralPos Pos
	for i := 1; i <= 3; i++ {
		for j := 4; j <= 6; j++ {
			//console.log("possible black general piece is: ", board[i][ j]);
			if board[i][j] == 254 {
				blackGeneralPos = Pos{i, j}
			}
			if board[i+7][j] == 126 {
				redGeneralPos = Pos{i + 7, j}
			}
		}
	}

	/* TODO
	if blackGeneralPos == nil || redGeneralPos == nil {
		return 0 // Missed a general.
	}
	*/

	var piece, score, distance, cannonBonusFactor int
	for i := 1; i <= BOARD_ROWS; i++ {
		for j := 1; j <= BOARD_COLUMNS; j++ {
			piece = board[i][j]

			thisPieceStr := pieceStr[piece]
			if matablePiece[thisPieceStr] == 0 {
				continue // Skip this piece
			}
			if piece < COLOR {
				distance = int(math.Abs(float64(i - blackGeneralPos.row + j - blackGeneralPos.col)))
			} else {
				distance = int(math.Abs(float64(i - redGeneralPos.row + j - redGeneralPos.col)))
			}

			switch true {
			case thisPieceStr == "C":
				cannonBonusFactor = getCannonBous(board, Pos{i, j}, blackGeneralPos)
				distance = 1 // distance is irrelevant
				//console.log('cannonBonusFactor is', cannonBonusFactor);

			case thisPieceStr == "c":
				cannonBonusFactor = getCannonBous(board, Pos{i, j}, redGeneralPos)
				distance = 1
				//console.log('cannonBonusFactor is', cannonBonusFactor);

			default:
				cannonBonusFactor = 1
			}

			pieceValue := piece
			if pieceValue > COLOR {
				pieceValue = -(pieceValue - COLOR)
			}
			score = score + pieceValue/distance*cannonBonusFactor
			//console.log('piece, mateWeight, distance, score', piece, PIECE_MATE_WEIGHT[piece], distance, PIECE_MATE_WEIGHT[piece] * cannonBonusFactor / distance);
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
