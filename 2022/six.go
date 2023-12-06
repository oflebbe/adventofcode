package main

import (
	"fmt"
	"io/ioutil"
)

func startSeq(s []byte, num int) int {
	for ptr := 0; ptr < len(s)-num; ptr++ {
		counter := make(map[byte]int)
		for i := 0; i < num; i++ {
			counter[s[ptr+i]] = 1
		}
		if len(counter) == num {
			return ptr + num
		}
	}
	return 0
}

func tests() {
	c := startSeq([]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 4)
	if c != 7 {
		fmt.Printf("%d rather 7", c)
	}
}

func main() {
	// tests()
	buf, _ := ioutil.ReadFile("input.txt")
	println(startSeq(buf, 14))
}
