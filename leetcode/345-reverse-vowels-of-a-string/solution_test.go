package main

import "testing"

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "official example 1",
			s:    "IceCreAm",
			want: "AceCreIm",
		},
		{
			name: "official example 2",
			s:    "leetcode",
			want: "leotcede",
		},
		{
			name: "no vowels",
			s:    "bcdfg",
			want: "bcdfg",
		},
		{
			name: "single uppercase vowel",
			s:    "A",
			want: "A",
		},
		{
			name: "mixed case vowels",
			s:    "aA",
			want: "Aa",
		},
		{
			name: "all vowels",
			s:    "aeiou",
			want: "uoiea",
		},
		{
			name: "punctuation and vowels",
			s:    ".,aE!",
			want: ".,Ea!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reverseVowels(tc.s)
			if got != tc.want {
				t.Errorf("reverseVowels(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
