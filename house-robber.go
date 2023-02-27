package leetcodeblind75

import (
	"math"
)

func rob(nums []int) int {
	// a = max money if we rob current house
	a := 0
	// b = max money if we don't rob current house
	b := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			// if we rob current house, we cannot rob previous house
			// so we take max of money from previous house and current house
			a = int(math.Max(float64(a+nums[i]), float64(b)))
		} else {
			// if we don't rob current house, we can take max of money from previous house
			b = int(math.Max(float64(a), float64(b+nums[i])))
		}
	}

	// return the max of money we can get when we rob the last house
	return int(math.Max(float64(a), float64(b)))
}
