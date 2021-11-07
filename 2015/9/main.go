package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Pair struct {
	a string
	b string
}

func NewPair(a, b string) pair {
	if a > b {
		return Pair{a: a, b: b}
	} else {
		return Pair{a: b, b: a}
	}
}

func solve() {


func travel(targets map[string]int, distance[depth int, distance int) int {
	last := ""
	for k, v := range targets {
		if v == depth-1 {
			last = k
			break
		}
	}
	for k, v := range targets {
		if v < 0 {
			targets[k] = depth
			d := 
			travel(targets, depth+1)
		}
	}
	return
}
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("open")
	}
	in, _ := ioutil.ReadAll(f)

	distance := make(map[Pair]int)
	targets := make(map[string]interface{}, 0)
	for _, s := range strings.Split(string(in), "\n") {
		var a, b string
		var d int
		_, err = fmt.Sscanf(s, "%s to %s = %d", &a, &b, &d)
		if err != nil {
			log.Fatal("Scanf")
		}
		distance[NewPair(a, b)] = d
		targets[a] = -1
		targets[b] = -1
	}
	travel()

}
