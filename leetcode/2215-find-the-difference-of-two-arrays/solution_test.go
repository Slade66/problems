package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDifference(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{
			name:  "official example 1",
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
			want:  [][]int{{1, 3}, {4, 6}},
		},
		{
			name:  "official example 2",
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			want:  [][]int{{3}, {}},
		},
		{
			name:  "all values overlap",
			nums1: []int{-1000, 0, 1000, 1000},
			nums2: []int{1000, 0, -1000, -1000},
			want:  [][]int{{}, {}},
		},
		{
			name:  "disjoint values at constraint boundaries",
			nums1: []int{-1000},
			nums2: []int{1000},
			want:  [][]int{{-1000}, {1000}},
		},
		{
			name:  "duplicate values appear only once",
			nums1: []int{-1, -1, 0, 0, 2, 2},
			nums2: []int{1, 1, 2, 2},
			want:  [][]int{{-1, 0}, {1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findDifference1(tc.nums1, tc.nums2)
			if !equalAnswers(got, tc.want) {
				t.Errorf("findDifference(%v, %v) = %v, want %v (order within each list does not matter)", tc.nums1, tc.nums2, got, tc.want)
			}
		})
	}
}

func equalAnswers(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := range want {
		gotValues := append([]int(nil), got[i]...)
		wantValues := append([]int(nil), want[i]...)
		sort.Ints(gotValues)
		sort.Ints(wantValues)

		if !reflect.DeepEqual(gotValues, wantValues) {
			return false
		}
	}

	return true
}
