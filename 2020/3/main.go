package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func problem(input string, xstep, ystep int) int {
	lines := strings.Split(input, "\n")
	YSIZE := len(lines)
	XSIZE := len(lines[0])
	y := 0
	x := 0
	tree := 0
	for y < YSIZE {
		ch := lines[y][x]
		if ch == '#' {
			tree++
		}
		x += xstep
		x %= XSIZE
		y += ystep
	}
	return tree
}

func main() {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	println(problem(input, 3, 1))

	fh, _ := os.Open("input.txt")
	defer fh.Close()
	buf, _ := ioutil.ReadAll(fh)
	println(problem(string(buf), 3, 1))

	xsteps := []int{1, 3, 5, 7, 1}
	ysteps := []int{1, 1, 1, 1, 2}
	res := 1
	res2 := 1
	for s, xs := range xsteps {
		ys := ysteps[s]
		res *= problem(input, xs, ys)
		res2 *= problem(string(buf), xs, ys)
	}
	println(res)
	println(res2)
}
