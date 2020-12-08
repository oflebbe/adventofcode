package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func group(input string) int {

	checked := make(map[rune]interface{})
	for _, person := range strings.Fields(input) {
		for _, c := range person {
			checked[c] = nil
		}
	}
	return len(checked)
}

func group2(input string) int {

	checked := make(map[rune]int)
	persons := strings.Fields(input)
	for _, person := range persons {
		for _, c := range person {
			v, ok := checked[c]
			if ok {
				checked[c] = v + 1
			} else {
				checked[c] = 1
			}
		}
	}
	numPersons := len(persons)
	count := 0
	for _, v := range checked {
		if numPersons == v {
			count++
		}
	}

	return count
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	fh.Close()

	sum := 0
	sum2 := 0
	for _, block := range strings.Split(string(buf), "\n\n") {
		if block == "" {
			continue
		}
		sum += group(block)
		sum2 += group2(block)
	}
	println(sum)

	println(sum2)
}
