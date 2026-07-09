package main

import "testing"

func TestMaxOperations(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "official example 1",
			nums: []int{1, 2, 3, 4},
			k:    5,
			want: 2,
		},
		{
			name: "official example 2",
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
			want: 1,
		},
		{
			name: "single number cannot form pair",
			nums: []int{1},
			k:    2,
			want: 0,
		},
		{
			name: "same number pairs by count",
			nums: []int{3, 3, 3, 3},
			k:    6,
			want: 2,
		},
		{
			name: "uneven complement counts",
			nums: []int{1, 1, 2, 2, 3, 3, 4},
			k:    5,
			want: 3,
		},
		{
			name: "large values near constraint",
			nums: []int{1000000000, 1, 999999999, 2},
			k:    1000000001,
			want: 2,
		},
		{
			name: "no available pair",
			nums: []int{1, 1, 1},
			k:    5,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxOperations2(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("maxOperations(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
