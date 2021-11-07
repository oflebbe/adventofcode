package main

import (
	"fmt"
	"strings"
	"testing"
)

func runAndCompare(t *testing.T, start []int, stop []int) {
	in := []int{}

	_ = runMemory(start, in)
	for k, v := range stop {
		if start[k] != v {
			t.Error(fmt.Sprintf("Result at %d value expected %d got %d", k, v, start[k]))
		}
	}
}

func TestMain(t *testing.T) {
	input := strings.Split("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", ",")
	mem := createMemory(input)
	in := []int{0}
	out := runMemory(mem, in)
	if len(out) != 1 {
		t.Error("len not expect")
	}

	runAndCompare(t, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	runAndCompare(t, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	runAndCompare(t, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	runAndCompare(t, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
