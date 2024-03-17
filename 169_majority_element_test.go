package leetcode

import (
	"testing"
)

func TestMajorityElementCase1(t *testing.T) {
	nums := []int{3, 2, 3}
	expected := 3
	result := MajorityElement(nums)
	if expected != result {
		t.Fail()
	}
}

func TestMajorityElementCase2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2
	result := MajorityElement(nums)
	if expected != result {
		t.Fail()
	}
}

func TestMajorityElementCase3(t *testing.T) {
	nums := []int{6, 5, 5}
	expected := 5
	result := MajorityElement(nums)
	if expected != result {
		t.Fail()
	}
}
