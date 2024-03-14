package leetcode

func RemoveDuplicatesFromSortedArray(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[k-1] != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
