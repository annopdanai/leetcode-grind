package leetcodeblind75

import (
	"regexp"
	"strings"
)

// clearStringAndSpace clears the string of all non-alphanumeric characters and spaces.
func clearStringAndSpace(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(str, "")
}

// isPalindrome returns true if the string is a palindrome.
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
