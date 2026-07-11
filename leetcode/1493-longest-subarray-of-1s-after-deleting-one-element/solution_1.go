package main

// 推理：
// 题意理解：必须从二进制数组中删除一个元素，找到最长的连续的 1 子数组。
// 窗口一直往右滑，遇到第一个 0 不管，记录索引和计数 +1，直到遇到第二个 0：
// 算一下 max = right - left - 1。
// left 移动到第一个 0 + 1 的位置，firstZero = 第二个 0，right 移动到遇到的那个 0 后面。
// 当 right == len(nums)，还要统计一次数，以防没有遇到第二个 0，对最长不计数。
// 这个算法是：最多包含一个 0 的滑动窗口
func longestSubarray1(nums []int) int {
	left, right := 0, 0
	firstZero, zeroCount := -1, 0
	longest1 := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
			if zeroCount == 1 {
				firstZero = right
			}
			if zeroCount == 2 {
				longest1 = max(longest1, right-left-1)
				left = firstZero + 1
				firstZero = right
				zeroCount = 1
			}
		}
		right++
	}
	return max(longest1, right-left-1)
}
