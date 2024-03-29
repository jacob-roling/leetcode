package leetcode

import (
	"reflect"
	"testing"
)

func TestGroupAnagramsCase1(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := GroupAnagrams(strs)
	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestGroupAnagramsCase2(t *testing.T) {
	strs := []string{""}
	result := GroupAnagrams(strs)
	expected := [][]string{{""}}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestGroupAnagramsCase3(t *testing.T) {
	strs := []string{"a"}
	result := GroupAnagrams(strs)
	expected := [][]string{{"a"}}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
