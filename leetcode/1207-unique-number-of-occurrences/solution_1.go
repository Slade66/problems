package main

// 推理：
// 找数组各个数字的出现次数是否是唯一的。
// 肯定是先统计各个数字的出现次数。然后再遍历次数，找是否唯一。
func uniqueOccurrences1(arr []int) bool {
	// 1. 记录数字的出现次数。
	freq := make(map[int]int, len(arr))
	for _, num := range arr {
		freq[num]++
	}

	// 2. 检查数字的出现次数是否唯一。
	seen := make(map[int]bool, len(freq))
	for _, count := range freq {
		if _, ok := seen[count]; ok {
			return false
		}
		seen[count] = true
	}

	return true
}
