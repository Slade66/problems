package main

import "testing"

func TestLongestZigZag(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "official example 1",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val:   1,
								Right: &TreeNode{Val: 1},
							},
						},
						Right: &TreeNode{Val: 1},
					},
				},
			},
			want: 3,
		},
		{
			name: "official example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   1,
							Right: &TreeNode{Val: 1},
						},
						Right: &TreeNode{Val: 1},
					},
				},
				Right: &TreeNode{Val: 1},
			},
			want: 4,
		},
		{
			name: "official example 3",
			root: &TreeNode{Val: 1},
			want: 0,
		},
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "only left children",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: 1,
		},
		{
			name: "zigzag starts below root",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val:   5,
								Right: &TreeNode{Val: 6},
							},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "longest path can start by moving right",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
			},
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestZigZag1(tc.root)
			if got != tc.want {
				t.Errorf("longestZigZag() = %d, want %d", got, tc.want)
			}
		})
	}
}
