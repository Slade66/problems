package main

func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST1(root.Right, val)
	}
	return searchBST1(root.Left, val)
}
