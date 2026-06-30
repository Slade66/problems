package main

// 推理：
// - 在循环中，从上到下，先取 word1 的字符，再取 word2 的字符，天然的交替顺序。
// - 循环次数取最长字符串的长度，较短的字符串取完了用 if 保护。
func mergeAlternately(word1 string, word2 string) string {
	word1Len, word2Len := len(word1), len(word2)
	maxLen := max(word1Len, word2Len)
	result := make([]byte, 0, maxLen)
	for i := 0; i < maxLen; i++ {
		if i < word1Len {
			result = append(result, word1[i])
		}
		if i < word2Len {
			result = append(result, word2[i])
		}
	}
	return string(result)
}
