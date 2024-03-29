package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOnceCase1(t *testing.T) {
	digits := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	result := PlusOne(digits)
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestPlusOnceCase2(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}
	result := PlusOne(digits)
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestPlusOnceCase3(t *testing.T) {
	digits := []int{9}
	expected := []int{1, 0}
	result := PlusOne(digits)
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
