package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sum(input string) int {
	length := len(input)
	last := input[length-1]
	sum := 0
	for i := 0; i < len(input); i++ {
		if last == input[i] {
			sum += int(input[i] - '0')
		}
		last = input[i]
	}
	return sum
}

func sum2(input string) int {
	length := len(input)
	sum := 0
	for i := 0; i < len(input); i++ {
		other := (i + length/2) % length
		if input[other] == input[i] {
			sum += int(input[i] - '0')
		}
	}
	return sum
}

func main() {
	inputs := []string{"1122", "1111", "1234", "91212129"}
	results := []int{3, 4, 0, 9}

	for i, j := range inputs {
		fmt.Printf("%d %d\n", sum(j), results[i])
		if sum(j) != results[i] {
			panic("incorrect")
		}
	}

	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	fh.Close()
	fmt.Printf("%d\n", sum(string(buf)))

	fmt.Printf("%d\n", sum2(string(buf)))
}
