package main

import (
	"fmt"
	"sort"
)

var container = []int{33,
	14,
	18,
	20,
	45,
	35,
	16,
	35,
	1,
	13,
	18,
	13,
	50,
	44,
	48,
	6,
	24,
	41,
	30,
	42}

var count = make(map[int]int)

func score(remaining int, containers []int, num int) {
	if remaining == 0 {
		count[num]++
		return
	}
	for i := 0; i < len(containers); i++ {
		if containers[i] <= remaining {
			score(remaining-containers[i], containers[i+1:], num+1)
		}
	}
}

func main() {
	sort.Sort(sort.Reverse(sort.IntSlice(container)))
	score(150, container, 0)
	sum := 0
	for _, v := range count {
		sum += v
	}
	fmt.Printf("%d\n", sum)

	for i := 0; i < len(container); i++ {
		if count[i] > 0 {
			fmt.Printf("%d\n", count[i])
			break
		}
	}
}
