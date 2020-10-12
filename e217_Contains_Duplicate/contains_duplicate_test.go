package e217_Contains_Duplicate

/*
https://leetcode.com/problems/contains-duplicate/
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false
Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		if _, exist := m[v]; !exist {
			m[v] = true
			continue
		}
		return true
	}
	return false
}
