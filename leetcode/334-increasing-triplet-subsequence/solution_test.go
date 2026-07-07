package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "official example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "official example 2",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "official example 3",
			nums: []int{2, 1, 5, 0, 4, 6},
			want: true,
		},
		{
			name: "short length",
			nums: []int{1, 2},
			want: false,
		},
		{
			name: "strictly increasing required",
			nums: []int{1, 1, 1, 1, 1},
			want: false,
		},
		{
			name: "decreasing sequence",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "negative values with triplet",
			nums: []int{-2, -1, -3, 0, 2},
			want: true,
		},
		{
			name: "mid updates after drop",
			nums: []int{5, 1, 5, 2, 3},
			want: true,
		},
		{
			name: "long drop before final triplet",
			nums: longDropBeforeFinalTriplet(),
			want: true,
		},
		{
			name: "many equal values before final triplet",
			nums: manyEqualValuesBeforeFinalTriplet(),
			want: true,
		},
		{
			name: "duplicates can still form triplet",
			nums: []int{1, 1, 2, 2, 3},
			want: true,
		},
		{
			name: "duplicates without triplet",
			nums: []int{1, 1, 2, 2},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := increasingTriplet3(tc.nums)
			if got != tc.want {
				t.Errorf("%s: increasingTriplet(len=%d) = %v, want %v", tc.name, len(tc.nums), got, tc.want)
			}
		})
	}
}

func longDropBeforeFinalTriplet() []int {
	nums := make([]int, 3790)
	nums[0] = 1
	nums[3770] = -1
	nums[3789] = 100000000
	return nums
}

// 数组是 500000 长度：499998 个 1，最后是 3, 7，所以期望是 true。用 helper 生成，不要把 50 万个元素直接写进测试文件。
func manyEqualValuesBeforeFinalTriplet() []int {
	nums := make([]int, 500000)
	for i := 0; i < len(nums)-2; i++ {
		nums[i] = 1
	}
	nums[len(nums)-2] = 3
	nums[len(nums)-1] = 7
	return nums
}
