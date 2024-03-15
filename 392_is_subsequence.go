package leetcode

func IsSubsequence(s string, t string) bool {
	for i, j := 0, 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
			if j == len(s) {
				return true
			}
		}
	}
	return false
}
