package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.Open("input.txt")
	in, _ := ioutil.ReadAll(f)
	ints := make([]int, 0)
	for _, l := range strings.Split(string(in), "\n") {
		v, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal("Atoi")
		}
		ints = append(ints, v)
	}

	for i := 0; i < len(ints); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				if ints[i]+ints[j]+ints[k] == 2020 {
					fmt.Printf("%d %d %d %d \n", ints[i], ints[j], ints[k], ints[i]*ints[j]*ints[k])
				}
			}
		}
	}

}
