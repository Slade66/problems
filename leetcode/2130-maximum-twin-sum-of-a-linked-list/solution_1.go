package main

func pairSum1(head *ListNode) int {

	// 1. 寻找中点，分割链表为前后两半。
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 反转后半部分链表
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next // 先存起来，避免等会修改完 curr.next 找不到下一个节点。
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 3. 同步遍历前后两半链表，计算最大和。
	maxSum := 0
	firstHalf := head
	secondHalf := prev
	for firstHalf != nil && secondHalf != nil {
		maxSum = max(maxSum, firstHalf.Val+secondHalf.Val)
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return maxSum
}
