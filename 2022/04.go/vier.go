package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"strings"
)

func fully(l1, r1, l2, r2 int) bool {
	if l1 == l2 || r1 == r2 {
		return true
	}
	if l1 < l2 {
		return r1 > r2
	}
	return r1 < r2
}

func partial(l1, r1, l2, r2 int) bool {
	return !(r1 < l2 || l1 > r2)
}

func find_elfs(name string) (int, int) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	sbuf := string(buf)
	lines := strings.Split(sbuf, "\n")
	count1 := 0
	count2 := 0
	for _, line := range lines {
		var l1, r1, l2, r2 int
		n, err := fmt.Sscanf(line, "%d-%d,%d-%d", &l1, &r1, &l2, &r2)

		if err != nil || n != 4 {
			log.Fatalf("error parsing elves %v %d", err, n)
		}
		if l1 > r1 || l2 > r2 {
			log.Fatal("inconsistent elf")
		}
		if fully(l1, r1, l2, r2) {
			count1++
		}
		if partial(l1, r1, l2, r2) {
			count2++
		}
	}
	return count1, count2
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		fmt.Printf("Hallo")
		defer pprof.StopCPUProfile()
	}

	count1, count2 := find_elfs("input.txt")
	fmt.Printf("%d, %d\n", count1, count2)

}
