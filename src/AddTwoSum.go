package src

/*
题目：
给出两个 非空 的链表用来表示两个非负的整数。
它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

example：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	if l1 == nil && l2 == nil {
		return result
	}
	newNode := result
	a, b := 0, 0
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		newNode.Next = new(ListNode)
		newNode = newNode.Next
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		newNode.Val = (a + b + carry) % 10
		carry = (a + b + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result.Next
}