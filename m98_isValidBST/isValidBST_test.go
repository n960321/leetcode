package m98_isValidBST

/*
https://leetcode.com/problems/validate-binary-search-tree/

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:

    2
   / \
  1   3

Input: [2,1,3]
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode = nil
	return isBSTHelper(&pre, root)
}

func isBSTHelper(pre **TreeNode, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBSTHelper(pre, root.Left) {
		return false
	}
	if *pre != nil && (*pre).Val >= root.Val {
		return false
	}
	*pre = root
	return isBSTHelper(pre, root.Right)
}
