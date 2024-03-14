package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesFromSortedArrayCase1(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	result := RemoveDuplicatesFromSortedArray(nums)
	if !reflect.DeepEqual(nums[:expected], []int{1, 2}) || result != expected {
		t.Fail()
	}
}

func TestRemoveDuplicatesFromSortedArrayCase2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	result := RemoveDuplicatesFromSortedArray(nums)
	if !reflect.DeepEqual(nums[:expected], []int{0, 1, 2, 3, 4}) || result != expected {
		t.Fail()
	}
}
