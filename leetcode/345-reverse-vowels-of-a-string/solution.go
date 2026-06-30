package main

// 算法：
// - 双指针，一个从左往右，一个从右往左，各自停在元音字母上交换，直到相遇。
func reverseVowels(s string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	s_bytes := []byte(s)
	for left, right := 0, len(s)-1; left < right; {
		var leftChar *byte
		for left < right {
			leftChar = &s_bytes[left]
			if _, ok := vowels[*leftChar]; ok {
				break
			}
			left++
		}
		if left >= right {
			break
		}
		var rightChar *byte
		for right > left {
			rightChar = &s_bytes[right]
			if _, ok := vowels[*rightChar]; ok {
				break
			}
			right--
		}
		if right <= left {
			break
		}
		swap(leftChar, rightChar)
		left++
		right--
	}
	return string(s_bytes)
}

func swap(a, b *byte) {
	*a, *b = *b, *a
}

// 易错点：
// 1. 循环内的 leftChar = &s_bytes[left] 不要写成 leftChar := &s_bytes[left]，否则就不是在给外层 leftChar 赋值，而是在内层循环里新建了一个同名变量。循环 break 后，外层的 leftChar 仍然是 nil。rightChar 也是如此。到 swap(leftChar, rightChar) 实际上传进去的是两个 nil 指针，最后在 swap 里解引用就 panic 了。
// 
// 优化点：
// 1. swap(leftChar, rightChar) 可以换成 s_bytes[left], s_bytes[right] = s_bytes[right], s_bytes[left]。这是 Go 的多重赋值写法，用来一次性交换两个位置的值。Go 在执行多重赋值时，会先计算右边所有表达式的值，再统一赋给左边。所以不会出现第一个赋值把原值覆盖掉、导致第二个赋值拿不到旧值的问题。
