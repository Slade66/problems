package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "official example 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "official example 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "minimum length",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
		{
			name: "single zero",
			nums: []int{0, 4, 5},
			want: []int{20, 0, 0},
		},
		{
			name: "multiple zeros",
			nums: []int{0, 4, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3, -4},
			want: []int{-24, -12, -8, -6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := productExceptSelf(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
