package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrayCase1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}
	MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Fail()
	}
}

func TestMergeSortedArrayCase2(t *testing.T) {
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0
	expected := []int{1}
	MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	expected := []int{1}
	MergeSortedArray(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Fail()
	}
}
