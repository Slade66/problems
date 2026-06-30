package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			name:         "official example 1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			name:         "official example 2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			name:         "official example 3",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := kidsWithCandies(tc.candies, tc.extraCandies)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("kidsWithCandies(%v, %d) = %v, want %v", tc.candies, tc.extraCandies, got, tc.want)
			}
		})
	}
}
