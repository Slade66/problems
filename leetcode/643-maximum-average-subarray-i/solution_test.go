package main

import (
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{
			name: "official example 1",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
		{
			name: "official example 2",
			nums: []int{5},
			k:    1,
			want: 5,
		},
		{
			name: "window length equals array length",
			nums: []int{3, -1, 4, -2},
			k:    4,
			want: 1,
		},
		{
			name: "best window at beginning",
			nums: []int{9, 8, 7, 1, 2, 3},
			k:    3,
			want: 8,
		},
		{
			name: "best window at end",
			nums: []int{-10, -5, 4, 8, 12},
			k:    3,
			want: 8,
		},
		{
			name: "all negative numbers",
			nums: []int{-7, -3, -5, -2},
			k:    2,
			want: -3.5,
		},
		{
			name: "single element windows choose max element",
			nums: []int{-1, 0, 10, -3},
			k:    1,
			want: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaxAverage(tc.nums, tc.k)
			if math.Abs(got-tc.want) > 1e-5 {
				t.Errorf("findMaxAverage(%v, %d) = %f, want %f", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
