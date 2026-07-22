package main

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	testCases := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "official example 1",
			head: []int{1, 3, 4, 7, 1, 2, 6},
			want: []int{1, 3, 4, 1, 2, 6},
		},
		{
			name: "official example 2",
			head: []int{1, 2, 3, 4},
			want: []int{1, 2, 4},
		},
		{
			name: "official example 3",
			head: []int{2, 1},
			want: []int{2},
		},
		{
			name: "single node",
			head: []int{1},
			want: []int{},
		},
		{
			name: "three nodes",
			head: []int{1, 2, 3},
			want: []int{1, 3},
		},
		{
			name: "five nodes",
			head: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "repeated values",
			head: []int{7, 7, 7, 7},
			want: []int{7, 7, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := deleteMiddle1(buildList(tc.head))
			gotValues := listValues(got)
			if !reflect.DeepEqual(gotValues, tc.want) {
				t.Errorf("deleteMiddle(%v) = %v, want %v", tc.head, gotValues, tc.want)
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
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	return values
}
