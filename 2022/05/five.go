package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Instruction struct {
	Num  int
	From int
	To   int
}

func parse(s string) ([][]byte, []Instruction) {
	start, instructions, found := strings.Cut(s, "\n\n")
	if !found {
		log.Fatal("parser split")
	}
	startLines := strings.Split(start, "\n")
	lastLine := len(startLines) - 1
	numStacks := len(startLines[lastLine])/4 + 1
	stacks := make([][]byte, numStacks, numStacks)
	for i := 0; i < lastLine; i++ {
		stackLine := startLines[lastLine-1-i]
		for j := 0; j < numStacks; j++ {
			box := stackLine[j*4+1]
			if box != ' ' {
				if len(stacks[j]) != i {
					log.Fatal("flying box")
				}
				stacks[j] = append(stacks[j], box)
			}
		}
	}

	lines := strings.Split(instructions, "\n")
	inst := []Instruction{}
	for _, line := range lines {
		var i Instruction

		n, err := fmt.Sscanf(line, "move %d from %d to %d", &i.Num, &i.From, &i.To)
		if n != 3 || err != nil {
			log.Fatalf("n = %d, err = %v", n, err)
		}
		i.From--
		i.To--
		inst = append(inst, i)
	}
	return stacks, inst
}

func process(stacks *[][]byte, inst Instruction) {
	for n := 0; n < inst.Num; n++ {
		st := (*stacks)[inst.From]
		el := st[len(st)-1]
		(*stacks)[inst.From] = st[:len(st)-1]
		(*stacks)[inst.To] = append((*stacks)[inst.To], el)
	}
}

func process9001(stacks *[][]byte, inst Instruction) {
	st := (*stacks)[inst.From]
	el := st[len(st)-inst.Num:]
	(*stacks)[inst.From] = st[:len(st)-inst.Num]
	(*stacks)[inst.To] = append((*stacks)[inst.To], el...)
}

func processAll(stacks *[][]byte, inst []Instruction) {
	for _, i := range inst {
		process(stacks, i)
	}
}

func processAll9001(stacks *[][]byte, inst []Instruction) {
	for _, i := range inst {
		process9001(stacks, i)
	}
}
func result(stacks [][]byte) string {
	ret := bytes.NewBufferString("")
	for i := 0; i < len(stacks); i++ {
		ret.WriteByte(stacks[i][len(stacks[i])-1])
	}
	return ret.String()
}

func part1() {

	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("read input.txt: %v", err)
	}
	top, bot := parse(string(buf))

	processAll(&top, bot)

	st := result(top)
	fmt.Println(st)
}

func part2() {

	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("read input.txt: %v", err)
	}
	top, bot := parse(string(buf))

	processAll9001(&top, bot)

	st := result(top)
	fmt.Println(st)
}

func main() {
	part1()
	part2()
}
