package leetcodeblind75

func climbStairs(n int) int {
	a, b := 1, 2 // 1. Initialize the first two steps

	for i := 0; i < n-1; i++ { // 2. Loop through n-1 steps
		a, b = b, a+b // 3. a is the current step, b is the next step
	}

	return a // 4. Return the last step
}
