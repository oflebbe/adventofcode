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

	Xdim, Ydim  int
	stride      int
	minCrossing int
}

type Wire struct {
	xp           []int
	yp           []int
	nearCrossing int
	numOfTests   int
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

func (w *Wire) testPoint(x, y int) {

	for i := 0; i < len(w.xp); i++ {
		if w.xp[i] == x && w.yp[i] == y {
			m := manhatten(x, y)
			if m < w.nearCrossing {
				w.nearCrossing = m
			}
		}
	}
}

func (w *Wire) testPointLen(x, y int) {
	w.numOfTests++
	for i := 0; i < len(w.xp); i++ {
		if w.xp[i] == x && w.yp[i] == y {
			m := i + w.numOfTests + 1
			if m < w.nearCrossing {
				w.nearCrossing = m
			}
		}
	}
}

func (w *Wire) dummy(x, y int) {

}

func (w *Wire) walk(str string, testPoint func(x, y int)) {
	toks := strings.Split(str, ",")
	xcur := 0
	ycur := 0
	w.xp = make([]int, 0)
	w.yp = make([]int, 0)
	for _, v := range toks {
		howmany, err := strconv.Atoi(v[1:])
		if err != nil {
			panic("parsing length")
		}
		switch v[0] {
		case 'R':
			for x := xcur + 1; x < xcur+howmany; x++ {
				w.xp = append(w.xp, x)
				w.yp = append(w.yp, ycur)
				testPoint(x, ycur)
			}
			xcur += howmany
		case 'L':
			for x := xcur - 1; x > xcur-howmany; x-- {
				w.xp = append(w.xp, x)
				w.yp = append(w.yp, ycur)
				testPoint(x, ycur)
			}
			xcur -= howmany
		case 'U':
			for y := ycur + 1; y < ycur+howmany; y++ {
				w.xp = append(w.xp, xcur)
				w.yp = append(w.yp, y)
				testPoint(xcur, y)
			}
			ycur += howmany
		case 'D':
			for y := ycur - 1; y > ycur-howmany; y-- {
				w.xp = append(w.xp, xcur)
				w.yp = append(w.yp, y)
				testPoint(xcur, y)
			}
			ycur -= howmany
		}
		w.xp = append(w.xp, xcur)
		w.yp = append(w.yp, ycur)
		testPoint(xcur, ycur)
	}
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	w := strings.Split(string(input), "\n")

	/*wire1 := Wire{}
	wire1.walk(w[0], func(x, y int) {})

	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(w[1], func(x, y int) {
		wire1.testPoint(x, y)
	})

	fmt.Println(wire1.nearCrossing) */

	wire3 := Wire{}
	wire3.walk(w[0], func(x, y int) {})

	wire4 := Wire{}
	wire3.nearCrossing = 100000000
	wire4.walk(w[1], func(x, y int) {
		wire3.testPointLen(x, y)
	})

	fmt.Println(wire3.nearCrossing)

}
