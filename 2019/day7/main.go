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

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	max := 0

	for i0 := 0; i0 < 5; i0++ {
		for i1 := 0; i1 < 5; i1++ {
			if i1 == i0 {
				continue
			}
			for i2 := 0; i2 < 5; i2++ {
				if i2 == i0 || i2 == i1 {
					continue
				}

				for i3 := 0; i3 < 5; i3++ {
					if i3 == i2 || i3 == i1 || i3 == i0 {
						continue
					}
					for i4 := 0; i4 < 5; i4++ {
						if i4 == i3 || i4 == i0 || i4 == i2 || i4 == i1 {
							continue
						}
						phases := []int{i0 + 5, i1 + 5, i2 + 5, i3 + 5, i4 + 5}
						chs := make([]chan int, 5)
						for i := 0; i < 5; i++ {
							chs[i] = make(chan int, 1)
							chs[i] <- phases[i]
						}

						var wg sync.WaitGroup
						wg.Add(5)
						lastOut := make(chan int, 1)
						for i := 0; i < 5; i++ {
							mem := createMemory(lines)
							if i == 4 {
								go runMemory(mem, chs[i], chs[(i+1)%5], lastOut, &wg)
							} else {
								go runMemory(mem, chs[i], chs[(i+1)%5], nil, &wg)
							}
						}
						chs[0] <- 0
						last := 0

						for v := range lastOut {
							last = v
							fmt.Printf("Wert:%d: %d %d %d %d %d\n", v, i0+5, i1+5, i2+5, i3+5, i4+5)
						}
						wg.Wait()
						if last > max {
							max = last
						}

					}
				}
			}
		}
	}
	fmt.Println(max)

}

/*
func mainA() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")
	max := 0
	for i0 := 0; i0 < 5; i0++ {
		for i1 := 0; i1 < 5; i1++ {
			if i1 == i0 {
				continue
			}
			for i2 := 0; i2 < 5; i2++ {
				if i2 == i0 || i2 == i1 {
					continue
				}

				for i3 := 0; i3 < 5; i3++ {
					if i3 == i2 || i3 == i1 || i3 == i0 {
						continue
					}
					for i4 := 0; i4 < 5; i4++ {
						if i4 == i3 || i4 == i0 || i4 == i2 || i4 == i1 {
							continue
						}
						phases := []int{i0, i1, i2, i3, i4}
						output := 0
						for i := 0; i < 5; i++ {
							mem := createMemory(lines)
							in := []int{phases[i], output}
							out := runMemory(mem, in)
							output = out[0]
						}
						if output > max {
							max = output
						}
					}
				}
			}
		}
	}
	fmt.Println(max)

}*/
