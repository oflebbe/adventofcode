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

func display(dis chan Display, wg *sync.WaitGroup) {

	fmt.Print(cursor.ClearEntireScreen())
	fmt.Print(cursor.Hide())
	fmt.Print(cursor.MoveTo(0+20, 0+30))
	fmt.Print("*")

	for d := range dis {
		fmt.Print(cursor.MoveTo(d.y+20, d.x+30))
		var x string
		switch d.what {
		case 0:
			x = "X"
		case 1:
			x = "."
		case 2:
			x = "!"
		}
		fmt.Print(x)
	}

	fmt.Print(cursor.MoveTo(40+40, 0))

	wg.Done()
}

/*
func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	inp := make(chan int, 1)
	out := make(chan int)
	dis := make(chan Display)
	var wg sync.WaitGroup
	wg.Add(2)
	mem := createMemory(lines)
	go runMemory(mem, inp, out, nil, &wg)
	coord := Coord{0, 0}
	karte := make(map[Coord]int)
	karte[coord] = 1
	go robot(out, inp, dis, coord, &karte, &wg)
	go display(dis, &wg)
	wg.Wait()
}*/

func createGraph(karte *map[Coord]int) map[Coord][]Coord {
	ret := make(map[Coord][]Coord)

	for c, v := range *karte {
		if v == 0 {
			continue
		}
		list := make([]Coord, 0)
		for i := 1; i <= 4; i++ {
			o := follow(c, i)
			if (*karte)[o] == 0 {
				continue
			}
			list = append(list, o)
		}
		ret[c] = list
	}
	return ret
}

func solveGraph(liste *map[Coord][]Coord, targetCoord Coord) int {
	alreadyVisited := make(map[Coord]struct{})

	var iter func(Coord) int
	min := len(*liste)
	count := 0

	iter = func(coord Coord) int {
		if coord == targetCoord {
			return count
		}
		if count > min {
			return min
		}
		count++
		alreadyVisited[coord] = struct{}{}
		for _, v := range (*liste)[coord] {
			_, ok := alreadyVisited[v]
			if ok {
				continue
			}
			min = iter(v)
		}
		count--
		return min
	}
	min = iter(Coord{0, 0})
	return min
}

func fillWithOxyGen(liste *map[Coord][]Coord, targetCoord Coord) int {
	alreadyFilled := make(map[Coord]int)

	count := 0
	alreadyFilled[targetCoord] = 0
	howmany := len(*liste)
	lasthowmany := 0
	for howmany > 0 && lasthowmany != howmany {
		lasthowmany = howmany

		for c, v := range alreadyFilled {
			if v == count {
				for _, next := range (*liste)[c] {
					_, ok := alreadyFilled[next]
					if ok {
						continue
					}
					alreadyFilled[next] = count + 1
					howmany--
				}
			}

		}
		count++
	}
	if howmany == lasthowmany {
		panic("no solution")
	}
	return count
}

func main() {
	input, err := ioutil.ReadFile("input")

	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), ",")

	inp := make(chan int, 1)
	out := make(chan int)
	dis := make(chan Display)
	var wg sync.WaitGroup
	wg.Add(2)
	mem := createMemory(lines)
	go runMemory(mem, inp, out, nil, &wg)
	coord := Coord{0, 0}
	karte := make(map[Coord]int)
	karte[coord] = 1
	go robot(out, inp, dis, coord, &karte, &wg)
	go display(dis, &wg)
	wg.Wait()

	var targetCoord Coord
	for k, v := range karte {
		if v == 2 {
			targetCoord = k
			break
		}
	}
	list := createGraph(&karte)
	min := solveGraph(&list, targetCoord)
	fmt.Println(min)
	fill := fillWithOxyGen(&list, targetCoord)
	fmt.Println(fill)

}
