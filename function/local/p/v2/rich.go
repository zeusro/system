package v2

// 求最大公约数 GCD
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 求最小公倍数 LCM
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 对一组数求 LCM
func LcmOfList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for _, n := range nums[1:] {
		result = lcm(result, n)
	}
	return result
}
