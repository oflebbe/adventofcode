package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type candleField []bool

const (
	Invalid = iota
	TurnOn
	TurnOff
	Toggle
)

func op(operation int, value bool) bool {
	switch operation {
	case TurnOn:
		return true
	case TurnOff:
		return false
	case Toggle:
		return !value
	default:
		log.Fatal("should not happen")
	}
	return false
}

func (c candleField) do(instruction string) {
	var state int
	if strings.HasPrefix(instruction, "turn on ") {
		state = TurnOn
		instruction = instruction[8:]
	} else if strings.HasPrefix(instruction, "turn off ") {
		state = TurnOff
		instruction = instruction[9:]
	} else if strings.HasPrefix(instruction, "toggle ") {
		state = Toggle
		instruction = instruction[7:]
	}
	var x1, y1, x2, y2 int
	fmt.Sscanf(instruction, "%d,%d  through %d,%d", &x1, &y1, &x2, &y2)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			c[x+y*1000] = op(state, c[x+y*1000])
		}
	}
}

func (c candleField) count() int {
	count := 0
	for _, i := range c {
		if i {
			count++
		}
	}
	return count
}

func main() {

	candles := make(candleField, 1000*1000)

	if candles.count() != 0 {
		log.Fatal("null")
	}
	candles.do("turn on 499,499 through 500,500")
	if candles.count() != 4 {
		log.Fatalf("4 %d", candles.count())
	}

	candles.do("turn off 499,499 through 500,500")
	if candles.count() != 0 {
		log.Fatal("00")
	}
	input, _ := os.Open("input.txt")
	line, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("readall")
	}

	lines := strings.Split(string(line), "\n")

	for _, l := range lines {
		candles.do(l)
	}
	fmt.Printf("Candles lit %d\n", candles.count())
}
