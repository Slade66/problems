package main

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "official example 1",
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			name: "official example 2",
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			name: "lower boundary zero",
			n:    0,
			want: []int{0},
		},
		{
			name: "single positive number",
			n:    1,
			want: []int{0, 1},
		},
		{
			name: "crosses a power of two",
			n:    8,
			want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := countBits1(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("countBits(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
