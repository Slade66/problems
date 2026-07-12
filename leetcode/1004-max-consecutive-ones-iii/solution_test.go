package main

import "testing"

func TestLongestOnes(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "official example 1",
			nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
			want: 6,
		},
		{
			name: "official example 2",
			nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
			want: 10,
		},
		{
			name: "no flips",
			nums: []int{1, 1, 0, 1, 1},
			k:    0,
			want: 2,
		},
		{
			name: "flip all zeros",
			nums: []int{0, 1, 0, 1, 0},
			k:    3,
			want: 5,
		},
		{
			name: "single zero without available flip",
			nums: []int{0},
			k:    0,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestOnes1(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("longestOnes(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
