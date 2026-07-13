package main

import "testing"

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		name string
		gain []int
		want int
	}{
		{
			name: "official example 1",
			gain: []int{-5, 1, 5, 0, -7},
			want: 1,
		},
		{
			name: "official example 2",
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			want: 0,
		},
		{
			name: "single positive gain",
			gain: []int{5},
			want: 5,
		},
		{
			name: "single negative gain",
			gain: []int{-5},
			want: 0,
		},
		{
			name: "highest altitude occurs at the end",
			gain: []int{-2, 3, 4},
			want: 5,
		},
		{
			name: "highest altitude occurs in the middle",
			gain: []int{3, -2, -5, 4},
			want: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := largestAltitude1(tc.gain)
			if got != tc.want {
				t.Errorf("largestAltitude(%v) = %d, want %d", tc.gain, got, tc.want)
			}
		})
	}
}
