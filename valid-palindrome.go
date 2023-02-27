package leetcodeblind75

import (
	"regexp"
	"strings"
)

func clearStringAndSpace(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(str, "")
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = clearStringAndSpace(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
