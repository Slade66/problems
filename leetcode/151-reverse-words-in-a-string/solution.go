package main

import "strings"

// 推理：
// - 单词内部没有空格。
// - 每个单词之间都用一个空格分隔。
// - 返回的字符串只在单词之间留一个空格，多余的空格删掉。
//
// 算法：
// - 从右往左扫 s，先跳过尾部的空格，然后识别第一个单词的边界，然后 append 到 result。
// - p1 识别单词尾部，碰到非空格之后在它的 -1 位置创建一个 p2，往前遍历找到单词的头部。
// - p2 从 p1 - 1 开始，遇到第一个空格视为找到单词的头部。
// - 如果 p1 < 0 了，说明扫完了，退出。
func reverseWords(s string) string {
	result := make([]string, 0, len(s))
	p1 := len(s) - 1
	for p1 >= 0 {
		for p1 >= 0 && s[p1] == ' ' {
			p1--
		}
		if p1 < 0 {
			break
		}
		p2 := p1 - 1
		for p2 >= 0 && s[p2] != ' ' {
			p2--
		}
		result = append(result, s[p2+1:p1+1])
		p1 = p2
	}
	return strings.Join(result, " ")
}
