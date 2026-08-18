package main

func longestZigZag1(rootNode *TreeNode) int {
	if rootNode == nil {
		return 0
	}

	pendingStartNodes := []*TreeNode{rootNode}
	longestZigZagLength := 0

	// 遍历树中的每一个节点。
	for len(pendingStartNodes) > 0 {
		currentStartNode := pendingStartNodes[0]
		pendingStartNodes = pendingStartNodes[1:] // pop 头节点

		// 将左右子节点加入待遍历队列。
		if currentStartNode.Left != nil {
			pendingStartNodes = append(pendingStartNodes, currentStartNode.Left)
		}
		if currentStartNode.Right != nil {
			pendingStartNodes = append(pendingStartNodes, currentStartNode.Right)
		}

		// 以当前节点为起点，分别尝试向左和向右移动。
		longestPathFromCurrentStartNode := max(
			zigZagLengthFromNode(currentStartNode, MoveLeft),
			zigZagLengthFromNode(currentStartNode, MoveRight),
		)
		longestZigZagLength = max(longestZigZagLength, longestPathFromCurrentStartNode)
	}

	return longestZigZagLength
}

type MoveDirection int

const (
	MoveLeft MoveDirection = iota
	MoveRight
)

func zigZagLengthFromNode(currentNode *TreeNode, requiredNextDirection MoveDirection) int {
	if currentNode.Left != nil && requiredNextDirection == MoveLeft {
		return 1 + zigZagLengthFromNode(currentNode.Left, MoveRight)
	}
	if currentNode.Right != nil && requiredNextDirection == MoveRight {
		return 1 + zigZagLengthFromNode(currentNode.Right, MoveLeft)
	}
	return 0
}
