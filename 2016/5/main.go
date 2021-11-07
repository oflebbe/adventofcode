package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func problem(input string) string {
	var count int64 = 0

	passwd := ""
	for {
		st := fmt.Sprintf("%s%d", input, count)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(st)))
		if strings.HasPrefix(hash, "00000") {
			var ch byte
			ch = hash[5]
			passwd += string(ch)
			if len(passwd) == 8 {
				return passwd
			}
		}
		count++

	}

}

func problem2(input string) string {
	var count int64 = 0

	passwd := make(map[int]byte)
	for {
		st := fmt.Sprintf("%s%d", input, count)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(st)))
		if strings.HasPrefix(hash, "00000") {
			pos := int(hash[5]) - '0'
			if pos > 7 {
				count++
				continue
			}
			val := hash[6]
			_, ok := passwd[pos]
			if !ok {
				passwd[pos] = val
			}

			if len(passwd) == 8 {
				passwdStr := ""
				for i := 0; i < 8; i++ {
					passwdStr += string(passwd[i])
				}
				return passwdStr
			}
		}
		count++

	}

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
		defer pprof.StopCPUProfile()
	}
	println(problem("abc"))
	/* println(problem("ffykfhsq"))
	println(problem2("abc"))
	println(problem2("ffykfhsq")) */
}
