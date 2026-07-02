package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "official example 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "official example 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "empty s is subsequence",
			s:    "",
			t:    "abc",
			want: true,
		},
		{
			name: "both empty",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "non-empty s with empty t",
			s:    "a",
			t:    "",
			want: false,
		},
		{
			name: "same string",
			s:    "abc",
			t:    "abc",
			want: true,
		},
		{
			name: "relative order matters",
			s:    "aec",
			t:    "abcde",
			want: false,
		},
		{
			name: "repeated characters need enough matches",
			s:    "aaa",
			t:    "aa",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isSubsequence(%q, %q) = %t, want %t", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
