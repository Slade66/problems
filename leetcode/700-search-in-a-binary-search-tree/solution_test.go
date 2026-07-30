package main

import "testing"

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestSearchBST(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}
	node7 := &TreeNode{Val: 7}
	tree1 := &TreeNode{
		Val:   4,
		Left:  node2,
		Right: node7,
	}

	testCases := []struct {
		name string
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			name: "official example 1",
			root: tree1,
			val:  2,
			want: node2,
		},
		{
			name: "official example 2",
			root: tree1,
			val:  5,
			want: nil,
		},
		{
			name: "empty tree",
			root: nil,
			val:  5,
			want: nil,
		},
		{
			name: "target is root node",
			root: tree1,
			val:  4,
			want: tree1,
		},
		{
			name: "target is leaf node",
			root: tree1,
			val:  3,
			want: node3,
		},
		{
			name: "single node match",
			root: &TreeNode{Val: 10},
			val:  10,
			want: &TreeNode{Val: 10},
		},
		{
			name: "single node mismatch",
			root: &TreeNode{Val: 10},
			val:  5,
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := searchBST1(tc.root, tc.val)
			if !isSameTree(got, tc.want) {
				t.Errorf("searchBST() = %v, want %v", got, tc.want)
			}
		})
	}
}
