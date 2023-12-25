package math

// GCD calculates the Greatest Common Divisor of two integers
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM calculates the Least Common Multiple of multiple integers
func LCM(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	lcm := nums[0]
	for _, num := range nums[1:] {
		lcm = (lcm * num) / GCD(lcm, num)
	}
	return lcm
}
