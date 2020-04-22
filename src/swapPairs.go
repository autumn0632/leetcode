package src

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

example:
给定 1->2->3->4, 返回 2->1->4->3.

*/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	firstNode := head
	secondNode := head.Next

	// 交换
	firstNode.Next = SwapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}