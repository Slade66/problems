package main

// 推理：
// 在包含 * 的字符串 s 中，你可以一次移除一个 * 和它左边最近的非 * 字符。
// 遍历一遍，从左往右加入切片，如果遇到 *，就不加 *，并把之前添加的那个元素删了。
func removeStars1(s string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		currChar := s[i]
		if currChar == '*' {
			result = result[:len(result)-1]
			continue
		}
		result = append(result, currChar)
	}
	return string(result)
}
