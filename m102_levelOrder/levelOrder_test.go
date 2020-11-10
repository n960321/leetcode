package m102_levelOrder

/*

https://leetcode.com/problems/binary-tree-level-order-traversal/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, -1, &res)
	return res
}

func dfs(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	currentLevel := level + 1
	for len(*res) <= currentLevel {
		*res = append(*res, []int{})
	}

	(*res)[currentLevel] = append((*res)[currentLevel], node.Val)

	dfs(node.Left, currentLevel, res)
	dfs(node.Right, currentLevel, res)

}

// func nextLevel(left, right *TreeNode) [][]int {
// 	arr := []int{}

// 	if left != nil {
// 		arr = append(arr, left.Val)
// 	}
// 	if right != nil {
// 		arr = append(arr, right.Val)
// 	}

// 	result := [][]int{}
// 	result = append(result, arr)
// 	if left != nil {
// 		result = append(result, nextLevel(left.Left, left.Right)...)
// 	}

// 	if right != nil {
// 		result = append(result, nextLevel(right.Left, left.Right)...)
// 	}
// 	return result
// }
