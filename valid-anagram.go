package leetcodeblind75

import (
	"sort"
	"strings"
)

// isAnagram returns true if two strings are anagrams, false otherwise
func isAnagram(s string, t string) bool {
	// split each string into a slice of characters
	first := strings.Split(s, "")
	second := strings.Split(t, "")

	// sort the characters in the slices
	sort.Strings(first)
	sort.Strings(second)

	// join the characters back into strings
	// and compare the strings
	return strings.Join(first, "") == strings.Join(second, "")
}
