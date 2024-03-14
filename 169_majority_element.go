package leetcode

func MajorityElement(nums []int) int {
	max := 0
	element := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if i+1-max > max {
				max = i + 1 - max
				element = nums[i-1]
			}
		}
	}
	return element
}
