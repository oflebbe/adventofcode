package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type point struct {
	x, y int
	phi  float64
	d    int
}

func euclid(a, b int) int {
	for b != 0 {
		h := a % b
		a = b
		b = h
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}

}
func simplify(a, b int) (int, int) {
	c := euclid(abs(a), abs(b))
	a /= c
	b /= c
	return a, b
}

func parse(lines []string) []point {
	ret := make([]point, 0)
	for y, l := range lines {
		l = strings.Trim(l, " ")
		for x, i := range l {
			if i == '#' {
				p := point{x, y, 0., 0}
				ret = append(ret, p)
			}
		}
	}
	return ret
}

type planetSorter struct {
	planets []point
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
/* func (s *planetSorter) Less(i, j int) bool {
	pi := s.planets[i]
	pj := s.planets[j]
	return pi.x*pi.x+pi.y*pi.y < pj.x*pj.x+pj.y*pj.y
}
*/
func (s *planetSorter) Less(i, j int) bool {
	pi := s.planets[i]
	pj := s.planets[j]
	return pi.phi < pj.phi || (pi.phi == pj.phi && pi.d < pj.d)
}

func angle(x, y int) float64 {
	if y == 0 {
		if x > 0 {
			return 0.
		}
		return 180.
	}
	phi := math.Atan(float64(x)/float64(y)) / math.Pi * 180.

	return phi
}

func dist(x, y int) int {
	return x*x + y*y
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	lines := strings.Split(string(input), "\n")
	list := parse(lines)

	for i, v := range list {
		list[i].phi = angle(20-v.x, 19-v.y)
		list[i].d = dist(v.x-20, v.y-19)
	}
	p := planetSorter{list}
	sort.Sort(&p)
	fmt.Printf("%+v", p)
	//p, m := findAstro(list)

	//fmt.Printf("%+v,%d\n", p, m)
}
