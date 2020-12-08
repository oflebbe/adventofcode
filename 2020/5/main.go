package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func seat(input string) (int, int, int) {
	input = strings.ReplaceAll(input, "B", "1")
	input = strings.ReplaceAll(input, "F", "0")
	input = strings.ReplaceAll(input, "R", "1")
	input = strings.ReplaceAll(input, "L", "0")

	row, err := strconv.ParseUint(input[0:7], 2, 64)
	if err != nil {
		panic("eins")
	}
	col, err := strconv.ParseUint(input[7:10], 2, 64)
	if err != nil {
		panic("zwei")
	}
	return int(row), int(col), int(row*8 + col)
}
func main() {
	expected_id := []int{567, 119, 820}
	for i, st := range []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"} {

		row, col, id := seat(st)
		fmt.Printf("%d %d %d\n", row, col, id)
		if id != expected_id[i] {
			panic("problem")
		}
	}

	fh, _ := os.Open("input.txt")
	defer fh.Close()
	buf, _ := ioutil.ReadAll(fh)
	max := 0

	for _, line := range strings.Split(string(buf), "\n") {
		_, _, id := seat(line)
		if max < id {
			max = id
		}
	}

	fmt.Printf("Highest %d\n", max)

	occupied := make([]bool, max+1)
	for _, line := range strings.Split(string(buf), "\n") {
		_, _, id := seat(line)
		occupied[id] = true
	}

	for i := 0; i < max-3; i++ {
		if occupied[i] == true && occupied[i+1] == false && occupied[i] == true {
			fmt.Printf("empty: %d", i+1)
		}
	}

}
