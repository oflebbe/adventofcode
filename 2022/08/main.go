package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func visible(arr [][]int, x, y, dx, dy int) (bool, int) {
	sz := len(arr)
	tree := arr[y][x]

	x += dx
	y += dy
	count := 0
	for x >= 0 && y >= 0 && x < sz && y < sz {
		count++
		if arr[y][x] >= tree {
			return false, count
		}
		x += dx
		y += dy
	}
	return true, count
}

func visibleAll(arr [][]int, x, y int) (bool, int) {
	sz := len(arr)
	if x == 0 || y == 0 || x == sz-1 || y == sz-1 {
		return true, 0
	}
	vis := false
	score := 1
	for i := 0; i < 4; i++ {
		dx := 0
		dy := 0
		switch i {
		case 0:
			dy = 0
			dx = 1
		case 1:
			dy = 0
			dx = -1
		case 2:
			dy = 1
			dx = 0
		case 3:
			dy = -1
			dx = 0
		}
		v, c := visible(arr, x, y, dx, dy)
		vis = vis || v
		score *= c

	}
	return vis, score
}

func main() {

	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\n")
	sz := len(lines)
	arr := make([][]int, 0)
	for k, l := range lines {
		arr = append(arr, []int{})
		for i := 0; i < len(l); i++ {
			arr[k] = append(arr[k], int(l[i]-'0'))
		}
	}
	fmt.Printf("%+v", arr)

	count := 0
	maxScore := 0
	for y := 0; y < sz; y++ {
		for x := 0; x < sz; x++ {
			v, s := visibleAll(arr, x, y)
			if v {
				count++
			}
			if maxScore < s {
				maxScore = s
			}
		}
	}
	fmt.Printf("%d\n", count)
	fmt.Printf("%d\n", maxScore)
}
