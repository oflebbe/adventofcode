package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	area := make(map[string]int)
	input, _ := os.Open("input")
	lines, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("readall")
	}
	x := 0
	y := 0
	x2 := 0
	y2 := 0
	area["0,0"] = 2
	for c, ch := range lines {
		if c%2 == 0 {
			switch ch {
			case '^':
				y--
			case 'v':
				y++
			case '>':
				x++
			case '<':
				x--
			default:
				log.Fatal("ooo")
			}
			index := fmt.Sprintf("%d,%d", x, y)

			val, ok := area[index]
			if !ok {
				area[index] = 1
			} else {
				area[index] = val + 1
			}
		} else {
			switch ch {
			case '^':
				y2--
			case 'v':
				y2++
			case '>':
				x2++
			case '<':
				x2--
			default:
				log.Fatal("oooxx")
			}
			index := fmt.Sprintf("%d,%d", x2, y2)

			val, ok := area[index]
			if !ok {
				area[index] = 1
			} else {
				area[index] = val + 1
			}
		}
	}
	fmt.Printf("area %d", len(area))
}
