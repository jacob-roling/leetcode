package leetcode

import (
	"testing"
)

func TestBestTimeToBuyAndSellStockCase1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	result := BestTimeToBuyAndSellStock(prices)
	if expected != result {
		t.Fail()
	}
}

func TestBestTimeToBuyAndSellStockCase2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0
	result := BestTimeToBuyAndSellStock(prices)
	if expected != result {
		t.Fail()
	}
}
