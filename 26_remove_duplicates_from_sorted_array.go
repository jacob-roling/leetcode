package leetcode

func RemoveDuplicatesFromSortedArray(nums []int) int {
	i := 0
	j := 1
	k := 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[k] = nums[j]
			k++
		}
		i++
		j++
	}
	return k
}
