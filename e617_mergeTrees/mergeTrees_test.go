package e617_mergeTrees

/*
https://leetcode.com/problems/merge-two-binary-trees/
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:

Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7


Note: The merging process must start from the root nodes of both trees.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil && t2 != nil {
		return &TreeNode{
			Val:   t2.Val,
			Left:  mergeTrees(nil, t2.Left),
			Right: mergeTrees(nil, t2.Right),
		}
	}

	if t1 != nil && t2 == nil {
		return &TreeNode{
			Val:   t1.Val,
			Left:  mergeTrees(t1.Left, nil),
			Right: mergeTrees(t1.Right, nil),
		}
	}

	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}
