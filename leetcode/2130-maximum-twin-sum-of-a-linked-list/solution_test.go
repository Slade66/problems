package main

import "testing"

func TestPairSum(t *testing.T) {
	testCases := []struct {
		name string
		head []int
		want int
	}{
		{
			name: "official example 1",
			head: []int{5, 4, 2, 1},
			want: 6,
		},
		{
			name: "official example 2",
			head: []int{4, 2, 2, 3},
			want: 7,
		},
		{
			name: "official example 3",
			head: []int{1, 100000},
			want: 100001,
		},
		{
			name: "minimum length and value",
			head: []int{1, 1},
			want: 2,
		},
		{
			name: "repeated values",
			head: []int{5, 5, 5, 5, 5, 5},
			want: 10,
		},
		{
			name: "maximum sum at outer pair",
			head: []int{100000, 1, 1, 1, 1, 100000},
			want: 200000,
		},
		{
			name: "maximum sum at inner pair",
			head: []int{1, 2, 100000, 100000, 2, 1},
			want: 200000,
		},
		{
			name: "multiple pairs share maximum",
			head: []int{9, 8, 7, 6},
			want: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := pairSum1(buildList(tc.head))
			if got != tc.want {
				t.Errorf("pairSum(%v) = %d, want %d", tc.head, got, tc.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}

	return dummy.Next
}
