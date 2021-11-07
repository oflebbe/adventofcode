package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type pic struct {
	w, h     int
	nrLayers int
	layers   []*layer
}

type layer struct {
	layer []int
}

func readInput(w, h int, st string) pic {
	if len(st)%(w*h) != 0 {
		panic("string does not fit")
	}

	nrLayers := len(st) / (w * h)
	picInstance := pic{w, h, nrLayers, make([]*layer, 0)}
	for i := 0; i < nrLayers; i++ {
		l := layer{make([]int, w*h)}
		picInstance.layers = append(picInstance.layers, &l)
		for j := 0; j < w*h; j++ {
			l.layer[j], _ = strconv.Atoi(string(st[i*(w*h)+j]))
		}
	}
	return picInstance
}

func decode(p pic) string {
	layer := make([]int, p.w*p.h)
	for i := 0; i < p.w*p.h; i++ {
		for j := 0; j < p.nrLayers; j++ {
			pixel := p.layers[j].layer[i]
			if pixel != 2 {
				layer[i] = pixel
				break
			}
		}
	}
	ret := ""
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			pixel := layer[y*p.w+x]
			if pixel == 0 {
				ret += " "
			} else {
				ret += "X"
			}
		}
		ret += "\n"
	}
	return ret

}
func score(p pic) int {
	min0 := p.w * p.h

	score := 0
	for i := 0; i < p.nrLayers; i++ {
		counter := [10]int{}
		for j := 0; j < p.w*p.h; j++ {
			counter[p.layers[i].layer[j]]++
		}
		if counter[0] < min0 {

			min0 = counter[0]
			score = counter[1] * counter[2]
		}
	}
	return score
}
func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	pic := readInput(25, 6, string(input))
	fmt.Println(score(pic))
	fmt.Print(decode(pic))
}
