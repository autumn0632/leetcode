package src

import "fmt"

/*
合并 k 个排序链表，返回合并后的排序链表。
请分析和描述算法的复杂度。

example:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

题解：
递归算法，

*/

func MergeNLists(lists []*ListNode) *ListNode {
	var temp *ListNode
	length := len(lists)
	for i := 0; i < length; i ++ {
		temp = MergeTwoListsK(temp, lists[i])
		new := temp
		for new != nil {
			fmt.Printf("%d \t", new.Val)
			new = new.Next
		}
		fmt.Println()
	}
	return temp
}

func MergeTwoListsK(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := new(ListNode)
	l := temp

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return l.Next
}