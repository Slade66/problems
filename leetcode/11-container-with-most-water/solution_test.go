package main

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "official example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "official example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "all zero heights",
			height: []int{0, 0, 0},
			want:   0,
		},
		{
			name:   "flat heights use widest distance",
			height: []int{5, 5, 5, 5},
			want:   15,
		},
		{
			name:   "width can beat a tall middle line",
			height: []int{1, 100, 1},
			want:   2,
		},
		{
			name:   "best pair is not at the edges",
			height: []int{2, 3, 4, 5, 18, 17, 6},
			want:   17,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea1(tc.height)
			if got != tc.want {
				t.Errorf("maxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}
