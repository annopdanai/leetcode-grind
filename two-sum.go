package leetcodeblind75

func twoSum(nums []int, target int) []int {
	// the outer loop starts from the first element of the array
	for i := 0; i < len(nums); i++ {
		// the inner loop starts from the element after the current element
		for j := i + 1; j < len(nums); j++ {
			// if the sum of the current element and the next element is equal to the target
			if nums[i]+nums[j] == target {
				// return the current and next element
				return []int{i, j}
			}
		}
	}

	return []int{}
}
