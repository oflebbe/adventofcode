package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("ReadFile %v", err)
	}
	lines := strings.Split(string(buf), "\n")
	col1 := make([]int, 0)
	col2 := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			break
		}
		var one int
		var two int
		e, err := fmt.Sscanf(line, "%d %d", &one, &two)
		if err != nil || e != 2 {
			log.Fatalf("line %s not scannable", line)
		}
		col1 = append(col1, one)
		col2 = append(col2, two)
	}
	slices.Sort(col1)
	slices.Sort(col2)

	sum := 0
	for i, v := range col1 {
		sum += Abs(v - col2[i])
	}
	fmt.Printf("Task1 %d\n", sum)

	lastValue := -1
	lastScore := 0
	totalScore := 0
	i2 := 0
	for _, v := range col1 {
		if v == lastValue {
			totalScore += lastScore
			continue
		}
		for v > col2[i2] {
			i2++
			if i2 >= len(col2) {
				break
			}
		}
		score := 0
		if i2 >= len(col1) {
			break
		}
		for v == col2[i2] {
			score++
			i2++
			if i2 >= len(col2) {
				break
			}
		}
		totalScore += score * v
		lastValue = v
		lastScore = score * v
	}
	fmt.Printf("Task2: %d\n", totalScore)
}
