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

// countSubstrings returns the number of palindromic substrings in the string.
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				count++
			}
		}
	}
	return count
}
