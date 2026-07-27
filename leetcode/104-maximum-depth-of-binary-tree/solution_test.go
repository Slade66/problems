package main

import "testing"

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "official example 1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "official example 2",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: 2,
		},
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: 1,
		},
		{
			name: "left-skewed tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
			want: 4,
		},
		{
			name: "deepest leaf is not on the outer edge",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}},
				},
				Right: &TreeNode{Val: 3},
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxDepth1(tc.root)
			if got != tc.want {
				t.Errorf("maxDepth() = %d, want %d", got, tc.want)
			}
		})
	}
}
