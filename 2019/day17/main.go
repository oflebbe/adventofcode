package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

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

type Coord struct {
	x, y int
}

type Map struct {
	what map[Coord]int
}

func follow(coord Coord, arrow int) Coord {
	switch arrow {
	case 1:
		coord.y++
	case 2:
		coord.y--
	case 3:
		coord.x++
	case 4:
		coord.x--
	}
	return coord
}

type Display struct {
	x, y int
	what int
}

func newDisplay(coord Coord, what int) Display {
	n := Display{coord.x, coord.y, what}
	return n
}

func robot(in chan int, out chan int, dis chan Display, coord Coord, karte *map[Coord]int, wg *sync.WaitGroup) {
	var iter func(coor Coord)

	iter = func(coord Coord) {
		for arrow := 1; arrow <= 4; arrow++ {
			new := follow(coord, arrow)
			_, ok := (*karte)[new]
			if ok {
				continue
			}
			out <- arrow
			newWhat := <-in
			(*karte)[new] = newWhat
			dis <- newDisplay(new, newWhat)
			if newWhat != 0 {
				iter(new)
				switch arrow {
				case 1:
					out <- 2
					<-in
				case 2:
					out <- 1
					<-in
				case 3:
					out <- 4
					<-in
				case 4:
					out <- 3
					<-in
				}
			}
		}
	}
	iter(coord)
	close(dis)
	wg.Done()
}

func display(dis chan int) map[Coord]int {

	fmt.Print(cursor.ClearEntireScreen())
	fmt.Print(cursor.Hide())
	fmt.Print(cursor.MoveTo(1, 1))

	maps := make(map[Coord]int)
	x := 0
	y := 0
	for d := range dis {
		fmt.Printf("%c", d)
		if d == 10 {
			x = 0
			y++
		} else {
			if d == 35 {
				c := Coord{x, y}
				maps[c] = d
			}
			x++
		}
	}

	return maps
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	inp := make(chan int, 1)
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	mem := createMemory(lines)
	go runMemory(mem, inp, out, nil, &wg)
	coord := Coord{0, 0}
	karte := make(map[Coord]int)
	karte[coord] = 1
	//	go robot(out, inp, dis, coord, &karte, &wg)
	maps := display(out)
	sum := 0
	howmany := 0
	fmt.Print(cursor.MoveTo(2, 2))
	fmt.Print("Z")

	for k := range maps {
		k2 := Coord{k.x - 1, k.y}
		_, ok := maps[k2]
		if !ok {
			continue
		}
		k2 = Coord{k.x, k.y - 1}
		_, ok = maps[k2]
		if !ok {
			continue
		}
		k2 = Coord{k.x, k.y + 1}
		_, ok = maps[k2]
		if !ok {
			continue
		}
		k2 = Coord{k.x + 1, k.y}
		_, ok = maps[k2]
		if !ok {
			continue
		}
		sum += (k.x) * (k.y)
		fmt.Print(cursor.MoveTo(k.y+1, k.x+1))
		fmt.Print("O")

		howmany++
	}
	fmt.Print(cursor.MoveTo(40, 0))
	fmt.Println(sum)
	fmt.Println(howmany)

	mem := createMemory(lines)
	mem[0] = 2
	go runMemory(mem, inp, out, nil, &wg)

}
