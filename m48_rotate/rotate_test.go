package m48_rotate

/*
https://leetcode.com/problems/rotate-image/

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.



Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
Example 2:


Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
Example 3:

Input: matrix = [[1]]
Output: [[1]]
Example 4:

Input: matrix = [[1,2],[3,4]]
Output: [[3,1],[4,2]]

*/

func rotate_V1(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for ii := range matrix[i] {
			newMatrix[ii][n-1-i] = matrix[i][ii]

		}
	}
	copy(matrix, newMatrix)
}

func rotate(matrix [][]int) {
	// m'(i,j) = m(n-1-j, i)
	// layer by layer
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-1-i; j++ {
			x, y, temp := i, j, matrix[i][j]
			// rotate four numbers
			for {
				if n-1-y == i && x == j {
					matrix[x][y] = temp
					break
				}
				matrix[x][y] = matrix[n-1-y][x]
				x, y = n-1-y, x
			}

		}
	}
}
