package leetcode

import "testing"

func TestValidPalindromeCase1(t *testing.T) {
	var s = "A man, a plan, a canal: Panama"
	var expected = true
	var result = ValidPalindrome(s)
	if expected != result {
		t.Fail()
	}
}

func TestValidPalindromeCase2(t *testing.T) {
	var s = "race a car"
	var expected = false
	var result = ValidPalindrome(s)
	if expected != result {
		t.Fail()
	}
}

func TestValidPalindromeCase3(t *testing.T) {
	var s = " "
	var expected = true
	var result = ValidPalindrome(s)
	if expected != result {
		t.Fail()
	}
}
