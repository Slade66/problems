package main

import "strconv"

// 推理：
// chars = ["a","a","b","b","c","c","c"]
// 2 个指针 + 1 个计数变量。
// 指针 p1 指向当前要数的字符。
// 另一个指针 p2 往后扫，看能扫到多少个一样的字符。
// 扫到一个就 +1，扫到不同的字符，就把 p1 + 1 位置的字符改为数量：
// 只有一个就不改了。
// 单数就直接改为数字。
// 超过 1 位数，就拆成多个字符。
// 把指针 p1 移到 p2 的位置：p1 = p2，p2 + 1。
// 问题：多余部分的边界无法确定。
// 还得有一个指针记录边界位置。它应该初始化在当前要数的字符的位置，然后每次把 p1 的字符 + 数量填过来，数量几位数移几位。
//
// 易错点：
// 1. byte(count % 10) 不是字符数字。byte(2) 是 ASCII 控制字符，不是字符 '2'。
// 2. count == 1 的时候不写数字，只写字符。count > 1：写字符 + 数量
// 3. 只在遇到不同字符时写上一组，但最后一组后面没有“不同字符”触发 else，所以必须在循环结束后补一次。
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	p1, p2, p3, count := 0, 0, 0, 0
	for p2 < len(chars) {
		if chars[p1] == chars[p2] {
			count++
			p2++
		} else {
			p3 = writeCountToChars(chars, p1, p3, count)
			p1 = p2
			p2++
			count = 1
		}
	}
	p3 = writeCountToChars(chars, p1, p3, count) // 防止最后循环完没遇到不同的字符

	return p3
}

// 返回新的 p3 位置
func writeCountToChars(chars []byte, p1, p3, count int) int {
	chars[p3] = chars[p1]
	p3++

	if count == 1 {
		return p3
	}

	countStr := strconv.Itoa(count)
	for i := 0; i < len(countStr); i, p3 = i+1, p3+1 {
		chars[p3] = countStr[i]
	}

	return p3
}
