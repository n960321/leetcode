package m78_subsets

import (
	"fmt"
	"reflect"
	"testing"
)

/*

https://leetcode.com/problems/subsets/

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func subsets(nums []int) [][]int {
	subsets := [][]int{[]int{}}
	for _, num := range nums {
		for _, subset := range subsets {
			fmt.Printf("===== start s: %+v  subset: %+v\n", subsets, subset)
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
			fmt.Printf("===== %+v\n", subsets)
		}
		fmt.Printf("==============\n")
	}
	return subsets
}

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{3},
				{1},
				{2},
				{1, 2, 3},
				{1, 3},
				{2, 3},
				{1, 2},
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
