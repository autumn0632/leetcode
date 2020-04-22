package src

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

example:
input: 1->2->3->4->5, 2
output: 1->2->3->5

题解：哨兵节点
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{} // 哨兵节点
	result.Next = head
	cur := result
	var pre *ListNode

	i := 1
	for head != nil {
		// 慢指针开始移动，等head移动到最后一个，pre指向倒数第N个节点
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i ++
	}
	// 快指针一直移动，当快指针移动到最后一个，慢指针指向倒数第N个
	pre.Next = pre.Next.Next
	return result.Next
}
