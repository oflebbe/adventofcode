package main

import (
	"crypto/md5"
	"fmt"
)

func try(st string, numZeros int) {
	i := 1
	for {
		s := fmt.Sprintf("%s%d", st, i)
		x := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		for j := 0; j < numZeros; j++ {
			if x[j] != '0' {
				goto next
			}
		}
		fmt.Printf("Solution for %s: %d\n", st, i)
		return
	next:
		i++
	}
}

func main() {

	start := "abcdef"
	try(start, 5)
	try("pqrstuv", 5)
	try("ckczppom", 5)
	try("ckczppom", 6)
}
