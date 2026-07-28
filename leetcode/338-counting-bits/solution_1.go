package main

func countBits1(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		count := 0
		for j := i; j > 0; j >>= 1{
			if j & 1 == 1 {
				count++
			}
		}
		result = append(result, count)
	}
	return result
}
