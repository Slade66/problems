package main

import (
	"strconv"
	"strings"
)

// 推理：
// 把字符串按照规则展开。
// k[...] 里的字符串会被左边外面紧挨着的数字重复 k 次。
// 没有多余的空格和数字，只有 [] 外面会有数字。
// 准备个结果字符串，一路遍历：
// 1. 遇到正常出现的字符，就直接加入结果。
// 遇到数字[]和嵌套的情况怎么办？
// 怎么样说通逻辑？
// 我需要一个函数，遇到数字触发它。
// 我需要界定范围，界定被 repeat 的范围。
// 肯定是有个函数，它的参数是 k, str 把 str 重复 k 遍。
// 先展开最外层的重复：即使内容中仍有嵌套括号，也先保留；不断重复这个展开过程，直到结果中不再包含括号。
// 准备一个栈，遇到数字之后，left 指针停，往栈里面塞 [ [ ] ]，另一个指针 right 往后遍历，直到栈里的括号都被抵消，此时范围界定完毕，
// 把 [left + 1, right] 加入 repeatString 函数中，直到扫到末尾。只要 result 里还有 [] 的存在，就重复展开。
// 怎么界定 result 里还有 []？一次展开只能解一层 []，如果在每轮展开在栈中发现多个 []，就记录有更多 []。
// 注意数字不只是一位数，在 1 ～ 300 之间。
// 那这样的做法就复杂了。
//
// 能否一次 O(n) 把所有数字和括号找出来？预处理一轮。如果只是找 [ 的话。
// [ 往前遇到的第一个非数字，就是数字的范围，往后第一个
//
// 还是往后遍历，遇到数字就进入特殊处理模式（往后找 []），然后把 k 和 str 丢进函数，然后后移 i。
func decodeString1(s string) string {
	for strings.Contains(s, "[") {
		var result strings.Builder
		for i := 0; i < len(s); i++ {
			currChar := s[i]
			if isNumber(currChar) {
				newI, k, str := findRepeatRange(s, i)
				i = newI
				result.WriteString(repeatString(k, str))
			} else {
				// 正常的字符直接添加到 result 里。
				result.WriteByte(currChar)
			}
		}
		s = result.String()
	}
	return s
}

func findRepeatRange(s string, i int) (newI int, k int, str string) {
	// 1. 找数字
	start := i
	for isNumber(s[i]) {
		i++
	}
	times, _ := strconv.Atoi(s[start:i])

	// 2. 找 [] 的范围
	brackets := 1
	start = i
	i++
	for ; brackets > 0; i++ {
		switch s[i] {
		case '[':
			brackets++
		case ']':
			brackets--
		}
	}
	i--
	return i, times, s[start+1 : i]
}

func isNumber(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func repeatString(k int, s string) string {
	var builder strings.Builder
	for i := 0; i < k; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}
