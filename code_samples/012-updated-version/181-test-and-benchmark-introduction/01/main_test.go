package main

import "testing"

// expected and got idiom
func TestMySum(t *testing.T)  {
	got := mySum(2,3)
	if got != 5 {
		t.Error("Expected", 5, "Got", got)
	}
}
