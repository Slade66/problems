package main

// 推理：
// - 要按顺序：只有 s 的第一个字符出现了，才能判断 s 的第二个字符是否出现。
//
// 算法：
// - 两个指针，一个指向 s，另一个指向 t；从左往右扫 t，遇到 s 的一个字符，s 的指针就往右移一位，看 t 扫完之后，s 扫完没有。
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	sPtr := 0
	for _, v := range t {
		if s[sPtr] == byte(v) {
			sPtr++
			if sPtr == len(s) {
				return true
			}
		}
	}
	return false
}
