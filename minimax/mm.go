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
	   {move:"10:9-9:9", board:[...], RoundColor:RED}
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

func rivalUnderCheck(node TreeNode, roundColor int) bool {

	nodes := expandNode(node, roundColor)
	for _, nodeItem := range nodes {
		// 如果此展开的节点自己老帅被吃了，则直接返回最小值
		if boardHelper.GeneralBeTaken(nodeItem.Board, roundColor) {
			return true
		}
	}
	return false
}
