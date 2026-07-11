package main

func longestSubarray2(nums []int) int {
	left, zeroCount, longest := 0, 0, 0

	for right, num := range nums {
		if num == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// 窗口长度减 1，因为必须删除一个元素。
		longest = max(longest, right-left)
	}

	return longest
}
