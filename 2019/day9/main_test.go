package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {
	input := strings.Split("1102,34915192,34915192,7,4,7,99,0", ",")

	output := 0
	for i := 0; i < 5; i++ {
		mem := createMemory(input)
		in := make(chan int, 5)
		var wg sync.WaitGroup
		wg.Add(1)
		out := make(chan int, 5)

		runMemory(mem, in, out, nil, &wg)
		wg.Wait()
		if len(out) != 1 {
			t.Error("len not expect")
		}
		output = <-out
	}
	if len(strconv.Itoa(output)) != 16 {
		t.Error("rechner")
	}

}

func TestMain2(t *testing.T) {

	input := strings.Split("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", ",")

	for i := 0; i < 5; i++ {
		mem := createMemory(input)
		mem2 := createMemory(input)
		in := make(chan int, 5)
		var wg sync.WaitGroup
		wg.Add(1)
		out := make(chan int, len(mem2))
		runMemory(mem, in, out, nil, &wg)
		wg.Wait()
		if len(out) != len(mem2) {
			t.Error("len not expect")
		}
		for _, v := range mem2 {
			output := <-out
			if v != output {
				t.Error("compare")
			}
		}
	}

}

func TestDay51(t *testing.T) {
	input, err := ioutil.ReadFile("../day5/input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")
	in := make(chan int, 5)
	in <- 1
	close(in)
	var wg sync.WaitGroup
	wg.Add(1)
	out := make(chan int, 50)
	mem := createMemory(lines)

	runMemory(mem, in, out, nil, &wg)
	wg.Wait()
	last := false
	for v := range out {
		output := v
		if output != 0 {
			last = true
			if output != 11049715 {
				t.Error("compare")
			}
		}

	}
	if last == false {
		t.Error("last")
	}

}
