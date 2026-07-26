package main

// 推理：
// [1,2,3,4,5]
func reverseList1(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
