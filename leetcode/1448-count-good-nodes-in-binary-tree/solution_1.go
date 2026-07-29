package main

func goodNodes1(root *TreeNode) int {
	initialMaxValueOnPath := -10001 // value 的下界是 -10^4
	return countGoodNodesInSubtree(root, initialMaxValueOnPath)
}

func countGoodNodesInSubtree(node *TreeNode, maxValueOnPath int) int {
	currentNodeGoodCount := 0
	if node == nil {
		return 0
	}
	if node.Val >= maxValueOnPath {
		currentNodeGoodCount = 1
		maxValueOnPath = node.Val
	}
	return countGoodNodesInSubtree(node.Left, maxValueOnPath) + currentNodeGoodCount + countGoodNodesInSubtree(node.Right, maxValueOnPath)
}
