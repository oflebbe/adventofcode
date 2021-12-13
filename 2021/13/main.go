package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Fold struct {
	XY   string
	Line int
}

//go:embed "input.txt"
var input string

func parse(input string) ([]Pair, []Fold) {
	blocks := strings.SplitN(input, "\n\n", 2)
	points := make([]Pair, 0)
	for _, line := range strings.Split(blocks[0], "\n") {
		tok := strings.Split(line, ",")
		x, ok := strconv.Atoi(tok[0])
		if ok != nil {
			panic("X conv")
		}
		y, ok := strconv.Atoi(tok[1])
		if ok != nil {
			panic("Y conv")
		}
		points = append(points, Pair{X: x, Y: y})
	}
	re := regexp.MustCompile(`fold along (.)=(\d+)`)
	folds := make([]Fold, 0)
	for _, line := range strings.Split(blocks[1], "\n") {
		match := re.FindStringSubmatch(line)
		num, ok := strconv.Atoi(match[2])
		if ok != nil {
			panic("num")
		}
		folds = append(folds, Fold{XY: match[1], Line: num})
	}
	return points, folds
}

func fold(points []Pair, fold Fold) []Pair {
	folder := make(map[Pair]struct{})
	if fold.XY == "y" {
		for _, p := range points {
			if p.Y > fold.Line {
				p.Y = 2*fold.Line - p.Y
			}
			folder[p] = struct{}{}
		}
	} else if fold.XY == "x" {
		for _, p := range points {
			if p.X > fold.Line {
				p.X = 2*fold.Line - p.X
			}
			folder[p] = struct{}{}
		}
	}
	ret := make([]Pair, 0)
	for k, _ := range folder {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	points, folding := parse(input)
	points = fold(points, folding[0])
	println(len(points))
	points, folding = parse(input)
	for k := range folding {
		points = fold(points, folding[k])
	}
	println(len(points))
	x := 0
	y := 0
	for _, p := range points {
		if x < p.X {
			x = p.X
		}
		if y < p.Y {
			y = p.Y
		}
	}
	x += 1
	y += 1
	grid := make([][]bool, y)
	for row := 0; row < y; row++ {
		grid[row] = make([]bool, x)
	}
	for _, p := range points {
		grid[p.Y][p.X] = true
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if grid[row][col] {
				print("X")
			} else {
				print(" ")
			}
		}
		println()
	}
}
