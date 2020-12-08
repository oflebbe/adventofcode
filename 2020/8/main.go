package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name     string
	argument int
}

const loop = 1
const term = 2

func problem(program []instruction) (int, int) {
	acc := 0
	pc := 0
	alreadySeen := make(map[int]bool)
	for {
		_, ok := alreadySeen[pc]
		if ok {
			return acc, loop
		}
		alreadySeen[pc] = true
		if pc >= len(program) {
			return acc, term
		}
		switch program[pc].name {
		case "acc":
			acc += program[pc].argument
		case "jmp":
			pc += program[pc].argument
			continue
		case "nop":
			break
		}
		pc++
	}
}

func fix(program []instruction) int {
	for i := 0; i < len(program); i++ {
		switch program[i].name {
		case "nop":
			program[i].name = "jmp"
		case "jmp":
			program[i].name = "nop"
		default:
			continue
		}
		acc, state := problem(program)
		if state == term {
			return acc
		}
		switch program[i].name {
		case "nop":
			program[i].name = "jmp"
		case "jmp":
			program[i].name = "nop"
		default:
			continue
		}
	}
	panic("no solution")
}

func parse(input string) []instruction {
	var prog []instruction
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		tok := strings.Fields(line)
		val, _ := strconv.Atoi(tok[1])
		prog = append(prog, instruction{tok[0], val})
	}
	return prog
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	prog := parse(string(buf))
	acc, _ := problem(prog)
	println(acc)
	acc = fix(prog)
	println(acc)
}
