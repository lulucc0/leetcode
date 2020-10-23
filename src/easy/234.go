package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	nums := []int{}
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}
