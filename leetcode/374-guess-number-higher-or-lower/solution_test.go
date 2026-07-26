package main

import "testing"

func TestGuessNumber(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		pick int
	}{
		{
			name: "official example 1",
			n:    10,
			pick: 6,
		},
		{
			name: "official example 2",
			n:    1,
			pick: 1,
		},
		{
			name: "official example 3",
			n:    2,
			pick: 1,
		},
		{
			name: "pick is lower bound",
			n:    100,
			pick: 1,
		},
		{
			name: "pick is upper bound",
			n:    100,
			pick: 100,
		},
		{
			name: "maximum n avoids integer overflow",
			n:    1<<31 - 1,
			pick: 1<<31 - 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pickedNumber = tc.pick

			got := guessNumber1(tc.n)
			if got != tc.pick {
				t.Errorf("guessNumber(%d) = %d, want %d", tc.n, got, tc.pick)
			}
		})
	}
}
