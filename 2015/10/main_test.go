package main

import (
	"testing"
)

func TestFunc1(t *testing.T) {
	input := []int{1}
	output := []int{1, 1}
	if !testEq(transform(input), []int{1, 1}) {
		t.Fatalf("expected %v git %v", output, transform(input))
	}
}

func TestFunc11(t *testing.T) {
	input := []int{1, 1}
	output := []int{2, 1}
	if !testEq(transform(input), output) {
		t.Fatalf("expected %v git %v", output, transform(input))
	}
}

func TestFunc21(t *testing.T) {
	input := []int{2, 1}
	output := []int{1, 2, 1, 1}
	if !testEq(transform(input), output) {
		t.Fatalf("expected %v git %v", output, transform(input))
	}
}

func TestFunc1211(t *testing.T) {
	input := []int{1, 2, 1, 1}
	output := []int{1, 1, 1, 2, 2, 1}
	if !testEq(transform(input), output) {
		t.Fatalf("expected %v git %v", output, transform(input))
	}
}

func TestFunc111221(t *testing.T) {
	input := []int{1, 1, 1, 2, 2, 1}
	output := []int{3, 1, 2, 2, 1, 1}
	if !testEq(transform(input), output) {
		t.Fatalf("expected %v git %v", output, transform(input))
	}
}

func TestIter(t *testing.T) {
	digits := []int{1}
	output := []int{3, 1, 2, 2, 1, 1}
	for i := 0; i < 4; i++ {
		digits = transform(digits)
	}
	if !testEq(transform(digits), output) {
		t.Fatalf("expected %v", output)
	}

}
