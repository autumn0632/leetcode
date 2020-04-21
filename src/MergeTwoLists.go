package src

/*
Merge Two Sorted Lists:
合并两个有序链表，

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
example：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

方法：使用哨兵节点，
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := new(ListNode)  // 哨兵节点
	result := preHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		}else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}

	return result.Next
}
