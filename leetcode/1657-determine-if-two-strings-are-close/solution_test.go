package main

import "testing"

func TestCloseStrings(t *testing.T) {
	testCases := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		{
			name:  "official example 1",
			word1: "abc",
			word2: "bca",
			want:  true,
		},
		{
			name:  "official example 2",
			word1: "a",
			word2: "aa",
			want:  false,
		},
		{
			name:  "official example 3",
			word1: "cabbba",
			word2: "abbccc",
			want:  true,
		},
		{
			name:  "single equal character",
			word1: "z",
			word2: "z",
			want:  true,
		},
		{
			name:  "different character sets with equal frequencies",
			word1: "aabbcc",
			word2: "ddeeff",
			want:  false,
		},
		{
			name:  "same character set with different frequency multisets",
			word1: "aabbccc",
			word2: "abbcccc",
			want:  false,
		},
		{
			name:  "same character set with permuted frequencies",
			word1: "aaabbbbcc",
			word2: "aabbbbccc",
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := closeStrings1(tc.word1, tc.word2)
			if got != tc.want {
				t.Errorf("closeStrings(%q, %q) = %t, want %t", tc.word1, tc.word2, got, tc.want)
			}
		})
	}
}
