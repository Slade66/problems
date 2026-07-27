package main

func maxDepth1(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}
