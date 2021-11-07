package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type receipt struct {
	howmuch  int
	recipies []receiptPart
}

type receiptPart struct {
	number int
	what   string
}

func parse(lines []string) map[string]receipt {
	all := make(map[string]receipt)
	for _, l := range lines {
		toks := strings.Split(l, " => ")
		rights := toks[1]
		var count int
		var ingredient string
		fmt.Sscanf(rights, "%d %s", &count, &ingredient)

		lefts := strings.Split(toks[0], ", ")
		receipies := make([]receiptPart, 0)

		for _, left := range lefts {
			var count int
			var ingredient string
			fmt.Sscanf(left, "%d %s", &count, &ingredient)
			receipies = append(receipies, receiptPart{count, ingredient})
		}
		result := receipt{count, receipies}
		_, ok := all[ingredient]
		if ok {
			panic("not easy")
		}
		all[ingredient] = result
	}
	return all
}

func addIng(sum *map[string]int, append map[string]int) {
	for k, v := range append {
		(*sum)[k] += v
	}
}

func add(recipies map[string]receipt, rest *map[string]int, count int, what string) int {
	sumOre := 0
	val, ok := (*rest)[what]
	if ok {
		if count <= val {
			(*rest)[what] = val - count
			return 0
		}
		(*rest)[what] = 0
		count -= val
	}
	valR, okR := recipies[what]
	if !okR {
		if what != "ORE" {
			panic("ORE!" + what)
		} else {
			return count
		}
	}
	factor := (count-1)/valR.howmuch + 1

	for _, v := range valR.recipies {
		ingAdd := add(recipies, rest, v.number*factor, v.what)
		sumOre += ingAdd
	}
	(*rest)[what] = factor*valR.howmuch - count

	return sumOre
}

func try(all map[string]receipt, n int) int {
	rest := make(map[string]int)
	return add(all, &rest, n, "FUEL")
}

func binary(all map[string]receipt, target int) int {
	n := 1
	var result int
	for ; result < target; n *= 2 {
		result = try(all, n)
		fmt.Println(n, result-target)
	}
	high := n / 2
	low := n / 4
	for high-low >= 2 {
		test := (high + low) / 2
		if test == low {
			test = test + 1
		}
		if test == high {
			test = test - 1
		}
		result = try(all, test)
		fmt.Println(test, result-target)
		if result > target {
			high = test
		} else if result < target {
			low = test
		} else {
			return test
		}
	}
	return low
}

func mainFirst() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	inputlines := strings.Split(string(input), "\n")
	all := parse(inputlines)
	rest := make(map[string]int)
	end := add(all, &rest, 1, "FUEL")
	fmt.Printf("%d\n", end)

}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	inputlines := strings.Split(string(input), "\n")
	all := parse(inputlines)
	end := binary(all, 1000000000000)

	fmt.Printf("%d\n", end)

}
