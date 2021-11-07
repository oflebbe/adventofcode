package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := strings.Split("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", ",")
	phases := []int{4, 3, 2, 1, 0}
	output := 0
	for i := 0; i < 5; i++ {
		mem := createMemory(input)
		in := []int{phases[i], output}
		out := runMemory(mem, in)
		if len(out) != 1 {
			t.Error("len not expect")
		}
		output = out[0]
	}
	if output != 43210 {
		t.Error("rechner")
	}

}

func TestMain2(t *testing.T) {
	input := strings.Split("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", ",")
	phases := []int{0, 1, 2, 3, 4}
	output := 0
	for i := 0; i < 5; i++ {
		mem := createMemory(input)
		in := []int{phases[i], output}
		out := runMemory(mem, in)
		if len(out) != 1 {
			t.Error("len not expect")
		}
		output = out[0]
	}
	if output != 54321 {
		t.Error("rechner")
	}

}

func TestMain3(t *testing.T) {
	input := strings.Split("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", ",")
	phases := []int{1, 0, 4, 3, 2}
	output := 0
	for i := 0; i < 5; i++ {
		mem := createMemory(input)
		in := []int{phases[i], output}
		out := runMemory(mem, in)
		if len(out) != 1 {
			t.Error("len not expect")
		}
		output = out[0]
	}
	if output != 65210 {
		t.Error("rechner")
	}

}
