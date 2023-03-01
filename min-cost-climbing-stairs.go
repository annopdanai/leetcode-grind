package leetcodeblind75

import (
	"math"
)

func minCostClimbingStairs(cost []int) int {
	// a = min cost to reach current house
	a := 0
	// b = min cost to reach previous house
	b := 0

	for i := 0; i < len(cost); i++ {
		// if we reach current house, we can either climb 1 or 2 steps
		// so we take min of cost to reach current house and previous house
		a, b = int(math.Min(float64(a+cost[i]), float64(b+cost[i]))), a
	}

	// return the min of cost to reach the last house
	return int(math.Min(float64(a), float64(b)))
}
