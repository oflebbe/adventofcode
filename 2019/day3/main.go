package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Grid is the layout
type Grid struct {
	Xmin, Xmax, Ymin, Ymax int
	Grid                   []rune
	Xdim, Ydim             int
	stride                 int
	minCrossing            int
}

func (g *Grid) updateExtension(str string) {
	toks := strings.Split(str, ",")
	xcur := 0
	ycur := 0
	for _, v := range toks {
		howmany, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("parsing length")
		}
		switch v[0] {
		case 'R':
			xcur += howmany
		case 'L':
			xcur -= howmany
		case 'U':
			ycur += howmany
		case 'D':
			ycur -= howmany
		}
		if xcur > g.Xmax {
			g.Xmax = xcur
		}
		if xcur < g.Xmin {
			g.Xmin = xcur
		}
		if ycur > g.Ymax {
			g.Ymax = ycur
		}
		if ycur < g.Ymin {
			g.Ymin = ycur
		}
	}
}

func (g *Grid) createGrid() {
	g.Xdim = g.Xmax - g.Xmin + 3
	g.Ydim = g.Ymax - g.Ymin + 3
	g.Grid = make([]rune, g.Xdim*g.Ydim)
	for i := 0; i < g.Xdim*g.Ydim; i++ {
		g.Grid[i] = '.'
	}
	g.Grid[-g.Xmin+1+(-g.Ymin+1)*g.Xdim] = 'o'
	g.minCrossing = g.Xdim + g.Ydim
}

func (g *Grid) print() string {
	st := ""
	for y := g.Ydim - 1; y >= 0; y-- {
		for x := 0; x < g.Xdim; x++ {
			st += string(g.Grid[x+y*g.Xdim])
		}
		st += string("\n")
	}
	return st
}

func manhatten(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func (g *Grid) setPoint(xcur, ycur int, ch rune) {
	orig := g.getPoint(xcur, ycur)
	switch orig {
	case 'o':
		ch = '?'
	case '-':
		if ch == '-' {
			ch = '='
		} else if ch == '|' {
			ch = 'X'
			//fmt.Printf("%d %d %d\n", xcur, ycur, manhatten(xcur, ycur))
			if manhatten(xcur, ycur) < g.minCrossing {
				g.minCrossing = manhatten(xcur, ycur)

			}
		} else {
			ch = '?'
		}
	case '|':
		if ch == '|' {
			ch = 'W'
		} else if ch == '-' {
			ch = 'X'
			//fmt.Printf("%d %d %d\n", xcur, ycur, manhatten(xcur, ycur))
			if manhatten(xcur, ycur) < g.minCrossing {
				g.minCrossing = manhatten(xcur, ycur)

			}
		} else {
			ch = '?'
		}
	case '+':
		ch = '?'
	}

	g.Grid[xcur-g.Xmin+1+(ycur-g.Ymin+1)*g.Xdim] = ch
}

func (g *Grid) getPoint(xcur, ycur int) rune {
	return g.Grid[xcur-g.Xmin+1+(ycur-g.Ymin+1)*g.Xdim]
}

func (g *Grid) drawWire(str string) {

	toks := strings.Split(str, ",")
	xcur := 0
	ycur := 0
	for k, v := range toks {
		howmany, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("parsing length")
		}
		lastChar := ' '
		switch v[0] {
		case 'R':
			for x := xcur + 1; x < xcur+howmany; x++ {
				lastChar = '-'
				g.setPoint(x, ycur, '-')
			}
			xcur += howmany
		case 'L':
			for x := xcur - 1; x > xcur-howmany; x-- {
				lastChar = '-'
				g.setPoint(x, ycur, '-')
			}
			xcur -= howmany
		case 'U':
			for y := ycur + 1; y < ycur+howmany; y++ {
				lastChar = '|'
				g.setPoint(xcur, y, '|')
			}
			ycur += howmany
		case 'D':
			for y := ycur - 1; y > ycur-howmany; y-- {
				lastChar = '|'
				g.setPoint(xcur, y, '|')
			}
			ycur -= howmany
		}
		if k != len(toks)-1 {
			g.setPoint(xcur, ycur, rune(strconv.Itoa(k)[0]))
		} else {
			g.setPoint(xcur, ycur, lastChar)
		}
	}
}

func support(w1, w2 string) Grid {
	grid := Grid{}
	grid.updateExtension(w2)
	grid.updateExtension(w1)
	grid.createGrid()
	grid.drawWire(w2)
	fmt.Println(grid.minCrossing)
	grid.minCrossing = 20000
	grid.drawWire(w1)
	return grid
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	w := strings.Split(string(input), "\n")

	grid := support(w[0], w[1])
	//	_ = ioutil.WriteFile("output", []byte(grid.print()), 0644)

	fmt.Println(grid.minCrossing)

}
