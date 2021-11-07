package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func solve(input string) uint64 {
	lines := strings.Split(input, "\n")
	time, _ := strconv.Atoi(lines[0])
	busLines := lines[1]
	earliest := uint64(0)
	minTime := uint64(1 << 62)
	for _, bus := range strings.Split(busLines, ",") {
		if bus == "x" {
			continue
		}
		nr, _ := strconv.Atoi(bus)
		diff := uint64((nr * ((time / nr) + 1)) - time)
		if diff < minTime {
			earliest = uint64(nr) * diff
			minTime = diff
		}
	}
	return earliest
}

func parse3(input string) ([]uint64, []uint64) {
	offset := uint64(0)
	var diff []uint64
	var nums []uint64
	for _, bus := range strings.Split(input, ",") {
		if bus == "x" {
			offset++
			continue
		}
		diff = append(diff, offset)
		offset++
		nr, _ := strconv.Atoi(bus)
		nums = append(nums, uint64(nr))
	}
	fmt.Printf("%d ", nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Printf(" - %d - %d", diff[i], nums[i])
	}
	println()
	return nums, diff
}

func parse2(input string) ([]uint64, []uint64) {
	lines := strings.Split(input, "\n")
	return parse3(lines[1])
}

func solve2(nums, diffs []uint64) uint64 {
	step := nums[0]
	v := uint64(0)
	i := 1
	D := nums[1] - diffs[1]
	for {
		if v%nums[i] == D {
			step *= nums[i]
			i++
			if i >= len(nums) {
				break
			}
			DD := int(nums[i]) - int(diffs[i])
			for DD < 0 {
				DD += int(nums[i])
			}
			D = uint64(DD)
		}
		v += step
	}
	return v
}

func main() {
	n, d := parse3("17,x,13,19")
	if 3417 != solve2(n, d) {
		panic("assert")
	}

	testInput := `939
7,13,x,x,59,x,31,19`
	n, d = parse2(testInput)
	if 1068781 != solve2(n, d) {
		panic("assert2")
	}

	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)

	n, d = parse2(string(buf))
	fmt.Printf("%d\n", solve2(n, d))
}
