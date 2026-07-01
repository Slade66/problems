package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "official example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "official example 2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "no zeroes",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all zeroes",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "zeroes already at end",
			nums: []int{1, 2, 0, 0},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "non-zero before middle zero",
			nums: []int{1, 0, 1},
			want: []int{1, 1, 0},
		},
		{
			name: "preserves negative non-zero order",
			nums: []int{0, -1, 0, -2, 3},
			want: []int{-1, -2, 3, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			moveZeroes(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Errorf("moveZeroes(%v) changed nums to %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}
