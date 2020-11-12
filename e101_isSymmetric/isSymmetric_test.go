package e101_isSymmetric

/*

https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(left, right *TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true
	case left == nil && right != nil, left != nil && right == nil:
		return false
	default:
		if left.Val != right.Val {
			return false
		} else {
			return bfs(left.Left, right.Right) && bfs(left.Right, right.Left)
		}
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return bfs(root.Left, root.Right)
}

/*
func isSymmetric(root *TreeNode) bool {
	res := [][]*int{}
	dfs(root, 0, &res)

	for i, v := range res {
		if i == 0 || i == 1 {
			continue
		}
		for ii := 0; ii < len(v)/2; ii++ {
			if v[ii] == nil && v[len(v)-1-ii] == nil {
				continue
			}

			if v[ii] != nil && v[len(v)-1-ii] == nil {
				return false
			}

			if v[ii] == nil && v[len(v)-1-ii] != nil {
				return false
			}

			if *v[ii] != *v[len(v)-1-ii] {
				return false
			}
		}
	}
	return true
}

func dfs(node *TreeNode, level int, res *[][]*int) {

	currentLevel := level + 1

	for len(*res) <= currentLevel {
		*res = append(*res, []*int{})
	}

	if node == nil {
		(*res)[currentLevel] = append((*res)[currentLevel], nil)
		return
	}
	(*res)[currentLevel] = append((*res)[currentLevel], &node.Val)

	dfs(node.Left, currentLevel, res)
	dfs(node.Right, currentLevel, res)
}
*/
