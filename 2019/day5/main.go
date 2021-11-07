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

func runMemory(mem []int, in []int) []int {
	out := make([]int, 0)
	for pc := 0; mem[pc] != 99; {
		print(pc)
		addr1 := mem[pc] / 100 % 10
		addr2 := mem[pc] / 1000 % 10
		switch mem[pc] % 100 {
		case 1:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			if addr1 == 0 {
				a = mem[a]
			} else if addr1 != 1 {
				panic("Unknwon mode")
			}
			if addr2 == 0 {
				b = mem[b]
			} else if addr2 != 1 {
				panic("Unknwon mode")
			}

			mem[c] = a + b
			pc += 4
		case 2:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			if addr1 == 0 {
				a = mem[a]
			} else if addr1 != 1 {
				panic("Unknwon mode")
			}
			if addr2 == 0 {
				b = mem[b]
			} else if addr2 != 1 {
				panic("Unknwon mode")
			}

			mem[c] = a * b
			pc += 4
		case 3:
			var a int
			a, in = in[0], in[1:]
			b := mem[pc+1]
			mem[b] = a
			if addr1 != 0 || addr2 != 0 {
				panic("Unknown mode")
			}
			pc += 2
		case 4:
			a := mem[pc+1]
			if addr1 == 0 {
				a = mem[a]
			}
			out = append(out, a)
			if addr2 != 0 {
				panic("Unknown mode")
			}
			pc += 2
		case 5:
			a := mem[pc+1]
			if addr1 == 0 {
				a = mem[a]
			}
			if a != 0 {
				b := mem[pc+2]
				if addr2 == 0 {
					b = mem[b]
				}
				pc = b
			} else {
				pc += 3
			}
		case 6:
			a := mem[pc+1]
			if addr1 == 0 {
				a = mem[a]

			}
			if a == 0 {
				b := mem[pc+2]
				if addr2 == 0 {
					b = mem[b]
				}
				pc = b
			} else {
				pc += 3
			}
		case 7:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			if addr1 == 0 {
				a = mem[a]
			} else if addr1 != 1 {
				panic("Unknwon mode")
			}
			if addr2 == 0 {
				b = mem[b]
			} else if addr2 != 1 {
				panic("Unknwon mode")
			}

			var r int
			if a < b {
				r = 1
			} else {
				r = 0
			}
			mem[c] = r
			pc += 4
		case 8:
			a := mem[pc+1]
			b := mem[pc+2]
			c := mem[pc+3]
			if addr1 == 0 {
				a = mem[a]
			} else if addr1 != 1 {
				panic("Unknwon mode")
			}
			if addr2 == 0 {
				b = mem[b]
			} else if addr2 != 1 {
				panic("Unknwon mode")
			}

			var r int
			if a == b {
				r = 1
			} else {
				r = 0
			}
			mem[c] = r
			pc += 4

		default:
			panic(fmt.Sprintf("unknown opconde %d at %d\n", mem[pc], pc))
		}
	}
	return out
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	inputlines := strings.Split(string(input), ",")
	mem := createMemory(inputlines)

	in := []int{5}
	out := runMemory(mem, in)
	fmt.Println(out)

}
