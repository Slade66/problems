package main

import "testing"

func TestMaxVowels(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "official example 1",
			s:    "abciiidef",
			k:    3,
			want: 3,
		},
		{
			name: "official example 2",
			s:    "aeiou",
			k:    2,
			want: 2,
		},
		{
			name: "official example 3",
			s:    "leetcode",
			k:    3,
			want: 2,
		},
		{
			name: "single character vowel",
			s:    "a",
			k:    1,
			want: 1,
		},
		{
			name: "window length equals string length",
			s:    "aeiobcdf",
			k:    8,
			want: 4,
		},
		{
			name: "no vowels",
			s:    "rhythms",
			k:    4,
			want: 0,
		},
		{
			name: "best window at end",
			s:    "zzzaeiou",
			k:    5,
			want: 5,
		},
		{
			name: "single character windows choose a vowel",
			s:    "bcdua",
			k:    1,
			want: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxVowels1(tc.s, tc.k)
			if got != tc.want {
				t.Errorf("maxVowels(%q, %d) = %d, want %d", tc.s, tc.k, got, tc.want)
			}
		})
	}
}
