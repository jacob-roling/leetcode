package leetcode

import (
	"reflect"
	"testing"
)

func TestRotateArrayCase1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	RotateArray(nums, k)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(nums, expected) {
		t.Fail()
	}
}

func TestRotateArrayCase2(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	RotateArray(nums, k)
	expected := []int{3, 99, -1, -100}
	if !reflect.DeepEqual(nums, expected) {
		t.Fail()
	}
}
