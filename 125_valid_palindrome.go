package leetcode

import (
	"strings"
	"unicode"
)

func ValidPalindrome(s string) bool {
	filter := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	s = strings.Map(filter, s)

	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// First method of formatting
// var alphanum = "abcdefghijklmnopqrstuvwxyz0123456789"
// s = strings.ToLower(s)
// var formatted = make([]byte, 0)

// for i := 0; i < len(s); i++ {
// 	if strings.ContainsRune(alphanum, rune(s[i])) {
// 		formatted = append(formatted, s[i])
// 	}
// }
