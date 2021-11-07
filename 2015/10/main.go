package main

import "fmt"

func transform(in []int) []int {
	lastDigit := 0
	count := 0
	output := []int{}
	for _, d := range in {
		if count == 0 {
			lastDigit = d
			count = 1
		} else {
			if lastDigit == d {
				count++
			} else {
				output = append(output, count, lastDigit)
				lastDigit = d
				count = 1
			}
		}
	}
	output = append(output, count, lastDigit)
	return output
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	digits := []int{3, 1, 1, 3, 3, 2, 2, 1, 1, 3}
	for i := 0; i < 40; i++ {
		digits = transform(digits)
	}

	fmt.Println(len(digits))

	digits = []int{3, 1, 1, 3, 3, 2, 2, 1, 1, 3}
	for i := 0; i < 50; i++ {
		digits = transform(digits)
	}

	fmt.Println(len(digits))
}
