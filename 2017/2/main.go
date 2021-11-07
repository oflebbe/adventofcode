package main

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func spreadsheet(input string) int64 {
	chksum := int64(0)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		//v := make([]int64, len(fields))
		min := int64(1) << 62
		max := int64(0)

		for _, f := range fields {
			//var err error
			v, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				panic("parse")
			}
			if v > max {
				max = v
			}
			if min > v {
				min = v
			}
		}
		chksum += max - min
	}
	return chksum
}

func spreadsheet2(input string) int {
	chksum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		v := make([]int, len(fields))

		for i, f := range fields {
			zw, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				panic("parse")
			}
			v[i] = int(zw)
		}
		sort.Ints(v)
		for i := 1; i < len(v); i++ {
			for j := 0; j < i; j++ {
				if v[i]%v[j] == 0 {
					chksum += v[i] / v[j]
					goto out
				}
			}
		}
	out:
	}
	return chksum
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	fh.Close()
	println(spreadsheet(string(buf)))
	println(spreadsheet2(string(buf)))
}
