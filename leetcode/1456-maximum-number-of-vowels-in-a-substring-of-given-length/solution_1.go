package main

// 推理：
// 在 s 中找出所有连续长度为 k 的子字符串，统计每个子字符串中的元音个数，返回最多的个数。
// 肯定是滑动窗口。
// 先扩展窗口到 k，记录窗口内的元音个数，然后往后前进一个，前面的窗口收缩一个，
// 每移动一次计算新增的是不是元音，减去的是不是元音，重算个数。
// 一开始 left 在 0，只有当 right - left > k 的时候，才移动 left
//
// 易错点：
// 1. left 右移之后不用再检查 s[left]，因为 right 右移的时候已经把它算入窗口中了。
//
// 复杂度：
// - 时间：O(n)
// - 空间：O(1)
func maxVowels1(s string, k int) int {
	vowelSet := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	windowVowelCount := 0
	maxVowelCount := 0
	left, right := 0, 0
	for right < len(s) {
		incomingChar := s[right]
		if _, ok := vowelSet[incomingChar]; ok {
			windowVowelCount++
			maxVowelCount = max(maxVowelCount, windowVowelCount)
		}
		right++
		if right == len(s) {
			break
		}
		if right-left >= k {
			outgoingChar := s[left]
			if _, ok := vowelSet[outgoingChar]; ok {
				windowVowelCount--
			}
			left++
		}
	}
	return maxVowelCount
}
