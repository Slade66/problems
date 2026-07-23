package main

func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		nextOdd := even.Next
		odd.Next = nextOdd
		odd = nextOdd

		nextEven := odd.Next
		even.Next = nextEven
		even = nextEven
	}

	odd.Next = evenHead
	return head
}
