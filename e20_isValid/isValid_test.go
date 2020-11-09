package e20_isValid

import "testing"

/*
https://leetcode.com/problems/valid-parentheses/
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	left := map[string]bool{"{": true, "[": true, "(": true}
	right := map[string]bool{"}": true, "]": true, ")": true}
	matchMap := map[string]string{"{": "}", "[": "]", "(": ")", "}": "{", "]": "[", ")": "("}
	arr := make([]string, 0, len(s))
	getLast := func(arr []string) string {
		if len(arr) == 0 {
			return ""
		}
		return arr[len(arr)-1]
	}

	removeArrLast := func(arr []string) []string {
		if len(arr) == 0 {
			return arr
		}
		return arr[0 : len(arr)-1]
	}

	for _, v := range s {
		vv := string(v)
		if _, exist := left[vv]; exist {
			arr = append(arr, vv)
			continue
		}

		if _, exist := right[vv]; exist {
			last := getLast(arr)
			if matchMap[last] == vv {
				arr = removeArrLast(arr)
				continue
			}
			return false
		}

	}

	if len(arr) != 0 {
		return false
	}
	return true
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s: "(",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
