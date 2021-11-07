package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
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

type memory struct {
	ram      []int
	extended map[int]int
	baseReg  int
}

func newMem(ram []int) memory {
	return memory{ram, make(map[int]int), 0}
}

func (m *memory) mem(addr int, mode int) int {
	var ret int
	addr = m.rawMem(addr)
	if mode == 1 {
		ret = addr
	} else if mode == 0 || mode == 2 {
		if mode == 2 {
			addr += m.baseReg
		}
		ret = m.rawMem(addr)
	} else {
		panic("wrong mode")
	}
	return ret
}

func (m *memory) rawMem(addr int) int {
	var ret int
	if addr < len(m.ram) {
		ret = m.ram[addr]
	} else {
		ret = m.extended[addr]
	}
	return ret
}

func (m *memory) setRawMem(addr, value int) int {
	var ret int

	if addr < len(m.ram) {
		m.ram[addr] = value
	} else {
		m.extended[addr] = value
	}
	return ret
}

func (m *memory) setMem(addr int, mode int, value int) {
	addr = m.rawMem(addr)
	if mode == 1 {
		panic("cannot set direct")
	} else if mode == 0 || mode == 2 {
		if mode == 2 {
			addr += m.baseReg
		}

		m.setRawMem(addr, value)

	}
}

func (m *memory) adjustBase(base int) {
	m.baseReg += base
}

func runMemory(initialMem []int, inCh chan int, outCh chan int, otherOut chan int, wg *sync.WaitGroup) {

	mem := newMem(initialMem)

	for pc := 0; mem.rawMem(pc) != 99; {
		addr1 := mem.rawMem(pc) / 100 % 10
		addr2 := mem.rawMem(pc) / 1000 % 10
		addr3 := mem.rawMem(pc) / 10000 % 10
		switch mem.rawMem(pc) % 100 {
		case 1:
			a := mem.mem(pc+1, addr1)
			b := mem.mem(pc+2, addr2)

			mem.setMem(pc+3, addr3, a+b)
			pc += 4
		case 2:
			a := mem.mem(pc+1, addr1)
			b := mem.mem(pc+2, addr2)

			mem.setMem(pc+3, addr3, a*b)
			pc += 4
		case 3:
			a := <-inCh
			mem.setMem(pc+1, addr1, a)

			pc += 2
		case 4:
			a := mem.mem(pc+1, addr1)
			outCh <- a
			if otherOut != nil {
				otherOut <- a
			}

			pc += 2
		case 5:
			a := mem.mem(pc+1, addr1)
			if a != 0 {
				b := mem.mem(pc+2, addr2)
				pc = b
			} else {
				pc += 3
			}
		case 6:
			a := mem.mem(pc+1, addr1)
			if a == 0 {
				b := mem.mem(pc+2, addr2)
				pc = b
			} else {
				pc += 3
			}
		case 7:
			a := mem.mem(pc+1, addr1)
			b := mem.mem(pc+2, addr2)

			var r int
			if a < b {
				r = 1
			} else {
				r = 0
			}
			mem.setMem(pc+3, addr3, r)
			pc += 4
		case 8:
			a := mem.mem(pc+1, addr1)
			b := mem.mem(pc+2, addr2)

			var r int
			if a == b {
				r = 1
			} else {
				r = 0
			}
			mem.setMem(pc+3, addr3, r)
			pc += 4
		case 9:
			a := mem.mem(pc+1, addr1)
			mem.adjustBase(a)
			pc += 2

		default:
			panic(fmt.Sprintf("unknown opconde %d at %d\n", mem.mem(pc, 0), pc))
		}
	}
	close(outCh)

	if otherOut != nil {
		close(otherOut)
	}
	wg.Done()
	return
}

type point struct {
	x, y int
}

func robot(input chan int, output chan int, result chan int, hull *map[point]int) {
	currentHeading := point{0, -1}
	currentPositioning := point{0, 0}
	count := 0
	for i := range input {
		if count%2 == 0 {
			(*hull)[currentPositioning] = i
		} else {
			if i == 0 {
				switch currentHeading {
				case point{0, -1}:
					currentHeading = point{-1, 0}
				case point{-1, 0}:
					currentHeading = point{0, 1}
				case point{0, 1}:
					currentHeading = point{1, 0}
				case point{1, 0}:
					currentHeading = point{0, -1}
				}
			} else {
				switch currentHeading {
				case point{0, -1}:
					currentHeading = point{1, 0}
				case point{1, 0}:
					currentHeading = point{0, 1}
				case point{0, 1}:
					currentHeading = point{-1, 0}
				case point{-1, 0}:
					currentHeading = point{0, -1}
				}
			}
			currentPositioning.x += currentHeading.x
			currentPositioning.y += currentHeading.y
			color, ok := (*hull)[currentPositioning]
			if !ok {
				output <- 0
			} else {
				output <- color
			}
		}
		count++
	}
	result <- len(*hull)
	close(result)
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	inp := make(chan int, 1)
	out := make(chan int, 1)
	result := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	mem := createMemory(lines)
	go runMemory(mem, inp, out, nil, &wg)
	hull := make(map[point]int)

	go robot(out, inp, result, &hull)
	inp <- 1
	<-result

	ymin := 999999999999999
	ymax := -999999999999999
	xmin := 999999999999999
	xmax := -999999999999999
	for k, v := range hull {
		if v == 0 {
			continue
		}
		if k.x < xmin {
			xmin = k.x
		}
		if k.x > xmax {
			xmax = k.x
		}
		if k.y < ymin {
			ymin = k.y
		}
		if k.y > ymax {
			ymax = k.y
		}
	}
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			c, ok := hull[point{x, y}]
			if ok && c == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
