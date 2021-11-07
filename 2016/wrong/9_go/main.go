package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var d map[string]int
var sol []string

func recurse(dist int, city string, cities map[string]interface{}) (int, int) {
	if len(cities) == 1 {
		// cities

		fmt.Printf("%v: %d\n", sol, dist)

		return dist, dist
	}
	mycities := make(map[string]interface{})
	for k := range cities {
		if k != city {
			mycities[k] = nil
		}
	}
	min := 1000000 // ARgh
	max := 0
	for c := range mycities {
		sol = append(sol, c)
		name := fmt.Sprintf("%s_%s", city, c)
		dd := d[name]

		wayMin, wayMax := recurse(dist+dd, c, mycities)
		if wayMin < min {
			min = wayMin
		}
		if wayMax > max {
			max = wayMax
		}
		sol = sol[:len(sol)-1]
	}
	return min, max
}

func problem(input string) (int, int) {
	cities := make(map[string]interface{})
	d = make(map[string]int)
	for _, line := range strings.Split(input, "\n") {
		var c1, c2 string
		var d12 int
		fmt.Sscanf(line, "%s to %s = %d", &c1, &c2, &d12)
		cities[c1] = nil
		cities[c2] = nil
		n1 := fmt.Sprintf("%s_%s", c1, c2)
		d[n1] = d12
		n2 := fmt.Sprintf("%s_%s", c2, c1)
		d[n2] = d12
	}
	min := 1000000
	max := 0
	for c := range cities {
		wayMin, wayMax := recurse(0, c, cities)
		if wayMin < min {
			min = wayMin
		}
		if wayMax > max {
			max = wayMax
		}
	}
	return min, max
}
func main() {

	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	i, j := problem(input)
	fmt.Printf("%d %d\n", i, j)

	fh, _ := os.Open("input.txt")
	input2, _ := ioutil.ReadAll(fh)
	i2, j2 := problem(string(input2))
	fmt.Printf("%d %d\n", i2, j2)
}
