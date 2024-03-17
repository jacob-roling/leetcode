package leetcode

func RotateArray(nums []int, k int) {
	length := len(nums)
	k %= length
	for i, j := 0, 0; i < length; j++ {
		next_idx := j
		temp := nums[j]
		for {
			next_idx = (next_idx + k) % length
			temp, nums[next_idx] = nums[next_idx], temp
			i++
			if j == next_idx {
				break
			}
		}
	}
}
