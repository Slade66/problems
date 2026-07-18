package main

import "testing"

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "official example 1",
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			name: "official example 2 nested repetition",
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			name: "official example 3 adjacent groups and suffix",
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			name: "single plain character",
			s:    "a",
			want: "a",
		},
		{
			name: "multi digit repeat count",
			s:    "10[a]",
			want: "aaaaaaaaaa",
		},
		{
			name: "nested groups with surrounding letters",
			s:    "x2[y3[z]]w",
			want: "xyzzzyzzzw",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := decodeString1(tc.s)
			if got != tc.want {
				t.Errorf("decodeString(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
