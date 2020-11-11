package e543_diameterOfBinaryTree

/*
https://leetcode.com/problems/diameter-of-binary-tree/

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	_, w := nextLevel(root, &result)
	if w > result {
		result = w
	}
	return result
}

func nextLevel(node *TreeNode, result *int) (d, w int) {
	if node == nil {
		return 0, 0
	}
	rLevel, rCount := nextLevel(node.Left, result)
	lLevel, lCount := nextLevel(node.Right, result)

	switch {
	case rCount > *result:
		*result = rCount
	case lCount > *result:
		*result = rCount
	}

	if rLevel > lLevel {
		return rLevel + 1, rLevel + lLevel
	} else {
		return lLevel + 1, rLevel + lLevel
	}
}
