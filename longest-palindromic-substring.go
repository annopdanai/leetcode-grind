package leetcodeblind75

// isPalindrome returns true if the string is a palindrome.
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// longestPalindrome returns the longest palindrome substring
// of a string.
func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > len(longest) {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}
