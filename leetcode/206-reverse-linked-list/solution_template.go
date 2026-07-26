package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 推理：遍历链表时，将当前节点的 Next 指向前一个节点；完成遍历后，前一个节点就是反转后的头节点。
func reverseList(head *ListNode) *ListNode {
	return nil
}
