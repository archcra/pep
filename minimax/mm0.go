package minimax

// 这个算法很慢，最多2回合(search level 4)，耗时已到分钟级别
// 而且，这个无法解决吃马的问题，即第一步如不知对方马，也会被对方吃另一马；而如果只
// 查询一个回合，则是个无解的问题；而且，这个的机动性(mobility)的权重不好，导致
// 走子比防止丢子重要
// 不过，将作为，如“有来有去”　的级别AI保留

import "github.com/archcra/pep/boardHelper"

func Minimax(board [13][12]int, depthLimit int, roundColor int) *TreeNode {

	// 由于minmax考虑了一次大（自己走），一次小（对方走），所以评分时，永远考虑为一种颜色即可
	// 所以roundColor永远是一样的，不需要改变；
	// 当minimax开始时，颜色为黑值时，即站在黑的角色，则传入roundColor为-1；
	// 评分后乘以这个roundColor就是正确的分数了；然后仍是最大最小地处理

	// DeepLimit must be even, must take a whole round steps as consideration,
	// not just move of one side.

	// roundColor RED: 1, BLACK: -1

	treeNode := TreeNode{
		nil,
		"ROOT",
		board,
		-10000,
		roundColor,
	}
	return max(expandNode(treeNode, roundColor), 1, depthLimit, roundColor)
}

func max(nodes []TreeNode, depth int, depthLimit int, roundColor int) *TreeNode {

	//console.log("in max: , nodes are: ", nodes);
	// max 展开的节点，表明当前的评估分数，和要得出步骤的棋子颜色相同
	var minNodes []TreeNode
	var maxNode TreeNode

	for _, node := range nodes {
		// 如果此展开的节点已吃了对方老帅，则直接返回
		if boardHelper.GeneralBeTaken(node.Board, node.RoundColor) {
			return &node
		}

		if depth == depthLimit {
			node.Score = boardHelper.Evaluate(node.Board) * roundColor
			minNodes = append(minNodes, node)

		} else {
			minNodes = append(minNodes, *min(expandNode(node, roundColor*-1), depth+1, depthLimit, roundColor))
		}
	}

	// find the max node in the minNodes
	maxScore := -SCORE_WIN_LIMIT

	for _, node := range minNodes {
		//fmt.Printf("\n++for each max, min score is: %d with move %s\n", node.Score, node.Parent.Move)
		if node.Score > maxScore {
			maxNode = node
			maxScore = node.Score
		}
	}
	//fmt.Printf("\n\n\n ____MaxScore is %d and MaxNode Parent Move is %s.\n", maxScore, maxNode.Parent.Move)

	// 如果没有可展开结点，即被完全憋死了（如3aga3/3ShS3/4S4/9/9/9/9/9/9/4G4，如轮到黑，则完全没有可走之步)
	//另一局，7子： (9/3ShS3/3aga3/9/9/9/9/9/9/3G5) // TODO
	//maxNode.Parent.Score = maxNode.Score
	if depth == 1 {
		return maxNode.Parent
	} else {
		//fmt.Printf("\n\n\nSHOULD RETURN HERE ....%s and %s and %d and maxScore is %d\n\n", maxNode.Move, maxNode.Parent.Move, maxNode.Score, maxScore)
		return &maxNode
	}

}

func min(nodes []TreeNode, depth int, depthLimit int, roundColor int) *TreeNode {

	var maxNodes []TreeNode
	var minNode TreeNode

	for _, node := range nodes {
		// 如果此展开的节点自己老帅被吃了，则直接返回最小值
		if boardHelper.GeneralBeTaken(node.Board, node.RoundColor) {
			node.Score = -SCORE_WIN_LIMIT
			return &node
		}

		if depth == depthLimit {
			node.Score = boardHelper.Evaluate(node.Board) * roundColor
			//fmt.Printf("======min node score is %d and board fen is %s=====", node.Score, boardHelper.Board2Fen(node.Board))
			maxNodes = append(maxNodes, node)
		} else {
			maxNodes = append(maxNodes, *max(expandNode(node, roundColor), depth+1, depthLimit, roundColor))
		}
	}

	// find the min node in the minNodes
	minScore := SCORE_WIN_LIMIT

	for _, node := range maxNodes {
		if node.Score < minScore {
			minNode = node
			minScore = node.Score
		}
	}

	//fmt.Printf("minNode is %q with score: %d.", minNode, minScore)
	//	minNode.Parent.Score = minNode.Score
	//fmt.Printf("min node move is %s and score is %d.", minNode.Move, minNode.Score)

	//if depth == depthLimit {
	return &minNode
	//} else {
	//	return minNode.Parent
	//}
}
