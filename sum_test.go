package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The reult must be 5!")
	}
}

func TestSub(t *testing.T) {
	result := sub(5, 2)

	if result != 3 {
		t.Error("The reult must be 3!")
	}
}
