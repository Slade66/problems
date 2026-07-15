package main

import "testing"

func TestEqualPairs(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "official example 1",
			grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
			want: 1,
		},
		{
			name: "official example 2",
			grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
			want: 3,
		},
		{
			name: "single element grid",
			grid: [][]int{{1}},
			want: 1,
		},
		{
			name: "no equal pairs",
			grid: [][]int{{1, 2}, {3, 4}},
			want: 0,
		},
		{
			name: "all rows and columns equal",
			grid: [][]int{{1, 1}, {1, 1}},
			want: 4,
		},
		{
			name: "two distinct equal pairs",
			grid: [][]int{{1, 2}, {2, 1}},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := equalPairs1(tc.grid)
			if got != tc.want {
				t.Errorf("equalPairs(%v) = %d, want %d", tc.grid, got, tc.want)
			}
		})
	}
}
