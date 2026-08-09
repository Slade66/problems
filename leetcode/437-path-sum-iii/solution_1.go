package main

func pathSum1(root *TreeNode, targetSum int) int {
	nodeQueue := []*TreeNode{root}
	pathCount := 0

	for len(nodeQueue) > 0 {
		currNode := nodeQueue[0]
		if currNode != nil {
			nodeQueue = append(nodeQueue, currNode.Left, currNode.Right)
		}
		nodeQueue = nodeQueue[1:]

		pathCount += calcPath(currNode, targetSum, 0)
	}

	return pathCount
}

func calcPath(currNode *TreeNode, targetSum, currSum int) int {
	if currNode == nil {
		return 0
	}

	currSum += currNode.Val
	isEqualTargetSum := 0
	if currSum == targetSum {
		isEqualTargetSum = 1
	}
	return calcPath(currNode.Left, targetSum, currSum) + isEqualTargetSum + calcPath(currNode.Right, targetSum, currSum)
}
