package leetcode

// Attempt 1 Going Forwards

// func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
// 	i := 0
// 	j := 0

// 	if n > 0 {
// 		for i < m {
// 			for nums2[j] <= nums1[i] {
// 				nums1 = append(nums1[:i+1], nums1[i:len(nums1)-1]...)
// 				nums1[i] = nums2[j]
// 				j++
// 			}
// 			i++
// 		}
// 	}

// 	i = m + j

// 	for j < n {
// 		nums1[i] = nums2[j]
// 		i++
// 		j++
// 	}
// }

// Attempt 2 Going Backwards

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if m > 0 && (n == 0 || nums1[m-1] >= nums2[n-1]) {
			nums1[m+n-1] = nums1[m-1]
			m--
			continue
		}
		nums1[m+n-1] = nums2[n-1]
		n--
	}
}
