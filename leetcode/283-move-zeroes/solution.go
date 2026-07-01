package main

// 算法：
// - zeroPtr 先找到最左边的 0。
// - nonZeroPtr 再从 zeroPtr 右边找到最近的非 0。
// - 两者交换后继续向右扫，直到右边已经没有非 0。
func moveZeroes(nums []int) {
	zeroPtr, nonZeroPtr := 0, 1
	for {
		for zeroPtr < len(nums) && nums[zeroPtr] != 0 {
			zeroPtr++
		}
		if nonZeroPtr <= zeroPtr {
			nonZeroPtr = zeroPtr + 1
		}
		for nonZeroPtr < len(nums) && nums[nonZeroPtr] == 0 {
			nonZeroPtr++
		}
		if zeroPtr >= len(nums) || nonZeroPtr >= len(nums) {
			break
		}
		nums[zeroPtr], nums[nonZeroPtr] = nums[nonZeroPtr], nums[zeroPtr]
		zeroPtr++
		nonZeroPtr++
	}
}
