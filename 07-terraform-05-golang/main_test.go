package main

import "testing"

func TestLittle_Element(t *testing.T) {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	res := little_element(x)

	if res != 9 {
		t.Error("Expected 9, got ", res)
	}
}

func TestDiv_three(t *testing.T) {
	numbers := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}
	var res []int
	res = div_three()

	if res[5] != numbers[5] {
		t.Error("Expected other, got ", res)
	}
}
