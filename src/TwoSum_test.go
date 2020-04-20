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
