package leetcode

import (
	"strings"
)

func ValidPalindrome(s string) bool {
	var alphanum = "abcdefghijklmnopqrstuvwxyz0123456789"
	s = strings.ToLower(s)
	var formatted = make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(alphanum, rune(s[i])) {
			formatted = append(formatted, s[i])
		}
	}

	for i, j := 0, len(formatted)-1; i < j; {
		if formatted[i] != formatted[j] {
			return false
		}
		i++
		j--
	}

	return true
}
