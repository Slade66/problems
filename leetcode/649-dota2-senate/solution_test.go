package main

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	testCases := []struct {
		name   string
		senate string
		want   string
	}{
		{
			name:   "official example 1",
			senate: "RD",
			want:   "Radiant",
		},
		{
			name:   "official example 2",
			senate: "RDD",
			want:   "Dire",
		},
		{
			name:   "single radiant senator",
			senate: "R",
			want:   "Radiant",
		},
		{
			name:   "single dire senator",
			senate: "D",
			want:   "Dire",
		},
		{
			name:   "dire acts first",
			senate: "DR",
			want:   "Dire",
		},
		{
			name:   "alternating senators",
			senate: "RDRD",
			want:   "Radiant",
		},
		{
			name:   "radiant survives next round",
			senate: "RRDDD",
			want:   "Radiant",
		},
		{
			name:   "dire survives multiple rounds",
			senate: "DRDRR",
			want:   "Dire",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := predictPartyVictory1(tc.senate)
			if got != tc.want {
				t.Errorf("predictPartyVictory(%q) = %q, want %q", tc.senate, got, tc.want)
			}
		})
	}
}
