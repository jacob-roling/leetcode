package leetcode

func RotateArray(nums []int, k int) {
	length := len(nums)
	target_index := 0
	rotated_index := 0

	var last_target_index int

	a := nums[0]
	b := nums[0]

	for i := 0; i < length; i++ {
		// Compute the next rotated index
		rotated_index = (target_index + k) % length
		// Store the what's currently at the rotated index
		// then set the rotated index to the last stored value
		if i%2 == 0 {
			a = nums[rotated_index]
			nums[rotated_index] = b
		} else {
			b = nums[rotated_index]
			nums[rotated_index] = a
		}
		// Ensure we won't get caught in a loop
		if last_target_index == rotated_index {
			rotated_index = (target_index + k + 1) % length
			a = nums[rotated_index]
			b = nums[rotated_index]
		}
		// Update the target index
		last_target_index = target_index
		target_index = rotated_index
	}
}

// Alternate Approach
// func Rotate(nums []int, k int)  {
//     length := len(nums)
//     k %= length
//     count := 0
//     start := 0
//     for count < length {
//         next_idx := start
//         temp := nums[start]
//         for {
//             next_idx = (next_idx+k) % length
//             temp, nums[next_idx] = nums[next_idx], temp
//             count++
//             if start == next_idx {
//                 break
//             }
//         }
//         start++
//     }
// }
