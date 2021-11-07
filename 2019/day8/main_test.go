package main

import "testing"

func TestMain(t *testing.T) {
	line := "123456789012"
	pic := readInput(3, 2, line)
	if len(pic.layers) != 2 {
		t.Error("Main")
	}
	if score(pic) != 1 {
		t.Error("score")
	}
}
