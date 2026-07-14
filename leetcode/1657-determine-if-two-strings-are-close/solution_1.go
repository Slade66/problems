package main

import "slices"

// 推理：
// 判断两个字符串是否相似。如果一个字符串 a 可以通过 2 种操作得到字符串 b，则它俩相似。
// 操作 1：任意两个字符换位。
// 操作 2：把字符 a 和 b 变换（a 变成 b，b 变成 a）。
// 长度不相等肯定不相似。
// 遍历俩串，记录字符的出现次数：
// 如果两个字符串的各个字符出现次数都一样，则相似（换位）。
// 如果字符的出现次数序列（1、2、3）两串都非顺序一样，则相似（变换）。
// 由于只能在出现的字符内变换，两个字符串包含的字符种类必须一样。
// 当只有 1 个字符的时候，无法交换和变换，不一样就真的不一样。
//
// 易错点：
// 隐藏条件：两个字符串必须包含相同的字符种类。例如 "aabbcc" 和 "ddeeff" 的频次都为 [2,2,2]，但字符集合不同，所以不能互相转换。
func closeStrings1(word1 string, word2 string) bool {

	word1Set := countChar(word1)
	word2Set := countChar(word2)

	isHaveSameChars := haveSameCharacters(word1Set, word2Set)
	if !isHaveSameChars {
		return false
	}

	m_v1 := getMapSortedValues(word1Set)
	m_v2 := getMapSortedValues(word2Set)

	if !slices.Equal(m_v1, m_v2) {
		return false
	}

	return true
}

func countChar(word1 string) map[byte]int {
	wordSet := map[byte]int{}
	for i := 0; i < len(word1); i++ {
		curChar := word1[i]
		wordSet[curChar]++
	}
	return wordSet
}

func haveSameCharacters(a, b map[byte]int) bool {
	for key := range a {
		if _, ok := b[key]; !ok {
			return false
		}
	}
	return true
}

func getMapSortedValues(m map[byte]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	slices.Sort(values)
	return values
}
