package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

//Sue 13: "perfumes": 6, "goldfish": 1, "cars": 8
var re = regexp.MustCompile(`Sue (\d+): (\w+): (-?\d+), (\w+): (-?\d+), (\w+): (-?\d+)`)

// list of sues : map of attributes
func parse(st string) []map[string]int {
	ret := make([]map[string]int, 0)
	for k, line := range re.FindAllStringSubmatch(st, -1) {
		number, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		if k+1 != number {
			panic("Sue # isn't aligned")
		}
		sueProp := make(map[string]int)
		// Sue has 3 properties at line[p*2+2]
		for p := 0; p < 3; p++ {
			v, err := strconv.Atoi(line[2*p+3])
			if err != nil {
				panic(err)
			}
			sueProp[line[2*p+2]] = v
		}
		ret = append(ret, sueProp)
	}
	return ret
}

var readings = map[string]int{"children": 3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1}

func eval(prop []map[string]int) int {
	for k, ant := range prop {
		found := true
		for k, v := range ant {
			fv, ok := readings[k]
			if !ok {
				found = false
				break
			}
			if fv != v {
				found = false
				break
			}
		}
		if found {
			return k + 1
		}
	}
	return -1
}

func eval2(prop []map[string]int) int {
	for k, ant := range prop {
		found := true
		for k, v := range ant {
			fv, ok := readings[k]
			if !ok {
				found = false
				break
			}
			switch k {
			case "trees", "cats":
				if fv >= v {
					found = false
					break
				}
			case "pomeranians", "goldfish":
				if fv <= v {
					found = false
					break
				}
			default:
				if fv != v {
					found = false
					break
				}
			}
		}
		if found {
			return k + 1
		}
	}

	return -1
}

func main() {
	val := parse(input)
	fmt.Printf("%d\n", eval(val))
	fmt.Printf("%d\n", eval2(val))
}
