package e21_mergeTwoLists

/*
https://leetcode.com/problems/merge-two-sorted-lists/

Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

Example 1:


Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: l1 = [], l2 = []
Output: []
Example 3:

Input: l1 = [], l2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both l1 and l2 are sorted in non-decreasing order.
*/

/**
 * Definition for singly-linked list.

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists(nil, l2.Next),
		}
	}

	if l1 != nil && l2 == nil {
		return &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists(l1.Next, nil),
		}
	}

	if l1.Val <= l2.Val {
		return &ListNode{
			Val:  l1.Val,
			Next: mergeTwoLists(l1.Next, l2),
		}
	}

	if l1.Val > l2.Val {
		return &ListNode{
			Val:  l2.Val,
			Next: mergeTwoLists(l1, l2.Next),
		}
	}

	return nil
}
