package main

func increasingTriplet3(nums []int) bool {
	miniLeft := make([]int, len(nums)) // 0...x 左侧（包括自己）最小的
	maxRight := make([]int, len(nums)) // x...len(nums)-1 右侧（包括自己）最大的

	// 预处理：构建 nums[i] 左侧（包括自己）的最小值
	miniLeft[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		miniLeft[i] = min(miniLeft[i-1], nums[i])
	}

	// 预处理：构建 nums[i] 右侧（包括自己）的最大值
	maxRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], nums[i])
	}

	// 如果存在一个数，它的左侧的最小值小于它，且右侧的最大值大于它，则存在递增三元组
	for i, v := range nums {
		if miniLeft[i] < v && v < maxRight[i] {
			return true
		}
	}

	return false
}
