package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "official example 1",
			head: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "official example 2",
			head: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "official example 3",
			head: nil,
			want: nil,
		},
		{
			name: "single node",
			head: []int{42},
			want: []int{42},
		},
		{
			name: "negative values",
			head: []int{-5000, 0, 5000},
			want: []int{5000, 0, -5000},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := listValues(reverseList1(buildList(tc.head)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("reverseList(%v) = %v, want %v", tc.head, got, tc.want)
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
