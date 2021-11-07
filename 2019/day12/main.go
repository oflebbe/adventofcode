package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type vec struct {
	coord [3]int
}

func (v *vec) Add(v2 *vec) {
	for i := range v.coord {
		v.coord[i] += v2.coord[i]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func (v *vec) SumAbs() int {
	ret := 0
	for i := range v.coord {
		ret += abs(v.coord[i])
	}
	return ret
}

func parse(st string) vec {
	//<x=13, y=-13, z=-2>
	ret := vec{}
	tok := strings.Split(st[1:len(st)-1], ", ")
	for i, val := range tok {
		v := strings.Split(val, "=")
		ret.coord[i], _ = strconv.Atoi(v[1])
	}
	return ret
}

func gravity(pl *[]vec, vel *[]vec) {
	lenpl := len(*pl)
	for i := 1; i < lenpl; i++ {
		for j := 0; j < i; j++ {
			for k := range (*pl)[i].coord {
				if (*pl)[i].coord[k] > (*pl)[j].coord[k] {
					(*vel)[i].coord[k]--
					(*vel)[j].coord[k]++
				} else if (*pl)[i].coord[k] < (*pl)[j].coord[k] {
					(*vel)[i].coord[k]++
					(*vel)[j].coord[k]--
				}
			}
		}
	}
}

func iter(pl *[]vec, vel *[]vec) {
	gravity(pl, vel)
	lenpl := len(*pl)
	for i := 0; i < lenpl; i++ {
		(*pl)[i].Add(&(*vel)[i])
	}
}

func energy(pl *[]vec, vel *[]vec) int {
	ret := 0
	lenpl := len(*pl)
	for i := 0; i < lenpl; i++ {
		ret += (*pl)[i].SumAbs() * (*vel)[i].SumAbs()
	}
	return ret
}

func print(pl *[]vec, vel *[]vec) {
	lenpl := len(*pl)
	for i := 0; i < lenpl; i++ {
		fmt.Printf("Pos %v Vel %v\n", (*pl)[i], (*vel)[i])
	}
}
func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}
	planets := make([]vec, 0)
	velo := make([]vec, 0)
	inputlines := strings.Split(string(input), "\n")
	for _, line := range inputlines {
		planets = append(planets, parse(line))
		velo = append(velo, vec{})
	}
	/*	for i := 0; i < 1000; i++ {
			iter(&planets, &velo)
			fmt.Println(energy(&planets, &velo))

		}
		print(&planets, &velo)
		fmt.Println(energy(&planets, &velo))
	*/
	config := make([]map[[8]int]int, 3)
	for i := 0; i < 3; i++ {
		config[i] = make(map[[8]int]int)
	}
	count := 0
	found := [3]bool{}
	nrfound := 0
	for {
		iter(&planets, &velo)
		for d := 0; d < 3; d++ {
			if !found[d] {
				c := [8]int{}
				for p := 0; p < 4; p++ {
					c[p] = planets[p].coord[d]
				}
				for p := 0; p < 4; p++ {
					c[p+4] = velo[p].coord[d]
				}
				r, ok := config[d][c]
				if ok {
					fmt.Println(d, r, count, energy(&planets, &velo))
					found[d] = true
					nrfound++
					if nrfound >= 3 {
						panic("found")
					}
				}
				config[d][c] = count
			}
		}
		count++
	}

}
