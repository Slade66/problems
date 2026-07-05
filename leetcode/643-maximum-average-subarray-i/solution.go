package main

// 推理：
// - 子数组是连续的序列。
//
// 算法：
// - 先算出前缀和，然后从前缀和数组的第 k 位开始往后，移动一位就用当前所在位置的前缀和减去 k 位前的前缀和，再算平均数。
func findMaxAverage(nums []int, k int) float64 {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	maxSum := prefix[k] - prefix[0]
	for i := k + 1; i <= len(nums); i++ {
		sum := prefix[i] - prefix[i-k]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}
