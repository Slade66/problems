package main

import "slices"

// 推理：
// 先从小到大排序，然后用双指针，一个在前，另一个在后，往中间扫：
// - 最小值 + 最大值 = k，组成一对
// - 和大于 k，说明右边大了，移动右边
// - 和小于 k，说明左边小了，移动左边
// - 直到相遇
//
// 复杂度：
// - 时间：O(n log n)，因为要排序
// - 空间：改原数组就无，否则 O(n)
func maxOperations2(nums []int, k int) int {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			left++
			right--
			count++
		} else if sum > k {
			right--
		} else {
			left++
		}
	}
	return count
}
