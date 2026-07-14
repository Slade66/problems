package main

// 推理：
// 在数组中找到最左侧的平衡点，它左侧（不包括它自己）所有元素和它右侧（不包括它自己）所有元素的和相等。
// 先扫 2 遍数组，第一遍构建左边和数组，第二遍构建右边和数组。
func pivotIndex1(nums []int) int {

	prefix := buildPrefix(nums)
	suffix := buildSuffix(nums)

	for i := 0; i < len(nums); i++ {
		if prefix[i] == suffix[i] {
			return i
		}
	}

	return -1
}

func buildPrefix(nums []int) []int {
	// a[i] = a[0] ~ a[i - 1] 的和，不包括 a[i]
	// [1, 7, 3, 6, 5, 6]
	//  0  1  8  11 17 22
	prefix := make([]int, len(nums))
	prefix[0] = 0
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	return prefix
}

func buildSuffix(nums []int) []int {
	// a[i] = a[0] ~ a[i - 1] 的和，不包括 a[i]
	// [1, 7, 3, 6, 5, 6]
	//  27 20 17 11 6  0
	suffix := make([]int, len(nums))
	lastIndex := len(nums) - 1
	suffix[lastIndex] = 0
	for i := lastIndex - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + nums[i+1]
	}
	return suffix
}
