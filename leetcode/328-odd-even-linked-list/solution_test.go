package main

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	testCases := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "official example 1",
			head: []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			name: "official example 2",
			head: []int{2, 1, 3, 5, 6, 4, 7},
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			name: "empty list",
			head: nil,
			want: nil,
		},
		{
			name: "single node",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "two nodes",
			head: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "even number of nodes",
			head: []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 3, 5, 2, 4, 6},
		},
		{
			name: "repeated values",
			head: []int{7, 7, 7, 7, 7},
			want: []int{7, 7, 7, 7, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := listValues(oddEvenList1(buildList(tc.head)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("oddEvenList(%v) = %v, want %v", tc.head, got, tc.want)
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

func listValues(head *ListNode) []int {
	if head == nil {
		return nil
	}

	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}

	return values
}
