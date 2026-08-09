package main

import "testing"

func TestPathSum(t *testing.T) {
	testCases := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      int
	}{
		{
			name: "official example 1",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: -2},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 1},
					},
				},
				Right: &TreeNode{
					Val:   -3,
					Right: &TreeNode{Val: 11},
				},
			},
			targetSum: 8,
			want:      3,
		},
		{
			name: "official example 2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			want:      3,
		},
		{
			name:      "empty tree",
			root:      nil,
			targetSum: 0,
			want:      0,
		},
		{
			name:      "single node equals target",
			root:      &TreeNode{Val: 5},
			targetSum: 5,
			want:      1,
		},
		{
			name:      "single node not equal to target",
			root:      &TreeNode{Val: 5},
			targetSum: 1,
			want:      0,
		},
		{
			name: "path does not start at root",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			targetSum: 5,
			want:      1,
		},
		{
			name: "negative values form a path",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  -2,
					Left: &TreeNode{Val: -3},
				},
				Right: &TreeNode{Val: 3},
			},
			targetSum: -1,
			want:      1,
		},
		{
			name: "multiple overlapping paths",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 1},
			},
			targetSum: 1,
			want:      5,
		},
		{
			name: "no matching path",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 10,
			want:      0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := pathSum1(tc.root, tc.targetSum)
			if got != tc.want {
				t.Errorf("pathSum() = %d, want %d", got, tc.want)
			}
		})
	}
}
