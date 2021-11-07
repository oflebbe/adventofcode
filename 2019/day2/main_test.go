package main

import (
	"fmt"
	"strings"
	"testing"
)

func runAndCompare(t *testing.T, start []int, stop []int) {
	runMemory(start)
	for k, v := range stop {
		if start[k] != v {
			t.Error(fmt.Sprintf("Result at %d value expected %d got %d", k, v, start[k]))
		}
	}
}

func TestMain(t *testing.T) {
	input := strings.Split("1,9,10,3,2,3,11,0,99,30,40,50", ",")
	mem := createMemory(input)
	expected := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	for k, v := range expected {
		if mem[k] != v {
			t.Error(fmt.Sprintf("Element at %d value expected %d got %d", k, v, mem[k]))
		}
	}
	runAndCompare(t, mem, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50})
	runAndCompare(t, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	runAndCompare(t, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	runAndCompare(t, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	runAndCompare(t, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
}
