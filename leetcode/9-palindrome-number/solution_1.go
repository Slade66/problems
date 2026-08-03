package main

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	if x != 0 && x%10 == 0 {
		return false
	}

	half := 0

	for x > half {
		last := x % 10
		half = half*10 + last
		x = x / 10
	}

	return (x == half) || (x == half/10)
}
