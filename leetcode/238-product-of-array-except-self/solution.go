package main

// 推理：
// - 要求得出索引 x 左右两侧所有的数相乘的结果（不包括 x）放到 x 上。
//
// 算法：
// - 遍历 2 次，一次从左往右，另一次从右往左，构建一个左右两侧所有数的乘积数组。
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	for i, _ := range nums {
		if i == 0 {
			left[i] = 1
			continue
		}
		left[i] = left[i-1] * nums[i-1] // 要得出 x 位置的数，首先要得到 x - 1 之前所有数的乘积，然后乘 x
	}

	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		right[i] *= left[i]
	}

	return right
}
