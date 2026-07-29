package main

import "testing"

func TestGoodNodes(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "official example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 4,
		},
		{
			name: "official example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 2},
				},
			},
			want: 3,
		},
		{
			name: "official example 3",
			root: &TreeNode{
				Val: 1,
			},
			want: 1,
		},
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "equal values on path are all good",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 2},
				},
			},
			want: 3,
		},
		{
			name: "negative values",
			root: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val:   -4,
					Right: &TreeNode{Val: -2},
				},
				Right: &TreeNode{Val: -1},
			},
			want: 3,
		},
		{
			name: "path maximum is independent between branches",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			want: 4,
		},
		{
			name: "smaller descendants are not good",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 6},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := goodNodes1(tc.root)
			if got != tc.want {
				t.Errorf("goodNodes() = %d, want %d", got, tc.want)
			}
		})
	}
}
