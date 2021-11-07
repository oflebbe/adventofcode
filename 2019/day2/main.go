package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func createMemory(program []string) []int {
	mem := make([]int, 0)
	for _, v := range program {
		t, err := strconv.Atoi(v)
		if err != nil {
			panic("token" + v)
		}
		mem = append(mem, t)
	}
	return mem
}

func runMemory(mem []int) {
	for pc := 0; mem[pc] != 99; pc += 4 {
		switch mem[pc] {
		case 1:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			mem[c] = mem[a] + mem[b]
		case 2:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			mem[c] = mem[a] * mem[b]
		default:
			panic("Unknown opcode")
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	inputlines := strings.Split(string(input), ",")
	mem := createMemory(inputlines)
	stash := make([]int, len(mem))
	copy(stash, mem)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(mem, stash)
			mem[1] = noun
			mem[2] = verb
			runMemory(mem)
			if mem[0] == 19690720 {
				fmt.Printf("noun %d, verb %d = %d\n", noun, verb, 100*noun+verb)
			}
		}
	}

}
