package leetcode

import (
	"testing"
)

func TestLongestSubstringWithoutReapeatingCharactersCase1(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := LongestSubstringWithoutReapeatingCharacters(s)

	if result != expected {
		t.Fail()
	}
}

func TestLongestSubstringWithoutReapeatingCharactersCase2(t *testing.T) {
	s := "bbbbb"
	expected := 1
	result := LongestSubstringWithoutReapeatingCharacters(s)

	if result != expected {
		t.Fail()
	}
}

func TestLongestSubstringWithoutReapeatingCharactersCase3(t *testing.T) {
	s := "pwwkew"
	expected := 3
	result := LongestSubstringWithoutReapeatingCharacters(s)

	if result != expected {
		t.Fail()
	}
}

func TestLongestSubstringWithoutReapeatingCharactersCase4(t *testing.T) {
	s := "dvdf"
	expected := 3
	result := LongestSubstringWithoutReapeatingCharacters(s)

	if result != expected {
		t.Fail()
	}
}
