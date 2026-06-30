package main

import "strings"

// 推理：
// - 如果一个字符串 a 能除字符串 b，那么 b = a + a + ... 即 b 是由 a 重复多次拼接而成。
// - 如果一个字符串 c 能除 a 和 b，那么 c 是 a 和 b 的公共基串（可以重复多次拼出两个字符串的基础字符串）。
// - a = c + c + ...
// - b = c + c + ...
// - 左侧都以 c 开头，因此，公共基串 c 必须是 a 和 b 的前缀。
// - 而且，a 的长度是 c 的 N 倍，b 的长度是 c 的 M 倍。因为要完整重复拼接，所以候选字符串的长度必须同时整除两个字符串的长度。也就是说，c 的长度就是 a 和 b 长度的公因数。
//
// 算法：找最长的公共基串 c，它首先是 a 和 b 的公因数，并且将它复制 n 次能得到 a 和 b。
func gcdOfStrings(str1 string, str2 string) string {
	str1Len, str2Len := len(str1), len(str2)
	gcds := findGCDs(str1Len, str2Len)
	largestGcdString := 0
	for _, gcd := range gcds {
		if isGcdOfString(str1, str2, gcd) {
			largestGcdString = gcd
		}
	}
	return str1[:largestGcdString]
}

func findGCDs(a, b int) []int {
	minVal := min(a, b)
	result := make([]int, 0, minVal)
	for i := 1; i <= minVal; i++ {
		if a%i == 0 && b%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func isGcdOfString(str1, str2 string, gcd int) bool {
	s := str1[:gcd]
	return canBuild(str1, s, gcd) && canBuild(str2, s, gcd)
}

func canBuild(target, base string, gcd int) bool {
	times := len(target) / gcd
	return strings.Repeat(base, times) == target
}
