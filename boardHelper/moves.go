package boardHelper

import (
    "fmt"
)

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
var pieceFuns map[int]func(board [13][12]int, row int, column int, isRedFlag int) []Pos

func init() {
	//posOffsetsGeneral := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//posOffsetsElephant := [][]int{{-2, -2}, {-2, 2}, {2, 2}, {2, -2}}
	//posOffsetsHorse := [][]int{{-2, -1}, {-1, 2}, {-1, -2}, {1, -2}, {-2, 1},{2, 1}, {1, 2}, {2, -1}, }

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
		ADVISOR: advisorMoves,
        ADVISOR+COLOR: advisorMoves,
		//GENERAL: generalMoves,
	}

}

const (
	PALACE_LEFT  = 4
	PALACE_RIGHT = 6
)

func validPos(board [13][12]int, row int, column int, isRedFlag int) bool {
	result := false
	//console.log('row, column is: ', row, column);
	if board[row][column] == 0 ||
		(board[row][column] > 0 && isRedFlag*(128-board[row][column]) < 0) {
		//Empty or the other side, a taken)){
		result = true
	}

	return result
}

func advisorMoves(board [13][12]int, row int, column int, isRedFlag int) []Pos {
    
    fmt.Println("...................................In move fun")
    
	// isRedFlag = 1 when red, -1 when black
	var pos []Pos
	advisorUpperRow := map[int]int{1: 8, -1: 1}
	advisorLowerRow := map[int]int{1: 10, -1: 3}
	posOffsetsAdvisor := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for i := 0; i < 4; i++ {
		if (row+posOffsetsAdvisor[i][0] >= advisorUpperRow[isRedFlag]) && (row+posOffsetsAdvisor[i][0] <= advisorLowerRow[isRedFlag]) && (column+posOffsetsAdvisor[i][1] >= PALACE_LEFT) && (column+posOffsetsAdvisor[i][1] <= PALACE_RIGHT) &&
			validPos(board, row+posOffsetsAdvisor[i][0], column+posOffsetsAdvisor[i][1], isRedFlag) {
			pos = append(pos, Pos{row + posOffsetsAdvisor[i][0], column + posOffsetsAdvisor[i][1]})
		}
	}
    
    fmt.Printf("Pos is: %q", pos)
	return pos
}

func getPossibleMovesInfo(board [13][12]int, row int, column int) MoveInfo {
	var movesInfo MoveInfo

    //fmt.Printf("Args are: %q, %d, %d", board, row, column)
    
	 movesInfo.weight = movesWeight[board[row][column]];
	
	           movesInfo.piece = board[row][column];
	           if movesInfo.piece > COLOR {
	               movesInfo.pieceColor = -1
	           }else{
	               movesInfo.pieceColor = 1
	           }
	   pieceFun := pieceFuns[movesInfo.piece]
    //fmt.Printf("piece fun is: %q", pieceFun)
        fmt.Printf("piece color is: %d==============",movesInfo.pieceColor)
	    movesInfo.moves = pieceFun(board, row, column, movesInfo.pieceColor)
	

	return movesInfo
}
