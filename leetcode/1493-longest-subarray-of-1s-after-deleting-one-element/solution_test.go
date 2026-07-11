package main

import "testing"

func TestLongestSubarray(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "official example 1",
			nums: []int{1, 1, 0, 1},
			want: 3,
		},
		{
			name: "official example 2",
			nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want: 5,
		},
		{
			name: "official example 3",
			nums: []int{1, 1, 1},
			want: 2,
		},
		{
			name: "single one must be deleted",
			nums: []int{1},
			want: 0,
		},
		{
			name: "single zero",
			nums: []int{0},
			want: 0,
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "one zero joins two runs",
			nums: []int{1, 1, 0, 1, 1, 1},
			want: 5,
		},
		{
			name: "zeros at both ends",
			nums: []int{0, 1, 1, 0},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := longestSubarray1(tc.nums)
			if got != tc.want {
				t.Errorf("longestSubarray(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
