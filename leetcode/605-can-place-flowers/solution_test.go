package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name:      "official example 1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "official example 2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			name:      "zero new flowers always fits",
			flowerbed: []int{1},
			n:         0,
			want:      true,
		},
		{
			name:      "single empty plot",
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		{
			name:      "single planted plot",
			flowerbed: []int{1},
			n:         1,
			want:      false,
		},
		{
			name:      "can use both edges",
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         2,
			want:      true,
		},
		{
			name:      "can stop after one edge placement",
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
			want:      true,
		},
		{
			name:      "long empty run can place alternating flowers",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			want:      true,
		},
		{
			name:      "middle gap too short for two flowers",
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			flowerbed := append([]int(nil), tc.flowerbed...)
			got := canPlaceFlowers(flowerbed, tc.n)
			if got != tc.want {
				t.Errorf("canPlaceFlowers(%v, %d) = %t, want %t", tc.flowerbed, tc.n, got, tc.want)
			}
		})
	}
}
