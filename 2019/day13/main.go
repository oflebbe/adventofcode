package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
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
	mem.setRawMem(0, 2)
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

func screen(input chan int, paddle *[2]int, ball *[2]int, wg *sync.WaitGroup) {
	count := 0
	x := 0
	y := 0
	bt := 0
	fmt.Print(cursor.ClearEntireScreen())
	fmt.Print(cursor.Hide())
	for i := range input {
		if count%3 == 0 {
			x = i
		} else if count%3 == 1 {
			y = i
		} else {

			if x == -1 {
				fmt.Print(cursor.MoveTo(y+1, x+1))
				fmt.Print(i)
			} else {
				fmt.Print(cursor.MoveTo(y+2, x+1))
				var ch string
				switch i {
				case 0:
					ch = " "
				case 1:
					ch = "|"
				case 2:
					ch = "X"
				case 3:
					(*paddle)[0] = x
					(*paddle)[1] = y
					ch = "="
				case 4:
					(*ball)[0] = x
					(*ball)[1] = y
					ch = "O"
					//					out <- [2]int{x,y}
				}
				fmt.Print(ch)
			}
		}
		count++
	}
	fmt.Print(cursor.MoveTo(30, 2))
	fmt.Printf("Blocktiles: %d\n", bt)
	wg.Done()
}

/*
func paddleControl(out chan int) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic("keyboard")
		} else if key == keyboard.KeyArrowLeft {
			out <- -1
		} else if key == keyboard.KeyArrowRight {
			out <- 1
		} else {
			out <- 0
		}
		//time.Sleep(1000000000)
	}
} */

func paddleControl(out chan int, paddle, ball *[2]int) {
	out <- 0
	for {
		time.Sleep(100000000000)
		if paddle[0] > ball[0] {
			out <- -1
		} else if paddle[0] < ball[0] {
			out <- 1
		} else {
			out <- 0
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	inp := make(chan int, 1)
	out := make(chan int)

	paddle := [2]int{}
	ball := [2]int{}
	var wg sync.WaitGroup
	wg.Add(2)
	mem := createMemory(lines)
	go paddleControl(inp, &paddle, &ball)
	go runMemory(mem, inp, out, nil, &wg)

	go screen(out, &paddle, &ball, &wg)
	wg.Wait()
}
