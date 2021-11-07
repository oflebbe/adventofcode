package main

import "testing"

func TestMain(t *testing.T) {
	i := -7
	if i%10 != -7 {
		t.Error("Main")
	}
	i = 11
	if i%10 != 1 {
		t.Error("Main")
	}
	i = -17
	if i%10 != -7 {
		t.Error("Main")
	}
}
