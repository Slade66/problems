package main

// 推理：
//
// 题意理解：返回 2 个切片，第一个切片包含只存在于 nums1 的元素，第二个切片包含只存在于 nums2 的元素。
//
// 分别将两个数组转换为集合，再遍历集合求差集。
func findDifference1(nums1 []int, nums2 []int) [][]int {
	nums1Set := buildSet(nums1)
	nums2Set := buildSet(nums2)

	onlyInNums1 := make([]int, 0, len(nums1Set))
	onlyInNums2 := make([]int, 0, len(nums2Set))

	for num := range nums1Set {
		if _, exists := nums2Set[num]; !exists {
			onlyInNums1 = append(onlyInNums1, num)
		}
	}

	for num := range nums2Set {
		if _, exists := nums1Set[num]; !exists {
			onlyInNums2 = append(onlyInNums2, num)
		}
	}

	return [][]int{onlyInNums1, onlyInNums2}
}

func buildSet(nums []int) map[int]bool {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		set[num] = true
	}

	return set
}
