package main

// 推理：
// 题意理解：在二进制字符串中寻找最长的连续 1 子字符串，可以通过翻转 k 个 0 达到。
// 可以转变为：滑动窗口内最多 k 个 0。
// 一旦遇到超过 k 个 0，就收缩窗口，减到 k 个为止。
//
// 复杂度：
// - 时间：O(n)
// - 空间：O(1)
func longestOnes1(nums []int, k int) int {
	left, zeroCount, maxLen := 0, 0, 0
	for right, num := range nums {
		if num == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
