package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElementCase1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2
	result := RemoveElement(nums, val)
	if !reflect.DeepEqual(nums[:expected], []int{2, 2}) || result != expected {
		t.Fail()
	}
}

func TestRemoveElementCase2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expected := 5
	result := RemoveElement(nums, val)
	if !reflect.DeepEqual(nums[:expected], []int{0, 1, 3, 0, 4}) || result != expected {
		t.Fail()
	}
}
