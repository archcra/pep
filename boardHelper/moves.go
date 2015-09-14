package boardHelper

//"fmt"
// "math"

type Pos struct {
	row int
	col int
}

type MoveInfo struct {
	weight     int
	piece      int
	pieceColor int // 1 for RED, -1 for BLACK
	moves      []Pos
}

var movesWeight map[int]int

var palaceUpperRow map[int]int
var palaceLowerRow map[int]int

var posOffsetsAdvisor [][]int
var posOffsetsGeneral [][]int
var posOffsetsHorse [][]int

var pieceFuns map[int]func(board [13][12]int, row int, column int, isRedFlag int) []Pos

func init() {

	movesWeight = make(map[int]int)

	movesWeight[0] = 0
	movesWeight[ADVISOR] = 1
	movesWeight[ELEPHANT] = 1
	movesWeight[SOLDIER] = 2
	movesWeight[HORSE] = 5
	movesWeight[CANNON] = 5
	movesWeight[ROOK] = 10
	movesWeight[GENERAL] = 1
	movesWeight[ADVISOR+COLOR] = 1
	movesWeight[ELEPHANT+COLOR] = 1
	movesWeight[SOLDIER+COLOR] = 2
	movesWeight[HORSE+COLOR] = 5
	movesWeight[CANNON+COLOR] = 5
	movesWeight[ROOK+COLOR] = 10
	movesWeight[GENERAL+COLOR] = 1

	pieceFuns = make(map[int]func(board [13][12]int, row int, column int, isRedFlag int) []Pos)
	pieceFuns = map[int]func(board [13][12]int, row int, column int, isRedFlag int) []Pos{
		ADVISOR:          advisorMoves, // lower case init for local fun
		ADVISOR + COLOR:  advisorMoves,
		ELEPHANT:         elephantMoves,
		ELEPHANT + COLOR: elephantMoves,
		GENERAL:          generalMoves,
		GENERAL + COLOR:  generalMoves,
		SOLDIER:          soldierMoves,
		SOLDIER + COLOR:  soldierMoves,
		HORSE:            horseMoves,
		HORSE + COLOR:    horseMoves,
		ROOK:             rookMoves,
		ROOK + COLOR:     rookMoves,
		CANNON:           cannonMoves,
		CANNON + COLOR:   cannonMoves,
	}

	palaceUpperRow = map[int]int{1: 8, -1: 1}
	palaceLowerRow = map[int]int{1: 10, -1: 3}

	posOffsetsAdvisor = [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	posOffsetsGeneral = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	posOffsetsHorse = [][]int{{-2, -1}, {-1, 2}, {-1, -2}, {1, -2}, {-2, 1}, {2, 1}, {1, 2}, {2, -1}}

}

const (
	PALACE_LEFT  = 4
	PALACE_RIGHT = 6
)

func validPos(board [13][12]int, row int, column int, isRedFlag int) bool {
	result := false
	if row < 0 || row > 10 || column < 0 || column > 9 {
		return false
	}
	if board[row][column] == 0 ||
		(board[row][column] > 0 && isRedFlag*(COLOR-board[row][column]) < 0) {
		//Empty or the other side, a taken)){
		result = true
	}

	return result
}

func advisorMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos

	for i := 0; i < 4; i++ {
		if (row+posOffsetsAdvisor[i][0] >= palaceUpperRow[isRedFlag]) && (row+posOffsetsAdvisor[i][0] <= palaceLowerRow[isRedFlag]) && (column+posOffsetsAdvisor[i][1] >= PALACE_LEFT) && (column+posOffsetsAdvisor[i][1] <= PALACE_RIGHT) &&
			validPos(board, row+posOffsetsAdvisor[i][0], column+posOffsetsAdvisor[i][1], isRedFlag) {
			pos = append(pos, Pos{row + posOffsetsAdvisor[i][0], column + posOffsetsAdvisor[i][1]})
		}
	}

	return pos
}

func elephantMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos
	elephantUpperRow := map[int]int{1: 6, -1: 1}
	elephantLowerRow := map[int]int{1: 10, -1: 5}
	posOffsetsElephant := [][]int{{-2, -2}, {-2, 2}, {2, 2}, {2, -2}}

	for i := 0; i < 4; i++ {
		if (row+posOffsetsElephant[i][0] >= elephantUpperRow[isRedFlag]) && (row+posOffsetsElephant[i][0] <= elephantLowerRow[isRedFlag]) &&
			validPos(board, row+posOffsetsElephant[i][0], column+posOffsetsElephant[i][1], isRedFlag) {

			//console.log('leg pos should have no piece: column-1, piece', row + posOffsetsElephant[i][0], column + posOffsetsElephant[i][1], row + ~~(posOffsetsElephant[i][0] / 2), column + ~~(posOffsetsElephant[i][1] / 2), board[row + ~~(posOffsetsElephant[i][0] / 2)][column + ~~(posOffsetsElephant[i][1] / 2)]);

			pos = append(pos, Pos{row + posOffsetsElephant[i][0], column + posOffsetsElephant[i][1]})

		}
	}
	return pos
}

func generalMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	//　Almost same as Advisor's code except face 2 dace
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos

	for i := 0; i < 4; i++ {
		if (row+posOffsetsGeneral[i][0] >= palaceUpperRow[isRedFlag]) && (row+posOffsetsGeneral[i][0] <= palaceLowerRow[isRedFlag]) && (column+posOffsetsGeneral[i][1] >= PALACE_LEFT) && (column+posOffsetsGeneral[i][1] <= PALACE_RIGHT) &&
			validPos(board, row+posOffsetsGeneral[i][0], column+posOffsetsGeneral[i][1], isRedFlag) {
			pos = append(pos, Pos{row + posOffsetsGeneral[i][0], column + posOffsetsGeneral[i][1]})
		}
	}

	// Face2face
	sameColumn := false
	oppositeGeneralRow := -1

	for i := 1; i <= 3; i++ {
		if board[i+7*(1-isRedFlag)/2][column] == GENERAL+
			COLOR*(1+isRedFlag)/2 {
			sameColumn = true
			oppositeGeneralRow = i + 7*(1-isRedFlag)/2
		}
	}
	if sameColumn {
		// No pieces on this column
		var rowFrom, rowTo int
		if isRedFlag == 1 {
			rowFrom = oppositeGeneralRow
			rowTo = row
		} else {
			rowFrom = row
			rowTo = oppositeGeneralRow
		}

		face2face := true
		//console.log('rowFrom, rowTo are: ', rowFrom, rowTo);
		for i := rowFrom + 1; i < rowTo; i++ {
			//console.log('i, board[i][column]', i , board[i][column]);
			if board[i][column] > 0 {
				face2face = false
			}
		}
		if face2face { // Add opposite general's position
			pos = append(pos, Pos{oppositeGeneralRow, column})
		}
	}
	return pos
}

func soldierMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos

	newRow := row - isRedFlag
	if validPos(board, newRow, column, isRedFlag) {
		pos = append(pos, Pos{newRow, column})
	}

	if isRedFlag*(5-row) >= (1-isRedFlag)/2 {
		for _, columnOffset := range []int{-1, 1} {
			newColumn := column + columnOffset
			if validPos(board, row, newColumn, isRedFlag) {
				pos = append(pos, Pos{row, newColumn})
			}
		}
	}
	return pos
}

func horseMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos
	//console.log('isRed: ', isRedFlag);
	for i := 0; i < 8; i++ {
		if board[row+posOffsetsHorse[i][0]/2][column+posOffsetsHorse[i][1]/2] == 0 &&
			validPos(board, row+posOffsetsHorse[i][0], column+posOffsetsHorse[i][1], isRedFlag) {

			//console.log('leg pos should have no piece: column-1, piece', row + posOffsetsHorse[i][0], column + posOffsetsHorse[i][1], row + ~~(posOffsetsHorse[i][0] / 2), column + ~~(posOffsetsHorse[i][1] / 2), board[row + ~~(posOffsetsHorse[i][0] / 2)][column + ~~(posOffsetsHorse[i][1] / 2)]);
			pos = append(pos, Pos{row + posOffsetsHorse[i][0], column + posOffsetsHorse[i][1]})

		}
	}

	return pos

}

func rookMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	// Get all possible moves of a piece as rook move at position row:column.

	//console.log("in rookMoves, isRedFlag is: ", isRedFlag);
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos

	// 1. collect all possible locations (a big cross)
	// 2. remove it's own location
	// 3. loop row, if < pieceRow, has a
	// Keep original algorithm for performance consideration.

	var rowIndex, columnIndex int
	searchValid := [2]bool{true, true}
	direction := 0

	for i := row + 1; i <= 10+row; i++ {
		if i <= 10 {
			rowIndex = i
		} else {
			rowIndex = row - i + 10
			direction = 1
		}

		if !searchValid[direction] {
			continue
		}

		//console.log('board is: ', board);
		//console.log('rowIndex is: ', rowIndex);
		switch board[rowIndex][column] {
		case -1:
			continue

		case 0:
			pos = append(pos, Pos{rowIndex, column})
		default:
			searchValid[direction] = false
			if isRedFlag*(128-board[rowIndex][column]) < 0 {
				// The other side, a taken
				pos = append(pos, Pos{rowIndex, column})
			}
		}
	}

	direction = 0                     //Reset
	searchValid = [2]bool{true, true} // Reset

	for i := column + 1; i <= 9+column; i++ {
		if i <= 9 {
			columnIndex = i
		} else {
			columnIndex = column - i + 9
			direction = 1
		}

		if !searchValid[direction] {
			continue
		}

		switch board[row][columnIndex] {
		case -1:
			continue

		case 0:
			pos = append(pos, Pos{row, columnIndex})
		default:
			searchValid[direction] = false
			if isRedFlag*(128-board[row][columnIndex]) < 0 {
				// The other side, a taken
				pos = append(pos, Pos{row, columnIndex})
			}
		}
	}

	return pos
}

func cannonMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos
	var rowIndex, columnIndex int
	attackModes := [2]bool{false, false}
	searchValid := [2]bool{true, true}
	direction := 0

	//console.log('isRed: ', isRedFlag);
	for i := row + 1; i <= 10+row; i++ {
		if i <= 10 {
			rowIndex = i
		} else {
			direction = 1
			rowIndex = row - i + 10
		}

		switch true {
		case !searchValid[direction]:
		case !attackModes[direction] && board[rowIndex][column] == 0:
			pos = append(pos, Pos{rowIndex, column})
		case attackModes[direction] && board[rowIndex][column] > 0 && isRedFlag*(128-board[rowIndex][column]) < 0:
			pos = append(pos, Pos{rowIndex, column})

			searchValid[direction] = false
		case attackModes[direction] && board[rowIndex][column] > 0 && isRedFlag*(128-board[rowIndex][column]) > 0:
			searchValid[direction] = false
		case !attackModes[direction] && board[rowIndex][column] > 0:
			attackModes[direction] = true
		default:
		}
	}

	/* */
	attackModes = [2]bool{false, false} //Reset
	direction = 0                       //Reset
	searchValid = [2]bool{true, true}   // Reset

	for i := column + 1; i <= 9+column; i++ {
		if i <= 9 {
			columnIndex = i
		} else {
			direction = 1
			columnIndex = column - i + 9
		}

		switch true { // same as above
		case !searchValid[direction]:

		case !attackModes[direction] && board[row][columnIndex] == 0:
			pos = append(pos, Pos{row, columnIndex})

		case attackModes[direction] && board[row][columnIndex] > 0 && isRedFlag*(128-board[row][columnIndex]) > 0:
			searchValid[direction] = false
		case attackModes[direction] && board[row][columnIndex] > 0 && isRedFlag*(128-board[row][columnIndex]) < 0:
			pos = append(pos, Pos{row, columnIndex})

			searchValid[direction] = false
		case !attackModes[direction] && board[row][columnIndex] > 0:
			attackModes[direction] = true
		default:
		}
	}
	return pos
}

func getPossibleMovesInfo(board [13][12]int, row int, column int) MoveInfo {
	var movesInfo MoveInfo

	movesInfo.weight = movesWeight[board[row][column]]

	movesInfo.piece = board[row][column]
	if movesInfo.piece > COLOR {
		movesInfo.pieceColor = -1
	} else {
		movesInfo.pieceColor = 1
	}
	pieceFun := pieceFuns[movesInfo.piece]
	if pieceFun == nil {
		//fmt.Printf("No valid piece of value %d on this position %d:%d of board %q.", movesInfo.piece, row, column, board)
		return movesInfo
	}
	movesInfo.moves = pieceFun(board, row, column, movesInfo.pieceColor)

	return movesInfo
}
