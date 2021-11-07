package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	inputlines := strings.Split(string(input), "\n")
	fmt.Println(len(inputlines))
}
