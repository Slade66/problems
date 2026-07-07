package main

// 推理：
// 最粗暴的方式，固定 2 个指针在前面的索引 0 和 1，然后第三个指针从索引 2 一直扫到尾。
// 然后把第二个指针从索引 1 移动到 2，第三个指针从 3 开始继续往后扫。
// 0 1 2 3 4
//
//	i j k
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}
