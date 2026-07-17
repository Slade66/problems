package main

// 推理：
// 第二种写法。
func asteroidCollision2(asteroids []int) []int {
	var stack []int

	for i := 0; i < len(asteroids); i++ {
		currAsteroid := asteroids[i]
		alive := true
		// 如果栈里还有元素，小行星还活着，且发生了碰撞
		for j := len(stack) - 1; j >= 0 && alive && isCollide2(stack[j], currAsteroid); j-- {
			switch {
			case stack[j] == -currAsteroid:
				stack = stack[:j]
				alive = false
			case stack[j] > -currAsteroid:
				alive = false
			case stack[j] < -currAsteroid:
				stack = stack[:j]
			}
		}
		if alive {
			stack = append(stack, currAsteroid)
		}
	}

	return stack
}

func isCollide2(asteroidA, asteroidB int) bool {
	return asteroidA > 0 && asteroidB < 0
}
