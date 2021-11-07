package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func s(w, l, h int) int {
	a1 := l * w
	a2 := w * h
	a3 := h * l

	a := 2 * (a1 + a2 + a3)
	min := a1
	if min > a2 {
		min = a2
	}
	if min > a3 {
		min = a3
	}
	return a + min
}

func r(w, l, h int) int {
	a1 := l + w
	a2 := w + h
	a3 := h + l

	min := a1
	if min > a2 {
		min = a2
	}
	if min > a3 {
		min = a3
	}

	return 2*min + l*w*h
}
func main() {
	input, _ := os.Open("input")
	lines, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("readall")
	}
	sumS := 0
	sumR := 0
	for _, line := range strings.Split(string(lines), "\n") {
		var w int
		var l int
		var h int
		fmt.Sscanf(line, "%dx%dx%d", &w, &l, &h)
		sumS += s(w, l, h)
		sumR += r(w, l, h)
	}
	fmt.Printf("area %d, ribbon %d", sumS, sumR)
}
