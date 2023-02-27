package leetcodeblind75

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
