package m448_findDisappearedNumbers

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

*/
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	res := make([]int, length)
	for _, v := range nums {
		res[v-1] = v
	}
	j := 0
	for i := 0; i < length; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

func onePass(nums []int) []int {
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i])) - 1)
		fmt.Println(index)
		if nums[index] > 0 {
			nums[index] = 0 - nums[index]
		}
		fmt.Println(nums[index])
	}
	fmt.Println(nums)
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			fmt.Println("=", nums[i])
			res = append(res, i+1)
		}
	}
	return res
}

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			want: []int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
