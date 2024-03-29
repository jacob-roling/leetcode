package leetcode

func LongestSubstringWithoutReapeatingCharacters(s string) int {
	length := len(s)
	if length < 1 {
		return 0
	}
	if length < 2 {
		return 1
	}
	maxWindowLength := 0
	for i, j := 0, 1; j < length; {
		for k := i; k < j; k++ {
			if s[k] == s[j] {
				i = k + 1
				break
			}
		}
		j++
		windowLength := j - i
		if windowLength > maxWindowLength {
			maxWindowLength = windowLength
		}
	}
	return maxWindowLength
}
