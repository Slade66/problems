package main

func deleteMiddle1(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	var prev *ListNode
	slow, fast := head, head

	// 旧写法：
	// for fast.Next != nil {
	// 	prev = slow
	// 	slow = slow.Next
	// 	if fast.Next.Next == nil {
	// 		break
	// 	}
	// 	fast = fast.Next.Next
	// }

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head
}
