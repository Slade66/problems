package main

// 推理：
// 1 to n，天然有序
// 两个指针，一个指向头，一个指向尾。
// 每轮算出一个中点，看被猜的数落在哪个区间。
// 如果比中点大，就把左指针移到中点右侧，反之亦然。
// 题目保证 1 <= pick <= n，答案一定存在。
func guessNumber1(n int) int {
	left, right := 1, n
	for {
		mid := (right-left)/2 + left
		guessResult := guess(mid)
		switch guessResult {
		case 0:
			return mid
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		}
	}
}
