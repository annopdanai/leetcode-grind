package leetcodeblind75

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	// Sort the numbers in ascending order
	sort.Ints(nums)

	// Iterate through the numbers
	for i := 0; i < len(nums)-1; i++ {
		// If the current number equals the next number, return true
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
