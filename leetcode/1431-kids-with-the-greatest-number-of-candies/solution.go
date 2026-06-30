package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := pickMax(candies)
	result := make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = candy+extraCandies >= max
	}
	return result
}

func pickMax(candies []int) int {
	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	return max
}
