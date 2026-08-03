package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "official example 1",
			x:    121,
			want: true,
		},
		{
			name: "official example 2",
			x:    -121,
			want: false,
		},
		{
			name: "official example 3",
			x:    10,
			want: false,
		},
		{
			name: "zero is a palindrome",
			x:    0,
			want: true,
		},
		{
			name: "single digit is a palindrome",
			x:    7,
			want: true,
		},
		{
			name: "even number of digits",
			x:    1221,
			want: true,
		},
		{
			name: "odd number of digits",
			x:    12321,
			want: true,
		},
		{
			name: "odd number of digits but not a palindrome",
			x:    213,
			want: false,
		},
		{
			name: "trailing zero makes a number invalid",
			x:    100,
			want: false,
		},
		{
			name: "minimum 32 bit integer",
			x:    -2147483648,
			want: false,
		},
		{
			name: "maximum 32 bit integer",
			x:    2147483647,
			want: false,
		},
		{
			name: "palindrome near maximum 32 bit integer",
			x:    2147447412,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome1(tc.x)
			if got != tc.want {
				t.Errorf("isPalindrome(%d) = %t, want %t", tc.x, got, tc.want)
			}
		})
	}
}
