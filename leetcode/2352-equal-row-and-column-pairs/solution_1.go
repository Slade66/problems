package main

import (
	"strconv"
	"strings"
)

// 推理：
// 在矩阵中找相同的行和列的数量。行和列只有组成元素一样、顺序一样，才一样。
// 最多 n * n 组。
// 可以想到的暴力法就是把每一列和每一行进行相比。
// 优化的版本是先计算出每一行和每一列的合，然后只去查那些和一样的行和列。
// 只是一个筛选条件，不是真正降低复杂度的优化：总和相同，但顺序不同，所以还得逐个元素比较。在最坏情况下，所有行列的总和都一样，你仍然需要比较 n × n 组。
//
// 哈希表：
// 先遍历每行，把每行的元素序列做成 key 存 hash 表，存出现次数。
// 再遍历每列，查 hash 表，有就 1 * value（1 列和 2 行都能匹配，是 1 x 2）。
//
// 复杂度：
// 时间：O(n^2)，遍历每行和每列
// 空间：O(n^2)，哈希表最多存 n 个行 key，每个 key 的长度是 O(n)
func equalPairs1(grid [][]int) int {
	count := 0

	rowCounts := make(map[string]int)
	for _, row := range grid {
		rowCounts[buildKey(row)]++
	}

	for col := 0; col < len(grid); col++ {
		colSlice := getColumn(grid, col)
		key := buildKey(colSlice)
		count += rowCounts[key]
	}

	return count
}

func buildKey(arr []int) string {
	var builder strings.Builder
	for _, num := range arr {
		builder.WriteString(strconv.Itoa(num))
		builder.WriteByte('_')
	}
	return builder.String()
}

func getColumn(grid [][]int, column int) []int {
	colSlice := make([]int, len(grid))
	for row := 0; row < len(grid); row++ {
		colSlice[row] = grid[row][column]
	}
	return colSlice
}
