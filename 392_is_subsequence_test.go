package leetcode

import "testing"

func TestIsSubsequenceCase1(t *testing.T) {
	if IsSubsequence("abc", "ahbgdc") != true {
		t.Fail()
	}
}

func TestIsSubsequenceCase2(t *testing.T) {
	if IsSubsequence("axc", "ahbgdc") != false {
		t.Fail()
	}
}
