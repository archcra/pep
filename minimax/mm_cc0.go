// cc - continually check - 连将
// 这个算法有如个特点：
// 1. 层数为奇数，到最低层如还不能吃掉对方老将（或完全困毙），则不要此步(分值最低)
// 2. 不要评分函数-以提高效率-只检查是否吃了老将
// 3. 奇数搜索层，每次都判断是否对方将处于被攻击状态-这个的判断
// 使用此局面再一次此颜色展开的方法来进行：看是否已吃了对方老将，如有则返回true

package minimax

import "github.com/archcra/pep/boardHelper"

func MinimaxCc(boardFen string, depthLimit int, roundColor int) *TreeNode {
	// 由于minmax考虑了一次大（自己走），一次小（对方走），所以评分时，永远考虑为一种颜色即可
	// 所以roundColor永远是一样的，不需要改变；
	// 当minimax开始时，颜色为黑值时，即站在黑的角色，则传入roundColor为-1；
	// 评分后乘以这个roundColor就是正确的分数了；然后仍是最大最小地处理

	// DeepLimit must be even, must take a whole round steps as consideration,
	// not just move of one side.

	// roundColor RED: 1, BLACK: -1

	// depthLimit must be odd number

	treeNode := TreeNode{
		nil,
		"ROOT",
		boardHelper.Fen2Board(boardFen),
		0,
		roundColor,
	}
	return maxCc(expandNode(treeNode, roundColor), 1, depthLimit, roundColor)
}

func maxCc(nodes []TreeNode, depth int, depthLimit int, roundColor int) *TreeNode {

	// max 展开的节点，表明当前的评估分数，和要得出步骤的棋子颜色相同
	var minNodes []TreeNode
	var maxNode TreeNode

	//fmt.Printf("total nodes are: %d with color:%d", len(nodes), nodes[0].RoundColor)
	for _, node := range nodes {
		// 如果此展开的节点已吃了对方老帅，则直接返回
		if boardHelper.GeneralBeTaken(node.Board, node.RoundColor) {
			node.Score = SCORE_WIN_LIMIT
			return &node
		} else {
			// 如果展开的节点不带将，将忽略此节点，不再展开
			if rivalUnderCheck(node, node.RoundColor) && depth != depthLimit {
				//fmt.Print("\n---->rival under check with move: %s \n", node.Move)
				minNodes = append(minNodes, *minCc(expandNode(node, roundColor*-1), depth+1, depthLimit, roundColor))
			}
		}
	}

	// find the max node in the minNodes
	maxScore := -SCORE_WIN_LIMIT

	for _, node := range minNodes {
		if node.Score >= maxScore {
			maxNode = node
			maxScore = node.Score
		}
	}

	// 如果没有可展开结点，即被完全憋死了（如3aga3/3ShS3/4S4/9/9/9/9/9/9/4G4，如轮到黑，则完全没有可走之步)
	//另一局，7子： (9/3ShS3/3aga3/9/9/9/9/9/9/3G5) // TODO
	//maxNode.Parent.Score = maxNode.Score
	if maxNode.Parent == nil { //所有的红方招法都没将，找不到下级展开
		//也可能是达到了深度
		//fmt.Printf("\n --> Nil found, with maxScore = %d and move =%s "+
		//"and minNodes count: %d and nodes.count=%d and depth is %d \n",
		//maxScore, maxNode.Move, len(minNodes), len(nodes), depth)
		nodes[0].Score = -SCORE_WIN_LIMIT
		return &nodes[0] // TODO，如果nodes的长度为0，如何处理？
	} else {
		return maxNode.Parent
	}
}

func minCc(nodes []TreeNode, depth int, depthLimit int, roundColor int) *TreeNode {

	var maxNodes []TreeNode
	var minNode TreeNode

	for _, node := range nodes {
		// 如果此展开的节点已吃了对方老帅，则直接返回
		if boardHelper.GeneralBeTaken(node.Board, node.RoundColor) {
			node.Score = -SCORE_WIN_LIMIT
			return &node
		}
		maxNodes = append(maxNodes, *maxCc(expandNode(node, roundColor), depth+1, depthLimit, roundColor))
	}

	// find the min node in the minNodes
	minScore := SCORE_WIN_LIMIT

	for _, node := range maxNodes {
		if node.Score <= minScore {
			minNode = node
			minScore = node.Score
		}
	}
	//fmt.Printf("\n\nscore of minnode is %d. with move %s\n and parent move: %s",
	//	minNode.Score, minNode.Move, minNode.Parent.Parent.Move)
	minNode.Parent.Score = minScore
	return minNode.Parent

}
