package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func scalarp(inpArray, mulArray []int, i, count int) int {
	sum := 0
	L := len(inpArray)
	for j := 0; j < count; j++ {
		sum += inpArray[j%L] * mulArray[((j+1)/i%4]
	}
	return sum % 10
}

func phase(input string) string {
	L := len(input)

	inpArray := make([]int, L)
	for i := 0; i < len(input); i++ {
		j, _ := strconv.Atoi(string(input[i]))
		inpArray[i] = j
	}

	helparray := [4]int{0, 1, 0, -1}
	ret := make([]byte, L)
	for i := 1; i <= L*10000; i++ {
		lcm := LCM(L, 4*i)
		scalarp( inpArray, mulArray, i, LCM(L, 4*i)

		sum = abs(sum % 10)
		ret[i-1] = byte(0x30 + sum)
	}
	return string(ret)
}

func main() {

	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic("ioutil")
	}

	st := string(input)
	for i := 0; i < 100; i++ {
		st += string(input)
	}
	off, _ := strconv.Atoi(st[0:7])
	for p := 0; p < 100; p++ {
		fmt.Println(p)
		st = phase(st)
	}
	off = 0
	fmt.Println(st[off : off+8])

}
