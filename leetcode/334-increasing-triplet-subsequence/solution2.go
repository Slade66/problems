package main

func increasingTriplet2(nums []int) bool {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i] >= 3 {
				return true
			}
		}
	}
	return false
}
