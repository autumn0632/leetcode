package src

import (
	"fmt"
	"testing"
)

func TestTwoSum_1(t *testing.T) {
	nums := []int{1, 2, 6, 9, 19}
	target := 28

	res := TwoSum_1(nums, target)

	fmt.Println(nums, target, res)
}

func TestTwoSum_2(t *testing.T) {
	nums := []int{1, 2, 6, 9, 19}
	target := 28

	res := TwoSum_2(nums, target)

	fmt.Println(nums, target, res)
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(ThreeSum(nums))
}

func TestMergeTwoLists(t *testing.T) {
	l4 := &ListNode{9, nil}
	l3 := &ListNode{5, l4}
	l2 := &ListNode{3, l3}
	l := &ListNode{1, l2}

	m4 := &ListNode{8, nil}
	m3 := &ListNode{7, m4}
	m2 := &ListNode{4, m3}
	m := &ListNode{0, m2}

	n := MergeTwoLists(l, m)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
func TestRemoveNthFromEnd(t *testing.T) {
	l4 := &ListNode{9, nil}
	l3 := &ListNode{5, l4}
	l2 := &ListNode{3, l3}
	l := &ListNode{1, l2}

	n := RemoveNthFromEnd(l, 4)
	for n != nil {
		fmt.Printf("%d\t", n.Val)
		n = n.Next
	}
	fmt.Println()
}