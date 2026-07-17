package main

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	testCases := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{
			name:      "official example 1",
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			name:      "official example 2",
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			name:      "official example 3",
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
		{
			name:      "official example 4",
			asteroids: []int{3, 5, -6, 2, -1, 4},
			want:      []int{-6, 2, 4},
		},
		{
			name:      "asteroids moving away from each other",
			asteroids: []int{-2, 1},
			want:      []int{-2, 1},
		},
		{
			name:      "all asteroids moving right",
			asteroids: []int{1, 2, 3},
			want:      []int{1, 2, 3},
		},
		{
			name:      "all asteroids moving left",
			asteroids: []int{-3, -2, -1},
			want:      []int{-3, -2, -1},
		},
		{
			name:      "left-moving asteroid survives multiple collisions",
			asteroids: []int{1, 2, -5},
			want:      []int{-5},
		},
		{
			name:      "left-moving asteroid is destroyed",
			asteroids: []int{5, -1},
			want:      []int{5},
		},
		{
			name:      "chain of equal-sized collisions",
			asteroids: []int{1, 3, 2, -2, -3},
			want:      []int{1},
		},
		{
			name:      "ignore asteroids already destroyed",
			asteroids: []int{-2, 1, -2, -2},
			want:      []int{-2, -2, -2},
		},
		{
			name:      "maximum asteroid sizes collide",
			asteroids: []int{1000, -1000},
			want:      []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := asteroidCollision2(tc.asteroids)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("asteroidCollision(%v) = %v, want %v", tc.asteroids, got, tc.want)
			}
		})
	}
}
