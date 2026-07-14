package main

import "testing"

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "official example 1",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "official example 2",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "official example 3",
			nums: []int{2, 1, -1},
			want: 0,
		},
		{
			name: "pivot at right edge",
			nums: []int{-1, -1, 0, 2, 1},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{42},
			want: 0,
		},
		{
			name: "leftmost among multiple pivots",
			nums: []int{0, 0, 0},
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := pivotIndex1(tc.nums)
			if got != tc.want {
				t.Errorf("pivotIndex(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
