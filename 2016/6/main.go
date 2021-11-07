package main

import (
	"io/ioutil"
	"os"
	"strings"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func problem(input string) (string, string) {

	lineArray := strings.Split(input, "\n")
	linelength := len(lineArray[0])
	freq := make([]map[byte]int, linelength)
	for i := 0; i < linelength; i++ {
		freq[i] = make(map[byte]int)
	}

	for _, line := range lineArray {
		for i := 0; i < linelength; i++ {
			freq[i][line[i]]++
		}
	}

	maxCh := make([]byte, linelength)
	minCh := make([]byte, linelength)
	for i := 0; i < linelength; i++ {
		max := 0
		min := MaxInt
		for k, v := range freq[i] {
			if v > max {
				maxCh[i] = k
				max = v
			}
			if v < min {
				minCh[i] = k
				min = v
			}
		}
	}
	return string(maxCh), string(minCh)

}

func main() {
	input := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`
	println(problem(input))

	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	println(problem(string(buf)))

}
