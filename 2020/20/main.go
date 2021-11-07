package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func mirror(st string) string {
	buf := []byte(st)
	length := len(buf)
	for i := 0; i < length/2; i++ {
		buf[i], buf[length-1-i] = buf[length-1-i], buf[i]
	}
	return string(buf)
}

func tile(part string) ([4]string, int) {
	var num int
	lines := strings.Split(part, "\n")
	fmt.Sscanf(lines[0], "Tile %d:", &num)
	length := len(lines) - 1
	first := lines[1]
	last := lines[length]
	side1 := make([]byte, length, length)
	side2 := make([]byte, length, length)

	for i := 1; i <= length; i++ {
		side1[i-1] = lines[i][length-1]
		side2[i-1] = lines[i][0]
	}
	return [4]string{first, string(side1), last, string(side2)}, num
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	parts := strings.Split(string(buf), "\n\n")
	dict := make(map[string][]int)
	for _, p := range parts {
		if p == "" {
			continue
		}
		s, num := tile(p)
		for _, i := range s {
			array, ok := dict[i]
			if ok {
				dict[i] = append(array, num)
			} else {
				dict[i] = make([]int, 1, 2)
				dict[i][0] = num
			}
			i2 := mirror(i)
			array, ok = dict[i2]
			if ok {
				dict[i2] = append(array, num)
			} else {
				dict[i2] = make([]int, 1, 2)
				dict[i2][0] = num
			}
			i3 := mirror(i2)
			if i3 != i {
				panic("olala")
			}
		}
	}
	count := make(map[int]int)
	allcount := 0
	for _, v := range dict {
		if len(v) == 1 {
			// fmt.Printf("%d %s\n", v, k)
			count[v[0]]++
			allcount++
		}
	}
	println(allcount)
	chk := uint64(1)
	for k, v := range count {
		if v == 4 {
			chk *= uint64(k)
		}
	}
	println(chk)

}
