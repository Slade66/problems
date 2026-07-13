package main

// 推理：
// 把数组里的数依次累加，每次比较是不是最大，找最大的那个。
func largestAltitude1(gain []int) int {
	maxAltitude := 0
	currentAltitude := 0
	for _, v := range gain {
		currentAltitude += v
		maxAltitude = max(maxAltitude, currentAltitude)
	}
	return maxAltitude
}
