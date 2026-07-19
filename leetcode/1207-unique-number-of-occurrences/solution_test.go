package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "official example 1",
			arr:  []int{1, 2, 2, 1, 1, 3},
			want: true,
		},
		{
			name: "official example 2",
			arr:  []int{1, 2},
			want: false,
		},
		{
			name: "official example 3",
			arr:  []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			want: true,
		},
		{
			name: "single element",
			arr:  []int{5},
			want: true,
		},
		{
			name: "all same values",
			arr:  []int{2, 2, 2, 2},
			want: true,
		},
		{
			name: "two values with same occurrence",
			arr:  []int{1, 1, 2, 2},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := uniqueOccurrences1(tc.arr)
			if got != tc.want {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tc.arr, got, tc.want)
			}
		})
	}
}
