package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{
			name:  "same length",
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			name:  "word2 is longer",
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			name:  "word1 is longer",
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
		{
			name:  "single character strings",
			word1: "a",
			word2: "b",
			want:  "ab",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := mergeAlternately(tc.word1, tc.word2)
			if got != tc.want {
				t.Errorf("mergeAlternately(%q, %q) = %q, want %q", tc.word1, tc.word2, got, tc.want)
			}
		})
	}
}
