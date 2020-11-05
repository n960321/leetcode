package e242_isAnagram

import (
	"testing"
)

/*

https://leetcode.com/problems/valid-anagram/

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := [26]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--
		if table[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func isAnagram_Aron(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[string]int{}
	tMap := map[string]int{}

	for i, _ := range s {
		sMap[string(s[i])] += 1
		tMap[string(t[i])] += 1
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {
		if vv, exist := tMap[k]; !exist || vv != v {
			return false
		}
	}

	return true
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "123",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "223",
			args: args{
				s: "car",
				t: "rat",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
