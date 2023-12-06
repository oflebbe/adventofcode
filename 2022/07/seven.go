package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	byteBuf, _ := ioutil.ReadFile("input.txt")
	blocks := strings.Split(string(byteBuf), "\n$ cd ")
	blocks[0] = strings.Replace(blocks[0], "$ cd /", "/", 1)
	dirStack := []string{}
	dirSize := make(map[string]int)
	for _, b := range blocks {
		lines := strings.Split(b, "\n")
		if lines[0] == ".." {
			dirStack = dirStack[:len(dirStack)-1]
		} else {
			dirStack = append(dirStack, lines[0])
		}
		if len(lines) == 1 {
			continue
		}
		if lines[1] != "$ ls" {
			fmt.Printf("ls expected")
		}
		for _, e := range lines[2:] {
			if strings.HasPrefix(e, "dir ") {
				continue
			}
			size := 0
			n, err := fmt.Sscanf(e, "%d ", &size)
			if n != 1 || err != nil {
				fmt.Printf("Error %d %v", n, err)
			}
			for i := 1; i <= len(dirStack); i++ {
				dirSize[strings.Join(dirStack[:i], "/")] += size
			}
		}
	}

	sum := 0
	for _, v := range dirSize {
		if v <= 100000 {
			sum += v

		}
	}
	fmt.Printf("%d\n", sum)
	currentFree := 70000000 - dirSize["/"]
	needFree := 30000000 - currentFree
	minMax := dirSize["/"]
	for k, v := range dirSize {
		if v >= needFree && minMax > v {
			minMax = v
			fmt.Printf("%s, %d\n", k, v)
		}
	}

	fmt.Printf("%d\n", minMax)
}
