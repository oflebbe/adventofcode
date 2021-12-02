package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func part1() {
	previous := 100000000
	count := 0
	for _, line := range strings.Split(input, "\n") {
		depth, _ := strconv.Atoi(line)
		if previous < depth {
			count++
		}
		previous = depth
	}
	println(count)
}

func part2() {
	lines := strings.Split(input, "\n")
	depths := make([]int, len(lines))
	for k, v := range lines {
		depths[k], _ = strconv.Atoi(v)
	}
	previous := 100000000000
	count := 0

	for nr := 2; nr < len(lines); nr++ {
		current := depths[nr] + depths[nr-1] + depths[nr-2]
		if previous < current {
			count++
		}
		previous = current
	}
	println(count)
}

func main() {
	part1()
	part2()
}
