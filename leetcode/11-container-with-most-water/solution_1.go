package main

// 推理：
// 题意理解：找到两条线 a 和 b，使它们和横轴围成的长方形的面积最大。
// 盛水量 = 横轴 x min(height[a], height[b])
// 所以肯定是从两边往中间扫，因为这样可以使得 x 尽可能长。
//
// 当前容器的水位由较短的边决定。每次计算完当前面积后，移动较短的那条边。
// 假设现在左右指针是 a 和 b，并且：height[a] < height[b]
// 当前面积是：(b - a) * height[a]
// 因为短板是 height[a]，只要左边这根短边 a 不动，右边怎么换，都不可能比当前这个面积更大。
// 所以当前面积已经记录进 maxArea 之后，a 这根边就可以淘汰了。下一步只能移动 a，去赌一个更高的左边界。
// 如果右边更短，同理淘汰 b。
func maxArea1(height []int) int {
	a, b := 0, len(height)-1
	maxArea := 0
	for a < b {
		x := b - a
		y := min(height[a], height[b])
		maxArea = max(maxArea, x*y)
		if height[a] < height[b] {
			a++
		} else {
			b--
		}
	}
	return maxArea
}
