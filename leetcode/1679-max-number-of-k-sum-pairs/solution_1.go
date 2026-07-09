package main

// 推理：
// 题意理解：可以从 nums 中找到多少对相加等于 k 的数。
// 可以做个 map，一次遍历，访问 nums[x]，找 k - nums[x] 在 map 中是否存在。
// 如果存在，y = k - nums[x]，map[y]--，count++；否则 map[nums[x]]++。
//
// 易错点：
//  1. map[int]int 里原来有 3:1，后来把 m[3] = 0，此时 v, ok := m[3] 的 ok 仍然是 true，
//     因为 Go map 的 comma ok 判断的是 key 是否存在，不是 value 是否为零值；只有执行 delete(m, 3) 后，ok 才会变成 false。
//
// 复杂度分析：
// - 时间 O(n)
// - 空间 O(n)
func maxOperations1(nums []int, k int) int {
	m := make(map[int]int)
	count := 0
	for _, v := range nums {
		pairNum := k - v
		if _, ok := m[pairNum]; ok {
			m[pairNum]--
			if m[pairNum] == 0 {
				delete(m, pairNum)
			}
			count++
		} else {
			m[v]++
		}
	}
	return count
}
