package main

// 推理
// - 能种的情况：
// 1. 遇到 0 才能种。
// 2. 0 紧挨着的不能为 1。
// 3. 题目已经说了花坛中没有紧挨着的 1，1 和 1 之间必须有 0。
// 4. 如果 x 地块为 1，那么左右一定是 0，可以检查 +2。
// 5. 如果 x 地块为 0，那么如果它左右没有 1，就可以种。
//
// 算法：
// - 从左扫到右，遇到一个能种的地块就 n-1，看扫完之后 n 是否为 0。
// - 遍历到 x 地块，如果 x 为 1，就 +2，如果 x 为 0，就检查左右是否为 0，为的话就种。
//
// 易错点：
// 1. 遇到 1 的时候不能 +2，因为循环自己也会 +1，最后变成 +3 了。
// 2. 种花的时候不能只 n--，还要修改花坛，不然下一个地块检测不到上一个地块的花。
// 3. 只判断 n 为 0 不稳定，如果不种完终止还继续种，n 会变成负数，最后不等于 0。
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		plot := flowerbed[i]
		if plot == 1 {
			i += 1
		} else if canPlace(flowerbed, i) { // plot == 0
			if placeOne(&n, flowerbed, i) {
				return true
			}
		}
	}
	return n == 0
}

// 返回值：如果还能继续种，返回 false；如果种完了，返回 true。
func placeOne(n *int, flowerbed []int, i int) bool {
	flowerbed[i] = 1
	(*n)--
	return *n <= 0
}

func canPlace(flowerbed []int, i int) bool {
	if i != 0 && flowerbed[i-1] != 0 {
		return false
	}
	if i != len(flowerbed)-1 && flowerbed[i+1] != 0 {
		return false
	}
	return true
}
