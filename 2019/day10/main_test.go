package main

import (
	"math"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := `.#..#
.....
#####
....#
...##`
	lines := strings.Split(input, "\n")

	list := parse(lines)
	if len(list) != 10 {
		t.Error("Main")
	}
	p, m := findAstro(list)
	if p.x != 3 && p.y != 4 && m != 8 {
		t.Error("point")
	}
}

func TestSimply(t *testing.T) {
	a, b := simplify(10, 5)
	if a != 2 && b != 1 {
		t.Error("TestSimply")

	}
}

func TestSimply2(t *testing.T) {

	a, b := simplify(-10, 5)
	if a != -2 && b != 1 {
		t.Error("TestSimply")
	}

}

func TestSimply3(t *testing.T) {

	a, b := simplify(-10, -300)
	if a != -1 && b != -30 {
		t.Error("TestSimply")
	}

}

func TestLarge(t *testing.T) {
	input := `.#..##.###...#######
	##.############..##.
	.#.######.########.#
	.###.#######.####.#.
	#####.##.#.##.###.##
	..#####..#.#########
	####################
	#.####....###.#.#.##
	##.#################
	#####.##.###..####..
	..######..##.#######
	####.##.####...##..#
	.#####..#.######.###
	##...#.##########...
	#.##########.#######
	.####.#.###.###.#.##
	....##.##.###..#####
	.#.#.###########.###
	#.#.#.#####.####.###
	###.##.####.##.#..##`

	lines := strings.Split(input, "\n")

	list := parse(lines)

	p, m := findAstro(list)
	if p.x != 11 && p.y != 13 && m != 210 {
		t.Error("point")
	}
}

func TestSmall(t *testing.T) {
	input := `......#.#.
	#..#.#....
	..#######.
	.#.#.###..
	.#..#.....
	..#....#.#
	#..#....#.
	.##.#..###
	##...#..#.
	.#....####`

	lines := strings.Split(input, "\n")

	list := parse(lines)

	p, m := findAstro(list)
	if p.x != 5 && p.y != 8 && m != 33 {
		t.Error("point")
	}
}

func TestAngle(t *testing.T) {
	if angle(0, 1) != 0. {
		t.Error("Angle0")
	}
	if angle(1, 0) != math.Pi/2. {
		t.Error("Angle10")
	}
	if angle(-1, 0) != math.Pi*3./2. {
		t.Error("Anglem10")
	}
}
