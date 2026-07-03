package main

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "official example 1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "official example 2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "official example 3",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "single word",
			s:    "single",
			want: "single",
		},
		{
			name: "only one word with spaces",
			s:    "   hello   ",
			want: "hello",
		},
		{
			name: "digits and multiple spaces",
			s:    "1  2   3",
			want: "3 2 1",
		},
		{
			name: "many words with leading and trailing spaces",
			s:    "  alpha   beta gamma  ",
			want: "gamma beta alpha",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reverseWords(tc.s)
			if got != tc.want {
				t.Errorf("reverseWords(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
