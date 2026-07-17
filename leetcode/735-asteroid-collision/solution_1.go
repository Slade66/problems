package main

// 推理：
// 一行小行星，有的向左，有的向右，速度相同。
// 数组元素的值是大小，正负号是方向。
// 计算出碰撞后的结果：
// 1. 同向不撞，反向炸。
// 2. 大小撞，小的炸。
// 3. 一样大，都炸。
//
// 一个个遍历加入结果集，处理遍历到的那一个，查看它跟上一个的大小方向关系。
//
// 易错点：
// 1. 不能只比数值大小，因为符号只是方向，要比较绝对值。
// 2. 当前的小行星比上一个大，也可能比上一个的上一个大，所以要依次往上找出来碾碎。
// 3. 有个大行星可以保护后面的小行星。
// 4. 方向相反不会碰撞。只有左边向右、右边向左时才会碰撞。
//
// 最后数组中都是同向的，或清空。
func asteroidCollision1(asteroids []int) []int {
	var result []int
	result = append(result, asteroids[0])

OuterLoop:
	for i := 1; i < len(asteroids); i++ {
		asteroid := asteroids[i]
		for j := len(result) - 1; j >= 0 && isCollide(result[j], asteroid); j-- {
			compareResult := compareAsteroid(result[j], asteroid)
			switch {
			case compareResult == 0:
				result = result[:j]
				continue OuterLoop
			case compareResult > 0:
				continue OuterLoop
			case compareResult < 0:
				result = result[:j]
			}
		}
		result = append(result, asteroid)
	}
	return result
}

func isCollide(asteroidA, asteroidB int) bool {
	return asteroidA > 0 && asteroidB < 0
}

// == 0: asteroidA == asteroidB
// > 0: asteroidA > asteroidB
// < 0: asteroidA < asteroidB
func compareAsteroid(asteroidA, asteroidB int) int {
	return abs(asteroidA) - abs(asteroidB)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
