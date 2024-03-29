package leetcode

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] - 9
		if carry < 1 {
			break
		}
		digits[i] = 0
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
