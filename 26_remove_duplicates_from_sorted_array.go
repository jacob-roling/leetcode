package leetcode

func RemoveDuplicatesFromSortedArray(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[count-1] != nums[i] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
