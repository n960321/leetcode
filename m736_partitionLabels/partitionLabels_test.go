package m736_partitionLabels

import (
	"fmt"
	"reflect"
	"testing"
)

/*
A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

Example 1:

Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.


Note:

S will have length in range [1, 500].
S will consist of lowercase English letters ('a' to 'z') only.
*/
func partitionLabels(S string) []int {
	start := 0
	end := 0
	res := []int{}
	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[c-'a'] = i
		fmt.Println(c)
	}
	fmt.Println('a')
	fmt.Println(lastIndex)
	for i, c := range S {
		end = max(end, lastIndex[c-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_partitionLabels(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				S: "ababcbacadefegdehijhklij",
			},
			want: []int{9, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
