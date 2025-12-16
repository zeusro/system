package problems

// ToOne computes the number of steps required to reduce a given integer i to 1
// 3n+1 角谷猜想
func ToOne(i int64) int64 {
	if i == 1 {
		return i
	}
	if i%2 == 0 {
		return ToOne(i / 2)
	}
	return ToOne(3*i + 1)
}
