package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type candleField []int

const (
	TurnOn = iota
	TurnOff
	Toggle
)

func op(operation int, value int) int {
	switch operation {
	case TurnOn:
		return value + 1
	case TurnOff:
		r := value - 1
		if r < 0 {
			r = 0
		}
		return r
	case Toggle:
		return value + 2
	default:
		log.Fatal("should not happen")
	}
	return 0
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
		count += i
	}
	return count
}

func main() {

	candles := make(candleField, 1000*1000)

	if candles.count() != 0 {
		log.Fatal("null")
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
	fmt.Printf("Power of Candles lit %d\n", candles.count())
	im := image.NewGray(image.Rect(0, 0, 1000, 1000))
	for i, c := range candles {
		x := i % 1000
		y := i / 1000
		im.SetGray(x, y, color.Gray{Y: uint8(c) * 2})
	}
	w, _ := os.Create("image.png")
	png.Encode(w, im)
	w.Close()
}
