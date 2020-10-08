package main

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/house-robber/
func TestHouseRobber(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
func rob(num []int) int {
	preMax := 0
	currentMax := 0

	for i := 0; len(num) > i; i++ {
		temp := currentMax
		if preMax+num[i] > currentMax {
			currentMax = preMax + num[i]
		}
		preMax = temp
	}
	return currentMax
}

// v1
// func rob(nums []int) int {
// 	a := 0
// 	b := 0
// 	for i := 0; len(nums) > i; i++ {
// 		if i%2 > 0 {
// 			a += nums[i]
// 		} else {
// 			b += nums[i]
// 		}
// 	}
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
