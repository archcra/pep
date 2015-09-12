package minimax

import "github.com/archcra/pep/boardHelper"

type TreeNode struct {
	Parent     *TreeNode //走到这个局面的上一局面
	Move       string    //走到这个局面的招法
	Board      [13][12]int
	Score      int //评分
	RoundColor int //当前局面该走的颜色
}

const (
	SCORE_WIN_LIMIT = 100000
)

func expandNode(node TreeNode, nextRoundColor int) []TreeNode {
	/*
	   A node is as this:
	   {move:"10:9-9:9", board:[...], nextMove:RED}
	*/
	var treeNodes []TreeNode
	var newNode TreeNode

	newNodes := boardHelper.Generate(node.Board, nextRoundColor)
	for _, item := range newNodes {
		newNode = TreeNode{
			&node, item.Move, item.Board, 0, nextRoundColor,
		}
		treeNodes = append(treeNodes, newNode)

	}
	return treeNodes
}

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
		"",
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
		score := boardHelper.Evaluate(node.Board) * roundColor
		if score > maxScore {
			maxNode = node
			maxScore = score
		}
	}

	// 如果没有可展开结点，即被完全憋死了（如3aga3/3ShS3/4S4/9/9/9/9/9/9/4G4，如轮到黑，则完全没有可走之步)
	//另一局，7子： (9/3ShS3/3aga3/9/9/9/9/9/9/3G5) // TODO
	return &maxNode
}

func min(nodes []TreeNode, depth int, depthLimit int, roundColor int) *TreeNode {
	//fmt.Printf("=====\nIn min with depth:%d; depthLimit:%d; roundColor: %d ======\n", depth, depthLimit, roundColor)

	var maxNodes []TreeNode
	var minNode TreeNode

	for _, node := range nodes {
		// 如果此展开的节点已吃了对方老帅，则直接返回
		if boardHelper.GeneralBeTaken(node.Board, node.RoundColor) {
			return &node
		}

		if depth == depthLimit {
			node.Score = boardHelper.Evaluate(node.Board) * roundColor

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

	return &minNode
}
