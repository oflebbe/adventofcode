package main

import "strings"

import "fmt"

import "strconv"

import "io/ioutil"

func calculate(str string) int {
	b := strings.Split(str, "\n")
	sum := 0
	for _, v := range b {
		module, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			panic("atoi" + v)
		}
		sum += rocket(module/3 - 2)
	}
	return sum
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	sum := calculate(string(input))
	fmt.Println(sum)
}

func rocket(mass int) int {
	add := mass/3 - 2
	for add > 0 {
		mass += add
		add = add/3 - 2
	}
	return mass
}
