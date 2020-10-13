package m2_addTwoNumbers

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func add(l1, l2 *ListNode, val int) *ListNode {
	result := &ListNode{}
	// quotient 商數
	// remainder 餘數
	t := 0
	quotient := 0

	if l1 != nil && l2 != nil {
		t = l1.Val + l2.Val + val
		result.Val = t % 10
		quotient = t / 10
		result.Next = add(l1.Next, l2.Next, quotient)
	} else if l1 != nil && l2 == nil {
		t = l1.Val + val
		result.Val = t % 10
		quotient = t / 10
		result.Next = add(l1.Next, nil, quotient)
	} else if l1 == nil && l2 != nil {
		t = l2.Val + val
		result.Val = t % 10
		quotient = t / 10
		result.Next = add(nil, l2.Next, quotient)
	} else if l1 == nil && l2 == nil && val > 0 {
		result.Val = val
		return result
	} else {
		return nil
	}

	return result
}
