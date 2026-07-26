package main

// pickedNumber simulates the hidden number selected by LeetCode.
var pickedNumber int

// guess simulates the API predefined by LeetCode.
func guess(num int) int {
	switch {
	case num > pickedNumber:
		return -1
	case num < pickedNumber:
		return 1
	default:
		return 0
	}
}

// 推理：
func guessNumber(n int) int {
	return 0
}
