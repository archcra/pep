package boardHelper

func positionScore(board [13][12]int) int {
	//  var nextPossibleMoves =  movesGenerator.generate(node.board, node.nextMoveColor);

	// 1. Mobility

	mobility := countMobility(board)

	// 2. count of cross river solder

	//4. attack targets count/Watch friend pieces count

	//5. prey targets count(the target is not protected)

	//6. attack the general count

	//7. matelity: distance differ from the opposite g
	matelity := countMateScore(board)

	return mobility + matelity
}

func Evaluate(board [13][12]int) int {
	pieceScore := BoardPiecesScore(board)
	posScore := positionScore(board)

	return posScore + pieceScore

}
