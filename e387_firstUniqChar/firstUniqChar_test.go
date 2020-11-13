package e387_firstUniqChar

import "strings"

func firstUniqChar(s string) int {
	for i := range s {
		if strings.LastIndex(s, s[i:i+1]) == strings.Index(s, s[i:i+1]) {
			return i
		}
	}
	return -1

}
