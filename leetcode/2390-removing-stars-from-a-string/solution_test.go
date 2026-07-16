package main

import "testing"

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "official example 1",
			s:    "leet**cod*e",
			want: "lecoe",
		},
		{
			name: "official example 2",
			s:    "erase*****",
			want: "",
		},
		{
			name: "no stars",
			s:    "a",
			want: "a",
		},
		{
			name: "single pair is fully removed",
			s:    "a*",
			want: "",
		},
		{
			name: "star removes nearest character",
			s:    "ab*c",
			want: "ac",
		},
		{
			name: "consecutive stars",
			s:    "abc**d*",
			want: "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeStars1(tc.s)
			if got != tc.want {
				t.Errorf("removeStars(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
